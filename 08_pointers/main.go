package main

import (
	"fmt"
)

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xaddress

	var b = &a
	fmt.Println(b)  // 0xaddress
	fmt.Println(*b) //43

	*b = 42
	fmt.Println(a) // 42

	zero(&a)
	fmt.Println(a) // 0
}
