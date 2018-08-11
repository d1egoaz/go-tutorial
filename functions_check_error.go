package main

import "fmt"

type user struct {
	a bool
}

func main() {
	u, error := retrieveUser("diego")
	if error != nil {
		fmt.Printf("error on retrieve user: %T: %+v\n", error, error)
		return
	}
	fmt.Printf("u: %T: %+v\n", *u, *u)

	if _, error := updateUser(u); error != nil {
		fmt.Printf("error on update user: %T: %+v\n", error, error)
		return
	}
	fmt.Printf("u: %T: %+v\n", u, u)
}

func retrieveUser(name string) (*user, error) {
	fmt.Printf("name: %T: %+v\n", name, name)

	u := user{
		a: true,
	}

	return &u, nil
}

func updateUser(u *user) (string, error) {
	u.a = false
	return "ok", nil
}
