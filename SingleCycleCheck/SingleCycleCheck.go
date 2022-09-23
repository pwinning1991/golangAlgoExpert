package main

func HasSingleCycle(array []int) bool {
	elementsVisited := 0
	currentIdx := 0
	for elementsVisited < len(array) {
		if elementsVisited > 0 && currentIdx == 0 {
			return false
		}
		elementsVisited += 1
		currentIdx = getNextIdx(currentIdx, array)
	}
	return currentIdx == 0
}

func getNextIdx(currentIdx int, array []int) int {
	jump := array[currentIdx]
	nextIdx := (currentIdx + jump) % len(array)
	if nextIdx < 0 {
		return nextIdx + len(array)
	}
	return nextIdx
}
