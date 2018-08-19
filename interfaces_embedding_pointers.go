package main

import "fmt"

type user struct {
	name string
}

type admin2 struct {
	user  // NOTE: embedding
	level string
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("sending email to: %T: %+v\n", u, u)
}
func (u *admin2) notify() {
	fmt.Printf("ADMIN, level: %s, sending email to: %T: %+v\n", u.level, u, u)
}

func main() {
	adm2 := admin2{
		user: user{
			name: "Diego",
		},
		level: "superuser",
	}

	// ./interfaces_embedding_pointers.go:30:18: cannot use adm2 (type admin2) as type notifier in argument to sendNotification:
	// admin2 does not implement notifier (notify method has pointer receiver)
	// sendNotification(adm2)

	sendNotification(&adm2)
	sendNotification(&adm2.user)
	// ADMIN, level: superuser, sending email to: *main.admin2: &{user:{name:Diego} level:superuser}
	// sending email to: *main.user: &{name:Diego}
}

func sendNotification(n notifier) {
	n.notify()
}
