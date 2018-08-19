package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

// Stringer is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
// type Stringer interface {
// 	String() string
// }
func (u *user) String() string {
	return fmt.Sprintf("My name is %s and my age is %v", u.name, u.age)
}

type printer interface {
	print() error
}

// func (u *user) ... works just for pointer receiver
func (u user) print() error {
	fmt.Println(u)
	return nil
}

func main() {

	u := user{name: "Diego", age: 10}

	// fmt.* if implements the String() interface on pointer semantic it would use that
	// otherwise will use the default toString

	// {Diego 10}
	fmt.Println(u)

	// My name is Diego and my age is 10
	fmt.Println(&u)

	list := []printer{
		// stores a copy of the user value in the interface value
		u,

		// stores a copy of the address of the user value interface
		&u,
	}

	u.name = "Diego CHANGED"
	u.age = 42

	for _, v := range list {
		v.print()
	}
	// {Diego 10}
	// My name is Diego and my age is 10
	// {Diego 10}
	// {Diego CHANGED 42}
}
