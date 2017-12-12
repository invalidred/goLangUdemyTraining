package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, number := range numbers {
		callback(number)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, number := range numbers {
		if callback(number) {
			xs = append(xs, number)
		}
	}
	return xs
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	visit(numbers, func(number int) {
		fmt.Println(number)
	})

	filteredNumbers := filter(numbers, func(number int) bool {
		return number > 4
	})

	fmt.Println(filteredNumbers)
}
