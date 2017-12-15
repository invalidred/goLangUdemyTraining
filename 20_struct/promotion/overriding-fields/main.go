package main
package main

import (
	"fmt"
)

// Person stuff
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero stuff
type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Not Really",
		LicenseToKill: false,
	}

	fmt.Println(p1.Person.First, p1.First)
	fmt.Println(p2.Person.First, p2.First)
}
