package main

import (
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	sort.Ints(array)
	firstNum := 0
	secondNum := len(array) - 1
	for firstNum < secondNum {
		sum := array[firstNum] + array[secondNum]
		if sum == target {
			return []int{array[firstNum], array[secondNum]}
		} else if sum > target {
			secondNum -= 1

		} else {
			firstNum += 1
		}
	}
	return []int{}
}
