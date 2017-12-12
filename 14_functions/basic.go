package main

import (
	"fmt"
)

// function params v/s args
func getName(my, your string) string {
	return fmt.Sprint(my, your)
}

// named returns BAD MAKES CODE UNREADABLE
func getNameNamedReturn(my string, your string) (abdul string) {
	abdul = fmt.Sprint(my, your)
	return
}

// multiple return values
func getNameMultipleReturn(my string, your string) (string, string) {
	return fmt.Sprint(my, your), fmt.Sprint(your, my)
}

func main() {
	fmt.Println(getName("abdul", "khan"))
	fmt.Println(getNameNamedReturn("Deen", "Khan"))
	fmt.Println(getNameMultipleReturn("sara", "abdo"))
}
