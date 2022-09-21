package main

import "sort"

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	sortedIntervals := make([][]int, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i][0] < sortedIntervals[j][0]
	})

	mergedIntervals := make([][]int, 0)
	currentInterval := sortedIntervals[0]
	mergedIntervals = append(mergedIntervals, currentInterval)

	for _, nextInterval := range sortedIntervals {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntercalEnd := nextInterval[0], nextInterval[1]

		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = max(currentIntervalEnd, nextIntercalEnd)
		} else {
			currentInterval = nextInterval
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
