package main

import (
	"fmt"
)

const p string = "Death & Taxes"
const q = 42

// PI for math
const (
	Pi       = 3.14
	Language = "G0"
)

// Memory Constants
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Printf("%b \t\t\t\t %d \n", KB, KB)
	fmt.Printf("%b \t\t\t %d \n", MB, MB)
	fmt.Printf("%b \t %d \n", GB, GB)
}
