package main

import (
	"fmt"
)

const metersToYards = 1.09361

func main() {
	var meters float64
	fmt.Println("Enter meters you swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Printf("%v meter is %v yards \n", meters, yards)
}
