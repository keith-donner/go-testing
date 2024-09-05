package list

import (
	"fmt"
	"reflect"
	"sort"
)

/*
listToAddTo: Pass in the list of any type
value: insert this value into that list
*/
/////////// --------------- ADD TO LIST /////////// ---------------

/*
if you want to update the existing list, send it in as a pointer with a & in front

otherwise you will need to include a return list
*/
func AddToMap(list interface{}, key interface{}, value interface{}) interface{} {
	v := reflect.ValueOf(list)

	// Handle the case where the input is a pointer to a map
	if v.Kind() == reflect.Ptr {
		// Dereference the pointer
		mapValue := v.Elem()
		if mapValue.Kind() != reflect.Map {
			panic("provided pointer does not point to a map")
		}
		// Set the key-value pair
		keyValue := reflect.ValueOf(key)
		valueValue := reflect.ValueOf(value)
		mapValue.SetMapIndex(keyValue, valueValue)
		return nil
	}

	// Handle the case where the input is a map (not a pointer)
	if v.Kind() == reflect.Map {
		// Create a new map of the same type
		mapType := v.Type()
		newMap := reflect.MakeMap(mapType)
		// Copy existing entries
		for _, key := range v.MapKeys() {
			newMap.SetMapIndex(key, v.MapIndex(key))
		}
		// Add the new key-value pair
		keyValue := reflect.ValueOf(key)
		valueValue := reflect.ValueOf(value)
		newMap.SetMapIndex(keyValue, valueValue)
		return newMap.Interface()
	}

	panic("provided value is not a map or pointer to a map")
}

/*
if you want to update the existing list, send it in as a pointer with a & in front

otherwise you will need to include a return list
*/
func AddToSlice(slice interface{}, value interface{}) interface{} {
	// Handle the case where the input is a pointer to a slice
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() == reflect.Ptr {
		sliceElem := sliceValue.Elem()
		if sliceElem.Kind() != reflect.Slice {
			fmt.Errorf("provided pointer does not point to a slice")
			return nil
		}

		// Create a new slice with an additional element
		sliceType := sliceElem.Type()
		newSlice := reflect.MakeSlice(sliceType, sliceElem.Len()+1, sliceElem.Cap()+1)
		reflect.Copy(newSlice, sliceElem)

		// Add the new value to the new slice
		newSlice.Index(sliceElem.Len()).Set(reflect.ValueOf(value))

		// Update the original slice pointer to point to the new slice
		sliceElem.Set(newSlice)

		return nil
	}

	// Handle the case where the input is a slice (not a pointer)
	if sliceValue.Kind() == reflect.Slice {
		sliceType := sliceValue.Type()
		newSlice := reflect.MakeSlice(sliceType, sliceValue.Len()+1, sliceValue.Cap()+1)
		reflect.Copy(newSlice, sliceValue)

		// Add the new value to the new slice
		newSlice.Index(sliceValue.Len()).Set(reflect.ValueOf(value))

		return newSlice.Interface()
	}
	fmt.Errorf("provided value is not a slice or pointer to a slice")
	return nil
}

/*
if you want to update the existing list, send it in as a pointer with a & in front

otherwise you will need to include a return list
*/
func AddToSliceAtIndex[T any](slice interface{}, index int, value T) []any {
	v := reflect.ValueOf(slice)
	if v.Kind() == reflect.Ptr {
		// Handle pointer to a slice
		sliceValue := v.Elem()
		if sliceValue.Kind() != reflect.Slice {
			panic("provided pointer does not point to a slice")
		}
		if index < 0 || index > sliceValue.Len() {
			panic("index out of bounds")
		}

		// Create a new slice with one more element than the original slice
		newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Cap()+1)

		// Copy elements before the index
		reflect.Copy(newSlice, sliceValue.Slice(0, index))

		// Insert the new value
		newSlice.Index(index).Set(reflect.ValueOf(value))

		// Copy elements after the index
		reflect.Copy(newSlice.Slice(index+1, newSlice.Len()), sliceValue.Slice(index, sliceValue.Len()))

		// Set the updated slice back to the pointer
		sliceValue.Set(newSlice)
		return nil
	}

	if v.Kind() == reflect.Slice {
		// Handle slice directly
		sliceValue := v
		if index < 0 || index > sliceValue.Len() {
			panic("index out of bounds")
		}

		// Create a new slice with one more element than the original slice
		newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Cap()+1)

		// Copy elements before the index
		reflect.Copy(newSlice, sliceValue.Slice(0, index))

		// Insert the new value
		newSlice.Index(index).Set(reflect.ValueOf(value))

		// Copy elements after the index
		reflect.Copy(newSlice.Slice(index+1, newSlice.Len()), sliceValue.Slice(index, sliceValue.Len()))
	}

	panic("provided value is not a slice or pointer to a slice")
}

/////////// -------------------------------------- ///////////
/////////// -------------------------------------- ///////////

/*
Pass in the list of any type
value: this value will be deleted from the list
It returns the modified slice and a boolean indicating if a value was removed.
*/
func RemoveFromSlice(slice interface{}, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	val := reflect.ValueOf(value)

	// Handle the case where the input is a pointer to a slice
	if v.Kind() == reflect.Ptr {
		sliceValue := v.Elem()
		if sliceValue.Kind() != reflect.Slice {
			panic("provided pointer does not point to a slice")
		}

		var indexToRemove int
		removed := false

		for i := 0; i < sliceValue.Len(); i++ {
			if reflect.DeepEqual(sliceValue.Index(i).Interface(), val.Interface()) {
				indexToRemove = i
				removed = true
				break
			}
		}

		if removed {
			// Remove the element
			newSlice := reflect.AppendSlice(sliceValue.Slice(0, indexToRemove), sliceValue.Slice(indexToRemove+1, sliceValue.Len()))
			sliceValue.Set(newSlice)
		}

		return nil
	}

	// Handle the case where the input is a slice (not a pointer)
	if v.Kind() == reflect.Slice {
		var indexToRemove int
		removed := false

		for i := 0; i < v.Len(); i++ {
			if reflect.DeepEqual(v.Index(i).Interface(), val.Interface()) {
				indexToRemove = i
				removed = true
				break
			}
		}

		if removed {
			// Return a new slice with the element removed
			return reflect.AppendSlice(v.Slice(0, indexToRemove), v.Slice(indexToRemove+1, v.Len())).Interface()
		}

		return v.Interface()
	}

	panic("provided value is not a slice or pointer to a slice")
}

// ///////// --------------- REMOVE FROM LIST /////////// ---------------
// RemoveKey removes the key-value pair with the specified key from the map.
// It returns the modified map.
func RemoveFromMap(list interface{}, key interface{}) interface{} {
	v := reflect.ValueOf(list)
	k := reflect.ValueOf(key)

	// Handle the case where the input is a pointer to a map
	if v.Kind() == reflect.Ptr {
		mapValue := v.Elem()
		if mapValue.Kind() != reflect.Map {
			panic("provided pointer does not point to a map")
		}

		// Delete the key from the map
		mapValue.SetMapIndex(k, reflect.Value{})
		return nil
	}

	// Handle the case where the input is a map (not a pointer)
	if v.Kind() == reflect.Map {
		// Create a new map of the same type
		mapType := v.Type()
		newMap := reflect.MakeMap(mapType)

		// Copy existing entries except the one to remove
		for _, mapKey := range v.MapKeys() {
			if !reflect.DeepEqual(mapKey.Interface(), key) {
				newMap.SetMapIndex(mapKey, v.MapIndex(mapKey))
			}
		}

		return newMap.Interface()
	}

	panic("provided value is not a map or pointer to a map")
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

/*
make sure to include the map or slice return type

int slice: mergedSlice := Merge(slice1, slice2).([]int)

string slice: mergedSlice := Merge(slice1, slice2).([]string)

string map: mergedMap := Merge(map1, map2).(map[string]string)

int map: mergedMap := Merge(map1, map2).(map[string]int)
*/
func Merge(v1, v2 interface{}) interface{} {
	if reflect.TypeOf(v1) != reflect.TypeOf(v2) {
		panic("Cannot merge values of different types")
	}

	switch v1.(type) {
	case []int:
		slice1 := v1.([]int)
		slice2 := v2.([]int)
		return append(slice1, slice2...)
	case []string:
		slice1 := v1.([]string)
		slice2 := v2.([]string)
		return append(slice1, slice2...)
	case map[interface{}]interface{}:
		map1 := v1.(map[interface{}]interface{})
		map2 := v2.(map[interface{}]interface{})
		for key, value := range map2 {
			map1[key] = value
		}
		return map1
	default:
		panic("Unsupported type for merging")
	}
}
