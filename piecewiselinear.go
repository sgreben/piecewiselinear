package piecewiselinear

// Function is a piecewise-linear 1-dimensional function
type Function struct {
	X []float64
	Y []float64
}

// At returns the function's value at the given point.
// Outside the domain, the function is constant at the respective boundary value.
//
// The function's X slice is expected to be sorted in ascending order. This sortedness property is _not_ verified.
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
	switch i {
	case 0:
		return Y[0]
	case len(X):
		return Y[len(X)-1]
	}
	w := (x - X[i-1]) / (X[i] - X[i-1])
	return (1-w)*Y[i-1] + w*Y[i]
}
