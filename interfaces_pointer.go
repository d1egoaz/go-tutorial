package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
}

// NOTE: pointer semantics works just for pointers!!!!!
// if written with value semantics, it could have been called with pointer or values!!!!
// if func (u user) notify() {
func (u *user) notify() {
	fmt.Printf("user: %T: %+v\n", u, u)
	fmt.Printf("u.name: %T: %+v\n", u.name, u.name)
}

func main() {
	u := user{name: "diego"}
	sendNotification(&u)

	// ./interfaces_pointer.go:22:18: cannot use u (type user) as type notifier in argument to sendNotification:
	// user does not implement notifier (notify method has pointer receiver)
	// sendNotification(u)
}

// NOTE: polymorphic method on notifier interface
// accepts values that implements the notifier interface
func sendNotification(n notifier) {
	n.notify()
}
