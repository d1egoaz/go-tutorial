package main

import "fmt"

func main() {
	a := [4]string{"a", "b", "c", "d"}
	for i, v := range a {
		fmt.Printf("v: %T: %+v\n", v, v)
		// %s	the uninterpreted bytes of the string or slice
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &a[i])

		// NOTE: internal v is the copy of element value
		// different address pointing to same value

		// ---
		// v: string: a
		// Value[a]	Address[0xc42000e1e0] IndexAddr[0xc42005a040]
		// v: string: b
		// Value[b]	Address[0xc42000e1e0] IndexAddr[0xc42005a050]
		// v: string: c
		// Value[c]	Address[0xc42000e1e0] IndexAddr[0xc42005a060]
		// v: string: d
		// Value[d]	Address[0xc42000e1e0] IndexAddr[0xc42005a070]
	}
}
