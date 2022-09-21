package main

func FirstDuplicateValue1(array []int) int {
	vals := map[int]int{}

	for i := range array {
		if _, ok := vals[array[i]]; ok {
			return array[i]
		} else {
			vals[array[i]] = i
		}
	}
	return -1
}

func FirstDuplicateValue(array []int) int {
	for _, value := range array {
		absValue := abs(value)
		if array[absValue-1] < 0 {
			return absValue
		}
		array[absValue-1] *= -1
	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
