package main

import (
	"fmt"
)

func main() {
	array := []int{7, 5, 10, 100, 10, 5}
	fmt.Println(getMaximumSumOfNonAdjacentElements(array, len(array)))
}

func getMaximumSumOfNonAdjacentElements(array []int, len int) int {
	if len == 0 {
		return 0
	}
	if len == 1 {
		return array[0]
	}
	incl := array[0]
	excl := 0
	newExcl := 0
	for i := 1; i < len; i++ {
		newExcl = max(incl, excl)
		incl = excl + array[i]
		excl = newExcl
	}
	return max(incl, excl)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
