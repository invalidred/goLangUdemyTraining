package main

import (
	"fmt"
)

func changeMap(val map[string]int) {
	val["Abdul"] = 100
}

func changeSlice(val []string) {
	val[0] = "Abdul"
}

func changeValue(val *int) {
	*val = 20
}

func main() {
	a := 10
	changeValue(&a)
	fmt.Println(a) // 20

	b := []string{"1", "2", "3"}
	c := make([]string, 1, 25) // make slice with lenth upto 25 elements
	changeSlice(b)
	changeSlice(c)

	fmt.Println(b) // [Abdul 2, 3]
	fmt.Println(c) // [Abdul]

	d := make(map[string]int)
	changeMap(d)
	fmt.Println(d)
}
