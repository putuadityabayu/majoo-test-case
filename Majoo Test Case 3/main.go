package main

import (
	"fmt"
)

func main() {
	var a1 int
	var a2 int
	var x int

	// Get input
	fmt.Println("Inputan deret pertama:")
	fmt.Scan(&a1)

	fmt.Println("Inputan deret kedua:")
	fmt.Scan(&a2)

	fmt.Println("Value X:")
	fmt.Scan(&x)

	// Calculate diff
	deret := a2 - a1

	// Calculate next integer
	data_str := fmt.Sprintf("%d,%d", a1, a2)
	temp := a2
	for i := 2; i < x; i++ {
		next := temp + deret
		data_str += fmt.Sprintf(",%d", next)
		temp = next
	}

	fmt.Println(data_str)
}
