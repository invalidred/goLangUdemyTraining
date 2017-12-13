package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Create slice of slice of string to hold slices of words
	buckets := make([][]string, 12)
	// Create slices to hold words words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
}

// HashBucket func, given a string a number of
// buckets the func returns which bucket the word should go to
func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// comment out the above then comment out the below
	// a more uneven distribution
	// return len(word) % buckets
}
