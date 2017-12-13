package main

import (
	"fmt"
)

func main() {
	myGreetings := map[int]string{
		1: "Good Morning",
		2: "Namastay",
		3: "Salam Alaikum",
	}

	for key, value := range myGreetings {
		fmt.Printf("key: %v - value: %v \n", key, value)
	}
}
