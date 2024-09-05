package main

import (
	"fmt"
	"or/conv"
	"or/http"
)

func TrySlice() {
	// Example usage with an int slice

	// Example usage with a string slice

	newTime := conv.DateFromString("2024-2-5", "600p", "y/m/d")

	if newTime.Hour() == 18 {
		fmt.Print(newTime)

	}

	http.Post()
}
