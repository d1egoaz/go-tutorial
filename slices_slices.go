package main

import "fmt"

func main() {
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Lemon"

	println("f1")
	// f1 and f11 share same addresses
	f1 := make([]string, len(fruits), 8)
	copy(f1, fruits)
	f11 := f1[2:4]
	inspectSlice(f1)
	inspectSlice(f11)

	// modifies shared address
	f11[0] = "CHANGED"
	inspectSlice(f1)
	inspectSlice(f11)

	println("\n\n\n", "f2")
	// f2 and f22 share same addresses
	f2 := make([]string, len(fruits))
	copy(f2, fruits)
	f22 := f2[2:4]
	// NOTE: this could give 1 extra capacity for a single element
	// same as f33 := f3[2:4:5]
	// adds element to shared capacity
	f22 = append(f22, "NEW")
	inspectSlice(f2)
	inspectSlice(f22)

	println("\n\n\n", "f3")
	// f3 and f33 share same addresses
	f3 := make([]string, len(fruits))
	copy(f3, fruits)
	f33 := f3[2:4:4]
	inspectSlice(f33)
	// adds element, however, as capacity == length, it'll create a copy value
	// of the slice on different memory address
	f33 = append(f33, "NEW on new copy")
	inspectSlice(f3)
	// NOTE: see how addresses have changed
	inspectSlice(f33)
}

func inspectSlice(slice []string) {
	fmt.Printf("%T: len: %d cap: %d\n", slice, len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d], %p %s\n", i, &slice[i], s)
	}
}
