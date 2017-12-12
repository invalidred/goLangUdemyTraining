package main

import (
	"fmt"
)

func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	g := func() {
		fmt.Println("hello world")
	}
	g()

	f := makeGreeter()
	fmt.Println(f())
}
