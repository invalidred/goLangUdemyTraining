package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("this is true")
	}

	if !false {
		fmt.Println("this is not false")
	}

	if false {
		fmt.Println("this line is ignored")
	} else if !true {
		fmt.Println("this line is also ignored")
	} else {
		fmt.Println("this line is printed")
	}

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
