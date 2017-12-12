package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	// defers the execution of world() right before the main function exits
	defer world()
	hello()
}
