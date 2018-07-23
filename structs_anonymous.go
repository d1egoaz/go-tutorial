package main

import (
	"fmt"
)

func main() {
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
