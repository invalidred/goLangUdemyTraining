package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) fullName() string {
	return p.First + p.Last
}

func main() {
	p1 := person{"Abdul", "Khan", 25}
	p2 := person{"Sara", "Abdo", 24}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
