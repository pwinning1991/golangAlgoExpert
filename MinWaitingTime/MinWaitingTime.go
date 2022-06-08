package main

import "sort"

func MinWaitingTime(queries []int) int {
	sort.Ints(queries)
	totalTime := 0
	for idx, num := range queries {
		queriesLeft := len(queries) - (idx + 1)
		totalTime += num * queriesLeft
	}

	return totalTime

}
