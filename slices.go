package main

import "fmt"

func main() {

	var s []string
	fmt.Printf("s: %T: %s %d %d %p\n", s, s, cap(s), len(s), &s)

	s2 := []string{}
	fmt.Printf("s2: %T: %s %d %d %p\n", s2, s2, cap(s2), len(s2), &s2)

	fruits := make([]string, 3)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"

	fmt.Println(fruits)
	inspectSlice(fruits)
}

func inspectSlice(slice []string) {
	fmt.Printf("%T: %d %d\n", slice, cap(slice), len(slice))
	for i, s := range slice {
		fmt.Printf("[%d], %p %s\n", i, &slice[i], s)
	}
}
