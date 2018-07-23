package main

import (
	"fmt"
)

type s1 struct {
	a bool
}

type s2 struct {
	a bool
}

func main() {
	var a s1
	var b s2

	// no implicit conversion
	// b = a
	// ERROR cannot use a (type s1) as type s2 in assignment
	b = s2(a)

	fmt.Println(a, b)
	fmt.Printf("a: %T: %+v\n", a, a)
	fmt.Printf("b: %T: %+v\n", b, b)

	// when anonymous, there is implicit conversion
	z := struct {
		a bool
	}{
		a: true,
	}

	// there is no cast here
	a = z
	fmt.Printf("a: %T: %+v\n", a, a)
}
