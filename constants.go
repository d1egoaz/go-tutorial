package main

import "fmt"
import "math"

func main() {

	// untyped constans
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating point

	fmt.Printf("ui: %T: %+v\n", ui, ui)
	fmt.Printf("uf: %T: %+v\n", uf, uf)

	const ti int = 12345       // type: int
	const tf float64 = 3.14159 // type: float64

	// ./constants.go:18:8: constant 1000 overflows uint8
	// const myuint uint8 = 1000

	// kind promotion, type vs kind, const == kind

	// type float64
	x := 3 * 0.333
	fmt.Printf("x: %T: %+v\n", x, x)

	// kind of floating point
	const y = 1 / 3.0
	fmt.Printf("y: %T: %+v\n", y, y)

	// kind integer
	const zero = 1 / 3
	fmt.Printf("zero: %T: %+v\n", zero, zero)

	const (
		maxInt = math.MaxInt64
		bigger = maxInt * maxInt
		// biggerInt int64 = bigger
		// won't compile
		// ./constants.go:38:3: constant 85070591730234615847396907784232501249 overflows int64
		// it'll work with bigger type
		biggerFloat float64 = bigger
	)

	// cannot be used directly as
	// ./constants.go:47:5: constant 85070591730234615847396907784232501249 overflows int
	// xx := integer

	// can be used with cast
	xx := float64(bigger)
	// or specifying the right type
	var yy float64 = bigger

	fmt.Printf("maxInt: %T: %+v\n", maxInt, maxInt)
	fmt.Printf("xx-bigger: %T: %+v\n", xx, xx)
	fmt.Printf("yy-bigger: %T: %+v\n", yy, yy)

}
