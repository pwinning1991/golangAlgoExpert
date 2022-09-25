package main

type IntPair struct {
	First, Second int
}

func MinPassOfMatrix(matrix [][]int) int {
	passes := convertNegatives(matrix)
	if !containsNegative(matrix) {
		return passes - 1
	} else {
		return -1
	}
}

func convertNegatives(matrix [][]int) int {
	queue := getAllPositivePositions(matrix)
	var passes = 0
	for len(queue) > 0 {
		var currentSize = len(queue)

		for currentSize > 0 {
			nextElemet := queue[0]
			queue = queue[1:]
			currentRow, currentCol := nextElemet.First, nextElemet.Second

			adjacentPositions := getAdjacentPositions(currentRow, currentCol, matrix)
			for _, position := range adjacentPositions {
				row, col := position.First, position.Second

				value := matrix[row][col]
				if value < 0 {
					matrix[row][col] *= -1
					queue = append(queue, IntPair{row, col})
				}
			}

			currentSize -= 1
		}
		passes += 1

	}
	return passes

}

func getAllPositivePositions(matrix [][]int) []IntPair {
	postivePositions := make([]IntPair, 0)
	for row := range matrix {
		for col := range matrix[row] {
			value := matrix[row][col]
			if value > 0 {
				postivePositions = append(postivePositions, IntPair{row, col})
			}
		}
	}
	return postivePositions
}

func getAdjacentPositions(row, col int, matrix [][]int) []IntPair {
	neighbors := make([]IntPair, 0)
	rowLen := len(matrix)
	colLen := len(matrix[row])
	if row+1 < rowLen {
		neighbors = append(neighbors, IntPair{row + 1, col})
	}
	if row-1 >= 0 {
		neighbors = append(neighbors, IntPair{row - 1, col})
	}
	if col-1 >= 0 {
		neighbors = append(neighbors, IntPair{row, col - 1})
	}
	if col+1 < colLen {
		neighbors = append(neighbors, IntPair{row, col + 1})
	}
	return neighbors
}

func containsNegative(matrix [][]int) bool {
	for _, row := range matrix {
		for _, value := range row {
			if value < 0 {
				return true
			}
		}
	}
	return false
}
