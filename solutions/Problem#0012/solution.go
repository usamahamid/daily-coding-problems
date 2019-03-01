package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 2}
	fmt.Println(getCountUtil(4, array))
}
func getCountUtil(n int, array []int) int {
	sort.Ints(array)
	return getCount(n, array)
}

func getCount(n int, array []int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for _, ithElement := range array {
		if ithElement <= n {
			count += getCount(n-ithElement, array)
		} else {
			break
		}
	}
	return count
}
