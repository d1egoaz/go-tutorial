package main

import "fmt"

type structure struct {
	a bool
	b int16
	c float32
}

var s structure

func main() {
	// shadows global
	var s structure
	{
		// shadows local to main
		var s *structure
		if s != nil {
			fmt.Printf("s internal: %T: %+v\n", s, s)
		}
	}
	// ./variables_scopes.go:18:4: no new variables on left side of :=
	// cannot name it to an already defined variable
	// s := f1(&s)
	if err := f1(&s); err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("s: %T: %+v\n", s, s)
	fmt.Printf("&s: %T: %+v\n", &s, &s)
}

func f1(x *structure) error {
	x.a = true
	return nil
}
