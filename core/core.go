package core

type Any interface{}

// AllMatch checks if all elements in the slice match the predicate
func AllMatch[T any](items []T, matchFunc func(T) bool) bool {
	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		if !matchFunc(item) {
			return false
		}
	}
	return true
}

func AnyMatch[T any](items []T, matchFunc func(T) bool) *T {
	for _, item := range items {
		if matchFunc(item) {
			return &item
		}
	}
	return nil
}

// AnyMatch checks if any element in a slice of Any objects matches the predicate
// func AnyMatch(elements []Any, predicate func(Any) bool) *Any {
// 	for _, element := range elements {
// 		if predicate(element) {
// 			return &element
// 		}
// 	}
// 	return nil
// }
