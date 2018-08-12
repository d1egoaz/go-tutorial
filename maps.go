package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	users := make(map[string]user)
	users["u1"] = user{"a", 1}
	users["u2"] = user{"b", 2}
	users["u3"] = user{"c", 3}

	for k, v := range users {
		fmt.Printf("[%s] v: %T: %+v\n", k, v, v)
	}

	users2 := map[string]user{
		"u1": {"a", 1},
		"u2": {"b", 2},
		"u3": {"c", 3},
	}

	for k, v := range users2 {
		fmt.Printf("[%s] v: %T: %+v\n", k, v, v)
	}

	u, found := users["u4"]
	if !found {
		fmt.Println("user u4 not found")
	}

	u2, found2 := users["u2"]
	if found2 {
		fmt.Printf("u2: %T: %+v\n", u2, u2)
	}

	// zero value
	fmt.Printf("u: %T: %+v\n", u, u)

	delete(users, "u1")
	// no-op
	delete(users, "u1")

}
