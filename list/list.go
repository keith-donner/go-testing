package list

import (
	"sort"
)

/*
listToAddTo: Pass in the list of any type
value: insert this value into that list
*/

// func AddToMap(listToAddTo *[]any, key string, value any) {
// 	mapType := (reflect.TypeOf(listToAddTo).Elem())

// 	if mapType == "string" {
// 		newMap

// 	}

// 	listToAddTo[key] = value

// }

/*
Pass in the list of any type
value: this value will be deleted from the list
*/
// RemoveFrom removes a value from a slice or a key from a map.
// It uses reflection to handle different types of data structures.
// RemoveValue removes the first occurrence of the value from the slice.
// It returns the modified slice and a boolean indicating if a value was removed.
func RemoveFromSlice[T comparable](slice []T, value T) ([]T, bool) {
	var indexToRemove int
	removed := false

	for i, v := range slice {
		if v == value {
			indexToRemove = i
			removed = true
			break
		}
	}

	if !removed {
		return slice, false
	}

	// Remove the element
	return append(slice[:indexToRemove], slice[indexToRemove+1:]...), true
}

// RemoveKey removes the key-value pair with the specified key from the map.
// It returns the modified map.
func RemoveFromMap[K comparable, V any](m map[K]V, key K) map[K]V {
	delete(m, key)
	return m
}

// send in list of any type and whether it should be Ascending (true) A-Z or Descending (false) Z-A

// Sortable is an interface for types that can be sorted.
// Interface that enforces the type T to be sortable
type Comparable interface {
	~int | ~float64 | ~string // Add other comparable types if needed
}

// Sort sorts a slice of any comparable type in ascending or descending order based on the reverse flag.
func Sort[T Comparable](slice []T, reverse bool) {
	// Create a custom sort.Interface implementation
	sortFunc := func(i, j int) bool {
		if reverse {
			return slice[i] > slice[j]
		}
		return slice[i] < slice[j]
	}

	// Sort the slice using sort.Slice
	sort.Slice(slice, sortFunc)
}
