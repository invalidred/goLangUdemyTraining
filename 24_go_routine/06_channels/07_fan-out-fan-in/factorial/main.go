package main

import (
	"fmt"
	"math/big"
	"sync"
)

func main() {
	perIteration := 1000
	max := 1000000
	c := makeSlices(perIteration, max)
	c2 := computeFactorial(c)
	c3 := computeFactorial(c)
	c4 := computeFactorial(c)
	c5 := computeFactorial(c)
	c6 := computeFactorial(c)
	c7 := computeFactorial(c)
	c8 := merge(c2, c3, c4, c5, c6, c7)
	answer := <-c8
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

func merge(channels ...chan big.Float) chan big.Float {
	var wg sync.WaitGroup
	var sumWait sync.WaitGroup
	wg.Add(len(channels))
	sumWait.Add(1)

	out := make(chan big.Float)
	sumCh := make(chan big.Float)

	for _, channel := range channels {
		go func(input chan big.Float, sumOut chan big.Float) {
			sum := new(big.Float).SetFloat64(1)
			for i := range input {
				sum.Set(sum.Mul(sum, &i))
			}
			sumOut <- *sum
			wg.Done()
		}(channel, sumCh)
	}

	go func() {
		total := new(big.Float).SetFloat64(1)
		for sum := range sumCh {
			total.Set(total.Mul(total, &sum))
		}
		out <- *total
		sumWait.Done()
	}()

	go func() {
		wg.Wait()
		close(sumCh)
		sumWait.Wait()
		close(out)
	}()

	return out
}
