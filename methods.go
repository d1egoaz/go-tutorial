package main

import (
	"fmt"
	"strconv"
)

type user struct {
	name string
}

func main() {

	fmt.Println("u")
	u := user{name: "diego"}
	fmt.Printf("user main: %T: %+v %p\n", u, u, &u)
	u.showName()
	// user.showName(u) COMPILER DOES THIS!

	// compiler changes u to a pointer &u so it can use the pointer receiver method
	u.changeName("diego2")
	(&u).changeName("diego3")
	// (*user).changeName(&u, "...") COMPILER DOES THIS!
	u.showName()

	fmt.Println("\nup")
	up := &user{name: "nathy"}
	fmt.Printf("user main: %T: %+v %p\n", up, up, up)

	// compilers changes up to *up so it uses the value the pointer points to
	up.showName()
	up.changeName("nathy2")
	(*up).showName()

	users := []user{
		{name: "u1"},
		{name: "u2"},
	}

	// range makes a copy of the value on v
	for i, v := range users {
		v.changeName(strconv.Itoa(i))
		// showName makes another copy of the value
		v.showName()
	}
}

// NOTE: this receives a copy of the value, see the different memory address
// user main: main.user: {name:diego} 0xc42000e1e0
// user showName: main.user: {name:diego} 0xc42000e210
// works for user and *user
//
// Pointers of a type can also be used to call methods with a value receiver
func (u user) showName() { // value receiver
	fmt.Printf("user showName: %T: %+v %p\n", u, u, &u)
}

// Works on pointer receiver
// pointer receiver, escapes to the heap
// NOTE: no for built in types
func (u *user) changeName(name string) {
	u.name = name
}
