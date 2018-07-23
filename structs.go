package main

import (
	"fmt"
)

type example struct {
	a bool
	b int16
	c float32
}

func main() {
	var x example
	fmt.Printf("x: %T: %+v\n", x, x)

	y := example{
		a: true,
		b: 5,
		c: 3.14159,
	}

	fmt.Printf("y: %T: %+v\n", y, y)
	fmt.Printf("y.a: %T: %+v\n", y.a, y.a)
	fmt.Printf("y.b: %T: %+v\n", y.b, y.b)
	fmt.Printf("y.c: %T: %+v\n", y.c, y.c)

	// anonymous struct
	z := struct {
		a bool
		b int16
		c float32
	}{
		a: true,
		b: 5,
		c: 3.14159,
	}
	fmt.Printf("z: %T: %+v\n", z, z)

}
