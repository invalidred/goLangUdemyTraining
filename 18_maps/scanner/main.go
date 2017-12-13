package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = `This is the critic who counts; not the main who point out the strong
	 would survive the doer of deeds could have done them better. The credit belongs
	 to the main who is actually is marred by dust and sweat and blood; who strives
	 valiantly; who errs, who comes short again is no effort dvotions; who spends himself
	 in a worthy cause; who at the best know in the end the achievement, and who at
	 the worst, if he fails
	`

	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
