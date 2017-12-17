package main

import (
	"fmt"
	"math/big"
)

func main() {
	perIteration := 1
	max := 100000
	c := makeSlices(perIteration, max)
	c2 := computeFactorial(c)
	c3 := combineFactorials(c2)
	answer := <-c3
	fmt.Println(answer.Text('e', 10))
}

func makeSlices(perIteration int, max int) chan []int {
	out := make(chan []int)
	go func() {
		for i := 1; i <= max; i += perIteration {
			var slice []int
			for j := i; j <= i+perIteration-1; j++ {
				slice = append(slice, j)
			}
			out <- slice
		}
		close(out)
	}()
	return out
}

func computeFactorial(c chan []int) chan big.Float {
	out := make(chan big.Float)
	go func() {
		for array := range c {
			factorial := new(big.Float).SetFloat64(1)
			for _, item := range array {
				floatItem := new(big.Float).SetFloat64(float64(item))
				factorial.Set(factorial.Mul(factorial, floatItem))
			}
			out <- *factorial
		}
		close(out)
	}()
	return out
}

func combineFactorials(c chan big.Float) chan big.Float {
	out := make(chan big.Float)
	go func() {
		sum := new(big.Float).SetFloat64(1)
		for i := range c {
			sum.Set(sum.Mul(sum, &i))
		}
		out <- *sum
		close(out)
	}()
	return out
}
