package main

import (
	"fmt"
)

func main() {
	x := 10
	// displays the value and address of x
	fmt.Println("before", x, &x)

	// pass address of x
	increment2(&x)
	fmt.Println("after ", x, &x)
}

func increment2(inc *int) {
	// increment the * value that the pointer points to
	*inc++
	fmt.Println("inc   ", *inc, &inc, inc)
}
