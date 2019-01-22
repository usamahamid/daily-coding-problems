package main

import "fmt"

func main() {
	array := []int{10, 15, 3, 7}
	fmt.Println(checkIfElementPairCanAddTo(array, 17))
}

func checkIfElementPairCanAddTo(array []int, sum int) (bool, int, int) {
	existingItemsMap := make(map[int]int)
	for i, ithElement := range array {
		j, isAvailable := existingItemsMap[sum-ithElement]
		if isAvailable {
			return true, j, i
		}
		existingItemsMap[ithElement] = i
	}
	return false, 0, 0
}
