package main

import "fmt"

func main() {
	x := 1
	error := f1(&x)
	if error != nil {
		return
	}
	fmt.Printf("x: %T: %+v\n", x, x)
	fmt.Printf("x: %T: %+v\n", &x, &x)
}

func f1(x *int) error {
	*x = 3
	return nil
}
