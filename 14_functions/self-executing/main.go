package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hello Self Executing World")
	}()
}
git