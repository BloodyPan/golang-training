package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// var result [] int
	numDict := make(map[int]int)
	for index, value := range nums {
		numIndex, ok := numDict[value]
		if ok {
			// result = append(result, numIndex)
			// result = append(result, index)
			return []int{numIndex, index}
		}
		numDict[target-value] = index
	}
	return []int{}
}

func main() {
	// var nums [] int = []int{2, 7, 11, 15}
	// var nums [4] int = [4]int{2, 7, 11, 15}
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
