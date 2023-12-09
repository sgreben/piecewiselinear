# piecewiselinear

[![](https://godoc.org/github.com/sgreben/piecewiselinear?status.svg)](http://godoc.org/github.com/sgreben/piecewiselinear) [![](https://goreportcard.com/badge/github.com/sgreben/piecewiselinear)](https://goreportcard.com/report/github.com/sgreben/piecewiselinear)

A tiny library for linear interpolation. `O(log(N))` per evaluation for `N` control points.

```go
import "github.com/sgreben/piecewiselinear"
```

- [Get it](#get-it)
- [Use it](#use-it)
  - [Fast special case](#fast-special-case)
- [Benchmarks](#benchmarks)


## Get it

```sh
go get -u "github.com/sgreben/piecewiselinear"
```

## Use it

```go
import "github.com/sgreben/piecewiselinear"

func main() {
    f := piecewiselinear.Function{Y:[]float64{0,1,0}} // range: "hat" function
    f.X = []float64{0, 0.5, 1}                        // domain: equidistant points along X axis
    fmt.Println(
		f.At(0),      // f.At(x) evaluates f at x
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
```

### Fast special case

If the control points are uniformly spaced, `piecewiselinear.FunctionUniform` is much faster (no search required):

```go
import "github.com/sgreben/piecewiselinear"

func main() {
	f := piecewiselinear.FunctionUniform{Y: []float64{0, 1, 0}} // range: "hat" function
	f.Xmin, f.Xmax = 0, 1                                       // domain: equidistant points along X axis
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
```

## Benchmarks

On an Apple M1 Pro:

- **6ns** per evaluation (`.At(x)`) for 10 control points
- **320ns** per evaluation for 10 million control points.

and, for `FunctionUniform`, **2ns** per evaluation regardless of the number of control points.

```
goos: darwin
goarch: arm64
pkg: github.com/sgreben/piecewiselinear
BenchmarkAt4-10                 230890022                5.499 ns/op           0 B/op          0 allocs/op
BenchmarkAt8-10                 199668106                6.084 ns/op           0 B/op          0 allocs/op
BenchmarkAt10-10                192352903                6.206 ns/op           0 B/op          0 allocs/op
BenchmarkAt100-10               138742411                8.613 ns/op           0 B/op          0 allocs/op
BenchmarkAt1k-10                46360660                25.50 ns/op            0 B/op          0 allocs/op
BenchmarkAt10k-10               16649996                70.02 ns/op            0 B/op          0 allocs/op
BenchmarkAt100k-10              11696936               100.4 ns/op             0 B/op          0 allocs/op
BenchmarkAt1M-10                 8512652               140.6 ns/op             0 B/op          0 allocs/op
BenchmarkAt10M-10                3769648               320.4 ns/op             0 B/op          0 allocs/op
BenchmarkUniformAt10M-10        571224222                2.185 ns/op           0 B/op          0 allocs/op
```
