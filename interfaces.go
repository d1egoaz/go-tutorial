package main

import "fmt"

type reader interface {
	read(b []byte) (int, error)
}

type user struct {
	name string
}

// type user is a concrete type that implements reader with the value receiver
func (user) read(b []byte) (int, error) {
	s := "Test"
	return copy(b, s), nil
}

type pipe struct {
}

func (pipe) read(b []byte) (int, error) {
	s := "very long string"
	return copy(b, s), nil
}
func main() {
	user := user{name: "Diego"}
	b := make([]byte, 20)
	user.read(b)

	retrieve(user)
	p := pipe{}
	retrieve(p)
}

func retrieve(r reader) error {
	b := make([]byte, 100)
	len, err := r.read(b)
	if err != nil {
		return err
	}
	f2()
	fmt.Printf("len: %T: %+v\n", len, len)
	return nil
}

func f2() {
	return nil
}
