package main

import (
	"fmt"
	"or/list"
)

func TestNow() {
	intSlice := []int{5, 3, 4, 1, 2}
	fmt.Println("Original int slice:", intSlice)

	list.Sort(intSlice, false)
	fmt.Println("Sorted int slice (ascending):", intSlice)

	list.Sort(intSlice, true)
	fmt.Println("Sorted int slice (descending):", intSlice)

	stringSlice := []string{"banana", "apple", "cherry"}
	fmt.Println("Original string slice:", stringSlice)

	list.Sort(stringSlice, false)
	fmt.Println("Sorted string slice (ascending):", stringSlice)

	list.Sort(stringSlice, true)
	fmt.Println("Sorted string slice (descending):", stringSlice)
}
