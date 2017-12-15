package main

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	First             string
	Last              string
	Age               int `json:"wisdom score"`
	notMarshalledFied string
}

func main() {
	p1 := Person{"Abdul", "Khan", 32, "wont be json'd"}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))

	var p2 Person
	bsa := []byte(`{"First":"James", "Last": "Bond", "wisdom score": 20}`)
	json.Unmarshal(bsa, &p2)
	fmt.Println(p2)
}
