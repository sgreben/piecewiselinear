# piecewiselinear

[![](https://godoc.org/github.com/sgreben/piecewiselinear?status.svg)](http://godoc.org/github.com/sgreben/piecewiselinear) [![](https://goreportcard.com/badge/github.com/sgreben/piecewiselinear/goreportcard)](https://goreportcard.com/report/github.com/sgreben/piecewiselinear) [![gocover.io](https://gocover.io/_badge/github.com/sgreben/piecewiselinear)](https://gocover.io/github.com/sgreben/piecewiselinear) [![Build](https://github.com/sgreben/piecewiselinear/workflows/Build/badge.svg)](https://github.com/sgreben/piecewiselinear/actions?query=workflow%3ABuild)

A tiny library for linear interpolation. `O(log(N))` per evaluation for `N` control points.

```go
import "github.com/sgreben/piecewiselinear"
```

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
