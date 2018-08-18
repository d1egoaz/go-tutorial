package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u *user) notify() {
	fmt.Printf("user: %T: %+v\n", u, u)
	fmt.Printf("u.name: %T: %+v\n", u.name, u.name)
}

func main() {
	u := user{name: "diego"}
	sendNotification(&u)
}

// accepts values that implements the notifier interface
func sendNotification(n notifier) {
	n.notify()
}
