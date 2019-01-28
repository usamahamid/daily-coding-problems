package main

import "fmt"

func main() {
	array := []int{3, 4, -1, 1, 2}
	fmt.Println(getLowestMissingPositiveNumberInArray(array, len(array)))
}

func getLowestMissingPositiveNumberInArray(array []int, length int) (bool, int) {
	var result = make([]int, length)
	for _, ithElement := range array {
		if ithElement > 0 && ithElement < length {
			result[ithElement-1] = 1
		}
	}

	fmt.Println(result)
	for i, ithElement := range result {
		if ithElement != 1 {
			return true, i + 1
		}
	}
	return false, -1
}
