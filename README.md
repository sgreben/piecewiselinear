# piecewiselinear

[![](https://godoc.org/github.com/sgreben/piecewiselinear?status.svg)](http://godoc.org/github.com/sgreben/piecewiselinear) [![](https://goreportcard.com/badge/github.com/sgreben/piecewiselinear)](https://goreportcard.com/report/github.com/sgreben/piecewiselinear)

A tiny library for linear interpolation. `O(log(N))` per evaluation for `N` control points.

```go
import "github.com/sgreben/piecewiselinear"
```

- [Get it](#get-it)
- [Use it](#use-it)
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
    f.X = piecewiselinear.Span(0, 1, len(f.Y)) // domain: equidistant points along X axis
    fmt.Println(
		f.At(0), // f.At(x) evaluates f at x
		f.At(0.25),
		f.At(0.5),
		f.At(0.75),
		f.At(1.0),
		f.At(123.0),  // outside its domain X the function is constant 0
		f.At(-123.0), //
	)
    // Output:
    // 0 0.5 1 0.5 0 0 0
}
```

## Benchmarks

On an Apple M1 Pro:

- **6ns** per evaluation (`.At(x)`) for 10 control points
- **320ns** per evaluation for 10 million control points.

```
goos: darwin
goarch: arm64
pkg: github.com/sgreben/piecewiselinear
BenchmarkAt4-10         217302646                5.461 ns/op           0 B/op          0 allocs/op
BenchmarkAt8-10         197175420                6.048 ns/op           0 B/op          0 allocs/op
BenchmarkAt10-10        188384818                6.283 ns/op           0 B/op          0 allocs/op
BenchmarkAt100-10       138276301                9.086 ns/op           0 B/op          0 allocs/op
BenchmarkAt1k-10        41258203                25.18 ns/op            0 B/op          0 allocs/op
BenchmarkAt10k-10       16852758                69.99 ns/op            0 B/op          0 allocs/op
BenchmarkAt100k-10      11745384               100.5 ns/op             0 B/op          0 allocs/op
BenchmarkAt1M-10         8501438               143.0 ns/op             0 B/op          0 allocs/op
BenchmarkAt10M-10        3659188               319.5 ns/op             0 B/op          0 allocs/op
```
