package main

import "fmt"

func main() {
	value := "111"
	fmt.Println(getDecodingCount(value, len(value)))
}

func getDecodingCount(message string, length int) int {
	countArray := make([]int, length+1)
	countArray[0] = 1
	countArray[1] = 1
	for i := 2; i <= length; i++ {
		countArray[i] = 0
		if message[i-1] > 48 {
			countArray[i] = countArray[i-1]
		}
		if message[i-2] == 49 || (message[i-2] == 50 || message[i-1] < 55) {
			countArray[i] += countArray[i-2]
		}
	}

	return countArray[length]
}
