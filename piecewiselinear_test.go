package piecewiselinear

import (
	"fmt"
	"math/rand"
	"testing"
)

func Example() {
	f := Function{Y: []float64{0, 1, 0}} // range: "hat" function
	f.X = []float64{0, 0.5, 1}           // domain: equidistant points along X axis
	fmt.Println(
		f.At(0), // f.At(x) evaluates f at x
		f.At(0.25),
		f.At(0.5),
		f.At(0.75),
		f.At(1.0),
		f.At(123.0),  // outside its domain X the function is constant 0
		f.At(-123.0), //

		f.Area(),
		f.AreaUpTo(0.5),
	)
	// Output:
	// 0 0.5 1 0.5 0 0 0 0.5 0.25
}

func ExampleUniform() {
	f := FunctionUniform{Y: []float64{0, 1, 0}} // range: "hat" function
	f.Xmin, f.Xmax = 0, 1                       // domain: equidistant points along X axis
	fmt.Println(
		f.At(0), // f.At(x) evaluates f at x
		f.At(0.25),
		f.At(0.5),
		f.At(0.75),
		f.At(1.0),
		f.At(123.0),  // outside its domain X the function is constant 0
		f.At(-123.0), //

		f.Area(),
		f.AreaUpTo(0.5),
	)
	// Output:
	// 0 0.5 1 0.5 0 0 0 0.5 0.25
}

func TestFunction_AreaUpTo(t *testing.T) {
	tests := []struct {
		name string
		X    []float64
		Y    []float64
		x    float64
		want float64
	}{
		{
			name: "simple",
			X:    []float64{0, 1},
			Y:    []float64{0, 1},
			x:    1.0,
			want: 0.5,
		},
		{
			name: "simple+1",
			X:    []float64{0, 1},
			Y:    []float64{1, 2},
			x:    0.5,
			want: 0.625,
		},
		{
			name: "three segments (1)",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			x:    2,
			want: 3.5,
		},
		{
			name: "three segments (2)",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			x:    2.5,
			want: 4.625,
		},
		{
			name: "three segments (3)",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			x:    3,
			want: 6,
		},
		{
			name: "three segments (4, outside domain)",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			x:    123,
			want: 6,
		},
		{
			name: "three segments (6, outside domain)",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			x:    -123,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.AreaUpTo(tt.x); got != tt.want {
				t.Errorf("Function.AreaUpTo(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

func TestFunction_Area(t *testing.T) {
	tests := []struct {
		name string
		X    []float64
		Y    []float64
		want float64
	}{
		{
			name: "empty",
			want: 0.0,
		},
		{
			name: "simple",
			X:    []float64{0, 1},
			Y:    []float64{0, 1},
			want: 0.5,
		},
		{
			name: "simple+1",
			X:    []float64{0, 1},
			Y:    []float64{1, 2},
			want: 1.5,
		},
		{
			name: "two segments",
			X:    []float64{0, 1, 2},
			Y:    []float64{1, 2, 2},
			want: 3.5,
		},
		{
			name: "three segments",
			X:    []float64{0, 1, 2, 3},
			Y:    []float64{1, 2, 2, 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.Area(); got != tt.want {
				t.Errorf("Function.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_At(t *testing.T) {
	tests := []struct {
		name string
		X    []float64
		Y    []float64
		x    float64
		want float64
	}{
		{
			name: "simple",
			X:    []float64{0, 1},
			Y:    []float64{0, 1},
			x:    0.5,
			want: 0.5,
		},
		{
			name: "simpleNonzeroBoundaries",
			X:    []float64{10, 20},
			Y:    []float64{123, 456},
			x:    10,
			want: 123,
		},
		{
			name: "simpleNonzeroBoundaries",
			X:    []float64{10, 20},
			Y:    []float64{123, 456},
			x:    20,
			want: 456,
		},
		{
			name: "simpleNonzeroBoundaries",
			X:    []float64{10, 20},
			Y:    []float64{123, 456},
			x:    21,
			want: 0,
		},
		{
			name: "saw(0.25)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.25,
			want: -1,
		},
		{
			name: "saw(0.125)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.125,
			want: -0.5,
		},
		{
			name: "saw(1.0)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    1.0,
			want: 0.0,
		},
		{
			name: "saw(0.0)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.0,
			want: 0.0,
		},
		{
			name: "saw(2.5) (outside domain)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    2.5,
			want: 0.0,
		},
		{
			name: "saw(-2.5) (outside domain)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    -2.5,
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.At(tt.x); got != tt.want {
				t.Errorf("Function.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_IsInterpolatedAt(t *testing.T) {
	tests := []struct {
		name           string
		X              []float64
		Y              []float64
		x              float64
		isInterpolated bool
	}{
		{
			name:           "simple",
			X:              []float64{0, 1},
			Y:              []float64{0, 1},
			x:              0.5,
			isInterpolated: true,
		},
		{
			name:           "simpleNonzeroBoundaries",
			X:              []float64{10, 20},
			Y:              []float64{123, 456},
			x:              10,
			isInterpolated: true,
		},
		{
			name:           "simpleNonzeroBoundaries",
			X:              []float64{10, 20},
			Y:              []float64{123, 456},
			x:              20,
			isInterpolated: true,
		},
		{
			name:           "simpleNonzeroBoundaries",
			X:              []float64{10, 20},
			Y:              []float64{123, 456},
			x:              21,
			isInterpolated: false,
		},
		{
			name:           "saw(0.25)",
			X:              []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:              []float64{0, -1, 0, 1, 0},
			x:              0.25,
			isInterpolated: true,
		},
		{
			name:           "saw(0.125)",
			X:              []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:              []float64{0, -1, 0, 1, 0},
			x:              0.125,
			isInterpolated: true,
		},
		{
			name:           "saw(1.0)",
			X:              []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:              []float64{0, -1, 0, 1, 0},
			x:              1.0,
			isInterpolated: true,
		},
		{
			name:           "saw(0.0)",
			X:              []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:              []float64{0, -1, 0, 1, 0},
			x:              0.0,
			isInterpolated: true,
		},
		{
			name:           "saw(2.5) (outside domain)",
			X:              []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:              []float64{0, -1, 0, 1, 0},
			x:              2.5,
			isInterpolated: false,
		},
		{
			name:           "saw(-2.5) (outside domain)",
			X:              []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:              []float64{0, -1, 0, 1, 0},
			x:              -2.5,
			isInterpolated: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.IsInterpolatedAt(tt.x); got != tt.isInterpolated {
				t.Errorf("Function.IsInterpolated() = %v, want %v", got, tt.isInterpolated)
			}
		})
	}
}

func TestUniform(t *testing.T) {
	var f Function
	var fU FunctionUniform
	ks := []int{2, 3, 4, 5, 8, 9, 16, 17, 32, 33, 64, 65, 128, 129, 256, 257}
	for _, k := range ks {
		f.Y = make([]float64, k)
		fU.Y = f.Y
		fU.Xmax = 1
		fU.Xmin = 0
		for i := range f.Y {
			f.Y[i] = rand.Float64()
		}
		f.X = Span(0, 1, len(f.Y))
		xs := make([]float64, len(f.X))
		for i := range f.X {
			xs[i] = rand.Float64()
		}
		for _, x := range xs {
			fx := f.At(x)
			fUx := fU.At(x)
			if fx != fUx {
				t.Errorf(".At(): %v != %v (expected)", fUx, fx)
			}
			Fx := f.AreaUpTo(x)
			FUx := fU.AreaUpTo(x)
			if Fx != FUx {
				t.Errorf(".AreaUpTo(): %v != %v (expected)", FUx, Fx)
			}
		}
		F := f.Area()
		FU := fU.Area()
		if F != FU {
			t.Errorf(".Area(): %v != %v (expected)", F, FU)
		}
	}
}

func benchmarkWithKPoints(b *testing.B, k int) {
	var f Function
	f.Y = make([]float64, k)
	for i := range f.Y {
		f.Y[i] = rand.Float64()
	}
	f.X = Span(0, 1, len(f.Y))
	xs := make([]float64, len(f.X))
	for i := range f.X {
		xs[i] = rand.Float64()
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = f.At(xs[n%len(f.X)])
	}
}

func benchmarkUniformWithKPoints(b *testing.B, k int) {
	var f FunctionUniform
	f.Y = make([]float64, k)
	for i := range f.Y {
		f.Y[i] = rand.Float64()
	}
	xs := make([]float64, len(f.Y))
	for i := range f.Y {
		xs[i] = rand.Float64()
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = f.At(xs[n%len(f.Y)])
	}
}

func BenchmarkAt4(b *testing.B) {
	benchmarkWithKPoints(b, 4)
}

func BenchmarkAt8(b *testing.B) {
	benchmarkWithKPoints(b, 8)
}

func BenchmarkAt10(b *testing.B) {
	benchmarkWithKPoints(b, 10)
}

func BenchmarkAt100(b *testing.B) {
	benchmarkWithKPoints(b, 100)
}

func BenchmarkAt1k(b *testing.B) {
	benchmarkWithKPoints(b, 1_000)
}

func BenchmarkAt10k(b *testing.B) {
	benchmarkWithKPoints(b, 10_000)
}

func BenchmarkAt100k(b *testing.B) {
	benchmarkWithKPoints(b, 100_000)
}

func BenchmarkAt1M(b *testing.B) {
	benchmarkWithKPoints(b, 1_000_000)
}

func BenchmarkAt10M(b *testing.B) {
	benchmarkWithKPoints(b, 10_000_000)
}

func BenchmarkUniformAt10M(b *testing.B) {
	benchmarkUniformWithKPoints(b, 10_000_000)
}
