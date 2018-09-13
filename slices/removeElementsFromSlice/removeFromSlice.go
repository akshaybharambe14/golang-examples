package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	p := []int{9, 7, 0, 2}
	fmt.Println("Result: ", deleteElements(a, p))
}

// removes elements present at given positions from provided slice
func deleteElements(s []int, p []int) []int {
	p = sortAndRemoveDuplicate(p)
	deleteCnt := 0
	for i := range p {
		n, d := deleteElement(s, p[i]-deleteCnt)
		fmt.Println("n: ", n)
		if d {
			deleteCnt++
			s = n
		}
	}
	return s
}

// removes single element from slice, returns new slice and true if element is removed
func deleteElement(s []int, p int) ([]int, bool) {
	if p < 0 || p >= len(s) {
		return s, false
	}
	return append(s[:p], s[p+1:]...), true
}

func sortAndRemoveDuplicate(s []int) []int {
	sort.Ints(s)
	n := make([]int, len(s))
	fmt.Println(len(n))
	last, cnt := 0, 0
	for i := range s {
		if i == 0 && !(s[i] < 0) {
			last = s[i]
			n[cnt] = s[i]
			cnt++
			continue
		}

		if s[i] == last || s[i] < 0 {
			continue
		}

		last = s[i]
		n[cnt] = s[i]
		cnt++
	}

	return n[:cnt]
}
