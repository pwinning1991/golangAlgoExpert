package main

func MoveElementToEnd(array []int, toMove int) []int {
	left, right := 0, len(array)-1
	for left < right {
		if array[right] == toMove {
			right -= 1
		} else if array[left] != toMove {
			left += 1
		} else {
			array[right], array[left] = array[left], array[right]
		}

	}

	return array

}
