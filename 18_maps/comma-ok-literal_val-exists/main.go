package main

import (
	"fmt"
)

func main() {
	myGreetings := map[int]string{
		0: "Good Morning",
		1: "Salam Alaikum",
		2: "Namastae",
	}

	fmt.Println(myGreetings)
	if val, exists := myGreetings[2]; exists {
		delete(myGreetings, 2)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("The value does not exists")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreetings)
}
