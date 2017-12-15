package main

import (
	"fmt"
)

// Person struct
type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"Abdul", "Khan", 12}
	p2 := Person{"Sara", "Abdul", 32}
	fmt.Println(p1)
	fmt.Println(p2)
}
