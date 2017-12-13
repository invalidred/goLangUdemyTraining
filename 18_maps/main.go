package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 2
	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)

	v, ok := m["k2"]
	fmt.Println("ok?:", ok, v)

	// You can also declare and init a new map in the same line
	myMap := map[string]int{"Abdul": 1, "Khan": 2}
	fmt.Println(myMap)
}
