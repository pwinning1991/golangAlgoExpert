package main

import (
	"math"
	"sort"
)

func SmallestDiff(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	pointer1, pointer2 := 0, 0
	pair := []int{}
	smallest := math.MaxInt32
	currentDiff := math.MaxInt32
	for pointer1 < len(array1) && pointer2 < len(array2) {
		first, second := array1[pointer1], array2[pointer2]
		if first == second {
			return []int{first, second}
		} else if first < second {
			currentDiff = second - first
			pointer1 += 1

		} else if second < first {
			currentDiff = first - second
			pointer2 += 1
		}
		if smallest > currentDiff {
			smallest = currentDiff
			pair = []int{first, second}
		}

	}

	return pair

}
