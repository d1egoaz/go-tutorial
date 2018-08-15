package main

import (
	"fmt"
)

type data struct {
	i int
}

func main() {
	d := data{i: 1}
	d.showData()

	// get a copy of the value, as it's a value receiver
	f1 := d.showData

	f1()
	d.i = 100
	f1()

	// as d.setI uses pointers semantics, f2 points to the reference not a copy@
	f2 := d.setI

	f2(5)
	d.showData()
	f2(6)
	d.showData()
}

func (u data) showData() {
	fmt.Printf("u: %T: %+v\n", u, u)
}

func (u *data) setI(i int) {
	u.i = i
}
