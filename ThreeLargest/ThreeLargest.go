package main

import "math"

func ThreeLargest(array []int) []int {
	threeNums := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for _, num := range array {
		updateLargest(threeNums, num)
	}

	return threeNums

}

func updateLargest(threeNums []int, num int) {
	if num > threeNums[2] {
		shiftAndUpdate(threeNums, num, 2)
	} else if num > threeNums[1] {
		shiftAndUpdate(threeNums, num, 1)
	} else if num > threeNums[0] {
		shiftAndUpdate(threeNums, num, 0)
	}
}

func shiftAndUpdate(array []int, num int, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}

}
