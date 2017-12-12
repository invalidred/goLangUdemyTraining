package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("AbdulKhan"[2])

	mySlice = append(mySlice, "g", "h")
	fmt.Println(mySlice[5:])

	newSlice := append([]string{"x", "y", "z"}, mySlice...)
	fmt.Println(newSlice[:3])
	fmt.Println("new slice capacity", cap(newSlice))
	fmt.Println("my slice capacity", cap(mySlice))

	// deleting c from mySlice
	newSlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(newSlice)
}
