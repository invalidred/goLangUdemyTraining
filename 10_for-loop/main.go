package main

import (
	"fmt"
)

func main() {
	// simple loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// nested loop
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}
	//while loop using for
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// print first 50 odd numbers
	j := 0
	for {
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
		if j >= 50 {
			break
		}
	}
	// print  rune
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
