package main

import (
	"fmt"
)

func wrapper() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	fmt.Printf("yo \n")
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	increment2 := wrapper()
	fmt.Println(increment2())
	fmt.Println(increment2())

	anonymousIncrementWrapper := func() func() int {
		counter := 0
		return func() int {
			counter++
			return counter
		}
	}

	anonymousIncrement := anonymousIncrementWrapper()
	fmt.Println(anonymousIncrement())
	fmt.Println(anonymousIncrement())
}
