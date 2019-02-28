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
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
