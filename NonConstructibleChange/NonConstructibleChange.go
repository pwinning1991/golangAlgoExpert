package main

import "sort"

func NonConstrucibleChange(coins []int) int {
	totalSum := 0
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > totalSum+1 {
			return totalSum + 1
		}
		totalSum += coin
	}

	return totalSum + 1
}
