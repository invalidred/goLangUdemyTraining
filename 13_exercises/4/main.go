package main

import (
	"fmt"
)

//Create a program that prints to the terminal asking for a user to enter a small
// number and a larger number. Print the remainder of the larger number divided
// by the smaller number.
func main() {
	var smallNumber int
	var largeNumber int
	fmt.Println("Enter small number")
	fmt.Scan(&smallNumber)
	fmt.Println("Enter large number")
	fmt.Scan(&largeNumber)
	fmt.Println("The remainder is", largeNumber%smallNumber)
}
