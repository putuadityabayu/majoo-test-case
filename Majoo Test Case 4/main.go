package main

import (
	"fmt"
	"sort"
)

func AscendingSort(a []float64) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

func DescendingSort(a []float64) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
}

// Ascending sorting
//
// a = float64 array
func ManualAscending(a []float64) {
	var n = len(a)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}

// Descending sorting
//
// a = float64 array
func ManualDescending(a []float64) {
	var n = len(a)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if a[j] > a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}

func main() {
	a := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	AscendingSort(a)
	fmt.Printf("Built in Ascending Sorting: %v\n", a)

	b := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	DescendingSort(b)
	fmt.Printf("Built in Descending Sorting: %v\n", b)

	c := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	ManualAscending(c)
	fmt.Printf("Manual Ascending Sorting: %v\n", c)

	d := []float64{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	ManualDescending(d)
	fmt.Printf("Manual Descending Sorting: %v", d)
}
