package main

import "fmt"

type User struct {
	name  string
	email string
}

func main() {
	stayOnStack()
	escapeToHeap()
	fmt.Printf("end\n")

}
func stayOnStack() User {
	u := createUser()
	return u
}

func escapeToHeap() *User {
	u := createUser()
	return &u
}

func createUser() User {
	u := User{
		name:  "name",
		email: "email",
	}
	return u
}
