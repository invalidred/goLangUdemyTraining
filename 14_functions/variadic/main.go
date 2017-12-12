package main

import "fmt"

func main() {
	values := []float64{43, 53, 34, 34, 11, 59}
	avg := average(1, 13.23, 343.34, 43.11)
	avg2 := average(values...)
	fmt.Println(avg)
	fmt.Println(avg2)
}

func average(numbers ...float64) float64 {
	fmt.Println(numbers)
	fmt.Printf("%T \n", numbers)
	var sum float64
	for _, value := range numbers {
		sum += value
	}
	return sum / float64(len(numbers))
}
