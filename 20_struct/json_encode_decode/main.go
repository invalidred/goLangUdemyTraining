package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Person struct
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)

	var p2 Person
	rdr := strings.NewReader(`{"First": "James", "Last":"Bond", "Age": 20}`)
	json.NewDecoder(rdr).Decode(&p2)

	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
}
