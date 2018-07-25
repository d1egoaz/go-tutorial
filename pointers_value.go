package main

import (
	"fmt"
)

func main() {
	x := 10
	// displays the value and address of x
	fmt.Println("before", x, &x)
	increment(x)
	fmt.Println("after ", x, &x)
}

// increment ...
func increment(inc int) {
	inc++
	fmt.Println("inc   ", inc, &inc)
}
