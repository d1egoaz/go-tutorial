package main

import "fmt"
import "strconv"

func main() {
	// iota start at 0 increment by 1
	const (
		A1 = iota //0
		A2 = iota //1
		A3 = iota //2
	)
	fmt.Println("A", A1, A2, A3)

	// not need to repeat iota
	const (
		B1 = iota //0
		B2        //1
		B3        //2
	)
	fmt.Println("B", B1, B2, B3)

	// starts with a fixed value
	const (
		C1 = iota + 10 //0
		C2             //1
		C3             //2
	)
	fmt.Println("C", C1, C2, C3)

	// starts with a fixed value
	const (
		D1 = 1 << iota //0
		D2             //1
		D3             //2
	)
	fmt.Println("D", fmtInt(D1), fmtInt(D2), fmtInt(D3))
	fmt.Println(33, fmtInt(33))
	fmt.Println(33, strconv.FormatInt(33, 2))

	fmt.Printf("%08b ", D1)
	fmt.Printf("%08b ", D2)
	fmt.Printf("%08b ", D3)

}

func fmtInt(n int64) string {
	return strconv.FormatInt(n, 2)
}
