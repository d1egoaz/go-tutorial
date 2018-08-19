package main

import "fmt"

type user struct {
	name string
}

type admin struct {
	person user // this is NOT embedding
	level  string
}

type admin2 struct {
	user  // NOTE: embedding
	level string
}

type notifier interface {
	notify()
}

func (u user) notify() {
	fmt.Printf("Admin: %T: %+v\n", u, u)
}

func main() {

	adm := admin{
		person: user{
			name: "Diego",
		},
		level: "superuser",
	}

	fmt.Printf("adm: %T: %+v\n", adm, adm)

	adm.person.notify()
	fmt.Println("")

	adm2 := admin2{
		user: user{
			name: "Diego",
		},
		level: "superuser",
	}

	fmt.Printf("adm2: %T: %+v\n", adm2, adm2)

	// we can access the inner type's method directly
	adm2.user.notify()

	// the inner type's method is promoted
	adm2.notify()
}
