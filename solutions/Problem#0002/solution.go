package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(getProductArray(array, len(array)))
}

func getProductArray(array []int, length int) []int {
	var result = make([]int, length)
	temp := 1
	for i, ithElement := range array {
		result[i] = temp
		temp *= ithElement
	}

	temp = 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= temp
		temp *= array[i]
	}
	return result
}
