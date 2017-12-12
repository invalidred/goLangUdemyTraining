package main

import (
	"fmt"
)

func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func greatestNumber(numbers ...int) int {
	var greatest int
	for _, number := range numbers {
		if number > greatest {
			greatest = number
		}
	}
	return greatest
}

func main() {
	// 1 & 2
	half, isEven := half(4)
	halfFuncExpression := func(a int) (float64, bool) {
		return float64(a) / 2, a%2 == 0
	}
	fmt.Println(half, isEven)
	fmt.Println(halfFuncExpression(4))

	// 3
	numbers := []int{1, 3, 100, 4, 5}
	fmt.Println(greatestNumber(numbers...))
}
