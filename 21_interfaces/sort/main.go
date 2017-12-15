package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[j] > p[i] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny", "Abdul"}

	sort.Strings(s)
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	//studyGroup := people{"Zeno", "John", "Al", "Jenny", "Abdul"}
	// sort.Sort(studyGroup)
	// fmt.Println(studyGroup)
	// fmt.Println(sort.Reverse(studyGroup))

	// i := []int{1, 3, 5, 0, 6, 10, 2}
	// // fmt.Println(sort.Ints(i))
	// fmt.Println(sort.Reverse(sort.IntSlice(i)))
}
