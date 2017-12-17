package main

import "fmt"
import "strconv"

func main() {
	var xInt = 12
	var yFloat64 = 12.123450

	fmt.Println(yFloat64 + float64(xInt))

	var xRune = 'a'
	var yInt32 = 'b' // rune is an alias for int32
	fmt.Println(xRune)
	fmt.Println(yInt32)
	fmt.Println(string(xRune))
	fmt.Println(string(yInt32))

	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))
	fmt.Println([]byte("hello"))

	var xASCII = "12"
	var yInt = 6
	zInt, _ := strconv.Atoi(xASCII)
	fmt.Println(yInt + zInt)

	yASCII := "I have this many: " + strconv.Itoa(yInt)
	fmt.Println(yASCII)
}
