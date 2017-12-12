package main

import "fmt"

func makeInc() func() int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}

func main() {
	increment := makeInc()
	fmt.Println(increment())
	fmt.Println(increment())
}
