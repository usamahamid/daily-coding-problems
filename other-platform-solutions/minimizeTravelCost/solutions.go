package main

import (
	"fmt"
)

func main() {
	testCases := 0
	fmt.Scanf("%d", &testCases)
	for index := 0; index < testCases; index++ {
		preformTestCase()
	}
}

func preformTestCase() {
	arrayLength := 0
	fmt.Scanf("%d", &arrayLength)

	var distanceArray = make([]int, arrayLength)
	for index := 0; index < arrayLength; index++ {
		fmt.Scanf("%d", &distanceArray[index])
	}

	var costArray = make([]int, arrayLength)
	for index := 0; index < arrayLength; index++ {
		fmt.Scanf("%d", &costArray[index])
	}

	findMinimumCost(distanceArray, costArray, arrayLength)
}

func findMinimumCost(distanceArray []int, costArray []int, arrayLength int) int {
	for i := 0; i < arrayLength; i++ {

	}
	return 0
}
