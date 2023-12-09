package piecewiselinear

import "math"

// Span generates `nPoints` equidistant points spanning [min,max]
func Span(min, max float64, nPoints int) []float64 {
	X := make([]float64, nPoints)
	min, max = math.Min(max, min), math.Max(max, min)
	step := (max - min) / float64(nPoints-1)
	for i := range X {
		X[i] = min + float64(i)*step
	}
	return X
}
