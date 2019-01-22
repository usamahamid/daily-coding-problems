package main

import "fmt"

func main() {
	array := []int{10, 15, 3, 7}
	fmt.Println(checkIfElementPairCanAddTo(array, len(array), 10))
}

func checkIfElementPairCanAddTo(array []int, length int, sum int) bool {
	existingItemsMap := make(map[int]int)
	for i := 0; i < length; i++ {
		ithElement := array[i]
		_, isAvailable := existingItemsMap[sum - ithElement]
		if isAvailable {
			return true
		}
		existingItemsMap[ithElement] = i
	}
	return false
}
