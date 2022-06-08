package main

import "sort"

func SortedSquaredArray(array []int) []int {
	squaredNums := make([]int, len(array))
	for idx, num := range array {
		squaredNumber := num * num
		squaredNums[idx] = squaredNumber

	}
	sort.Ints(squaredNums)
	return squaredNums
}
