package main

import "fmt"

func main() {
	array := []int{10, 15, 3, 7}
	fmt.Println(checkIfElementPairCanAddTo(array, 11))
}

func checkIfElementPairCanAddTo(array []int, sum int) bool {
	existingItemsMap := make(map[int]int)
	for i, ithElement := range array {
		_, isAvailable := existingItemsMap[sum-ithElement]
		if isAvailable {
			return true
		}
		existingItemsMap[ithElement] = i
	}
	return false
}
