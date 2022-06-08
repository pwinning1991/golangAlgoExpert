package main

func IsValidSubsequence(array []int, sequence []int) bool {
	currentNum := 0
	for _, num := range array {
		if num == sequence[currentNum] {

			currentNum += 1
			if currentNum > len(sequence)-1 {
				return true
			}

		} else {
			continue
		}

	}
	return false
}
