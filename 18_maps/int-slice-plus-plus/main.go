package main

import (
	"fmt"
)

func main() {
	a := make([]int, 1)
	fmt.Println(a)
	a[0] = 1
	fmt.Println(a)
	a[0]++
	fmt.Println(a)
}
