package main

// This takes in a matrix of 1s and 0s
// 1s rep[resent a river, find the lens of all rivers
// rivers can bend into Ls
func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			sizes = traverseNode(i, j, matrix, sizes, visited)
		}
	}

	return sizes
}

func traverseNode(i, j int, matrix [][]int, sizes []int, visited [][]bool) []int {
	currentRiverSize := 0
	riverNodes := [][]int{{i, j}}
	for len(riverNodes) > 0 {
		currentNode := riverNodes[0]
		riverNodes = riverNodes[1:]
		i, j := currentNode[0], currentNode[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}
		currentRiverSize++
		unvisitedNeighbors := getUnvisitedNeighbors(i, j, matrix, visited)
		riverNodes = append(riverNodes, unvisitedNeighbors...)
	}
	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}

	return sizes
}

func getUnvisitedNeighbors(i, j int, matrix [][]int, visited [][]bool) [][]int {
	neighbors := [][]int{}
	if i > 0 && !visited[i-1][j] {
		neighbors = append(neighbors, []int{i - 1, j})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		neighbors = append(neighbors, []int{i + 1, j})
	}
	if j > 0 && !visited[i][j-1] {
		neighbors = append(neighbors, []int{i, j - 1})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {
		neighbors = append(neighbors, []int{i, j + 1})
	}
	return neighbors
}
