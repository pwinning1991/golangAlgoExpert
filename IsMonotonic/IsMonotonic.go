package main

func IsMonotonic(array []int) bool {
	isNonDecreasing := true
	isNonIncreasing := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonDecreasing = false
		}
		if array[i] > array[i-1] {
			isNonIncreasing = false
		}
	}

	return isNonDecreasing || isNonIncreasing
}
