package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello, my name is Abdul Khan")
	fmt.Println("Enter you name")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v \n", name)
}
