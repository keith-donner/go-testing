package main

import (
	"fmt"
	"or/list"
)

func TrySlice() {
	// Example usage with an int slice

	// Example usage with a string slice
	intSlice := []int{1, 2, 3}

	newIntSlice := list.AddToSliceAtIndex(intSlice, 1, 4)

	newIntSlice2 := list.ExtractSlice(newIntSlice)
	list.AddToSliceAtIndex(&newIntSlice2, 2, 5)

	fmt.Println("Updated int slice:", newIntSlice) // Output: [1 99 2 3]

	stringSlice := []string{"a", "b", "c"}
	list.AddToSliceAtIndex(&stringSlice, 2, "z")
	fmt.Println("Updated string slice:", stringSlice) // Output: [a b z c]
}
