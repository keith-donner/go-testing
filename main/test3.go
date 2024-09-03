package main

import (
	"fmt"
	"or/list"
)

func TestNow() {

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	list.AddToMap(m, "test", 2)
	fmt.Print(m)

	fmt.Print(m)

}
