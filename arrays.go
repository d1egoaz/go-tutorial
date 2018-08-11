package main

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "apple0"
	fruits[1] = "apple1"
	fruits[2] = "apple2"
	fruits[3] = "apple3"
	fruits[4] = "apple4"

	fmt.Printf("fruits: %T: %+v\n", fruits, fruits)

	for _, f := range fruits {
		fmt.Println(f)
	}

	for i, f := range fruits {
		fmt.Println(i, f)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(i, fruits[i])
	}

	numbers := [4]int{1, 2, 3, 4}
	fmt.Printf("numbers: %T: %+v\n", numbers, numbers)
}
