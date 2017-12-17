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

	i := []int{1, 3, 5, 0, 6, 10, 2}
	fmt.Println(i)
	sort.Ints(i)
	fmt.Println(i)
	sort.Sort(sort.Reverse(sort.IntSlice(i)))
	fmt.Println(i)

	studyGroup := people{"Zeno", "John", "Al", "Jenny", "Abdul"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}
