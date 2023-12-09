// Package piecewiselinear is a tiny library for linear interpolation.
package piecewiselinear

import (
	"math"
)

// Function is a piecewise-linear 1-dimensional function
type Function struct {
	X []float64
	Y []float64
}

// Area returns the definite integral of the function on its domain X.
//
// Time complexity: O(N), where N is the number of points.
// Space complexity: O(1)
func (f Function) Area() (area float64) {
	X, Y := f.X, f.Y
	for i := 1; i < len(X); i++ {
		area += (X[i] - X[i-1]) * (Y[i] + Y[i-1]) / 2
	}
	return area
}

// AreaUpTo returns the definite integral of the function on its domain X intersected with [-Inf, x].
//
// Time complexity: O(N), where N is the number of points.
// Space complexity: O(1)
func (f Function) AreaUpTo(x float64) (area float64) {
	X, Y := f.X, f.Y
	for i := 1; i < len(X); i++ {
		dX := X[i] - X[i-1]
		if x < X[i] {
			if x >= X[i-1] {
				dxX := x - X[i-1]
				w := dxX / dX
				y := (1-w)*Y[i-1] + w*Y[i]
				area += dxX * (y + Y[i-1]) / 2
			}
			return area
		}
		area += dX * (Y[i] + Y[i-1]) / 2
	}
	return area
}

// IsInterpolatedAt returns true if x is within the given range of points, false if outside of that range
func (f Function) IsInterpolatedAt(x float64) bool {
	n := len(f.X)
	if n == 0 {
		return false
	}
	left, right := f.X[0], f.X[n-1]
	return x >= left && x <= right
}

// At returns the function's value at the given point.
// Outside its domain X, the function is constant at 0.
//
// The function's X and Y slices are expected to be the same legnth. The length property is _not_ verified.
// The function's X slice is expected to be sorted in ascending order. The sortedness property is _not_ verified.
//
// Time complexity: O(log(N)), where N is the number of points.
// Space complexity: O(1)
func (f Function) At(x float64) float64 {
	X, Y := f.X, f.Y
	i, j := 0, len(X)
	for i < j {
		h := int(uint(i+j) >> 1)
		if X[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	if i == 0 {
		if len(X) > 0 && x == X[0] {
			return Y[0]
		}
		return 0
	}
	if i == len(X) {
		return 0
	}
	w := (x - X[i-1]) / (X[i] - X[i-1])
	return (1-w)*Y[i-1] + w*Y[i]
}

// Function is a piecewise-linear 1-dimensional function with uniformly spaced control points. Faster to interpolate than the equivalent `Function`.
type FunctionUniform struct {
	Xmin, Xmax float64
	Y          []float64
}

// At returns the function's value at the given point.
// Outside its domain X, the function is constant at 0.
//
// Time complexity: O(1)
// Space complexity: O(1)
func (f FunctionUniform) At(x float64) float64 {
	if x < f.Xmin {
		return 0
	}
	if x > f.Xmax {
		return 0
	}
	step := (f.Xmax - f.Xmin) / float64(len(f.Y)-1)
	i := math.Floor(x / step)
	j := math.Ceil(x / step)
	if i == j {
		return f.Y[int(i)]
	}
	Xi := f.Xmin + (i * step)
	Xj := f.Xmin + (j * step)
	w := (x - Xi) / (Xj - Xi)
	return (1-w)*f.Y[int(i)] + w*f.Y[int(j)]
}

// Area returns the definite integral of the function on its domain X.
//
// Time complexity: O(N), where N is the number of points.
// Space complexity: O(1)
func (f FunctionUniform) Area() (area float64) {
	Y := f.Y
	step := (f.Xmax - f.Xmin) / float64(len(f.Y)-1)
	for i := 1; i < len(Y); i++ {
		Xi := f.Xmin + (float64(i) * step)
		Xi1 := f.Xmin + (float64(i-1) * step)
		area += (Xi - Xi1) * (Y[i] + Y[i-1]) / 2
	}
	return area
}

// AreaUpTo returns the definite integral of the function on its domain X intersected with [-Inf, x].
//
// Time complexity: O(N), where N is the number of points.
// Space complexity: O(1)
func (f FunctionUniform) AreaUpTo(x float64) (area float64) {
	Y := f.Y
	step := (f.Xmax - f.Xmin) / float64(len(f.Y)-1)
	for i := 1; i < len(Y); i++ {
		Xi := f.Xmin + (float64(i) * step)
		Xi1 := f.Xmin + (float64(i-1) * step)
		dX := Xi - Xi1
		if x < Xi {
			if x >= Xi1 {
				dxX := x - Xi1
				w := dxX / dX
				y := (1-w)*Y[i-1] + w*Y[i]
				area += dxX * (y + Y[i-1]) / 2
			}
			return area
		}
		area += dX * (Y[i] + Y[i-1]) / 2
	}
	return area
}
