package main

import "fmt"

func main() {
	var data []string
	fmt.Println("data - dynamicCap", len(data), cap(data))
	dynamicCap(data)

	// NOTE: cap already allocated, lenght 0
	data2 := make([]string, 0, 1e5)
	fmt.Println("data2 - dynamicCap", len(data2), cap(data2))
	dynamicCap(data2)

	// NOTE: cap already allocated, lenght 1e5
	data3 := make([]string, 1e5)
	fmt.Println("data3 - dynamicCapFixed", len(data3), cap(data3))
	inspectSlice(data3)
	dynamicCapFixed(data3)
}

func dynamicCap(data []string) {
	lastCap := float64(cap(data))

	// 100k
	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		newCap := float64(cap(data))
		if lastCap != newCap {
			capChange := (newCap - lastCap) / lastCap * 100
			lastCap = newCap
			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChange)
		}
	}
}

func dynamicCapFixed(data []string) {
	lastCap := float64(cap(data))

	// 100k
	for record := 0; record < 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data[record] = value

		newCap := float64(cap(data))
		if lastCap != newCap {
			capChange := (newCap - lastCap) / lastCap * 100
			lastCap = newCap
			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChange)
		}
	}
}
