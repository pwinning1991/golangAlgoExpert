package main

func ArrayOfProducts1(array []int) []int {
	products := make([]int, len(array))
	for i := range array {
		runningproduct := 1
		for j := range array {
			if i != j {
				runningproduct *= array[j]
			}
		}
		products[i] = runningproduct
	}

	return products
}

func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))
	leftProducts := make([]int, len(array))
	rightProducts := make([]int, len(array))
	for i := range array {
		products[i] = 1
		leftProducts[i] = 1
		rightProducts[i] = 1
	}
	leftRunningProduct := 1
	for i := range array {
		leftProducts[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}
	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		rightProducts[i] = rightRunningProduct
		rightRunningProduct *= array[i]
	}

	for i := range array {
		products[i] = leftProducts[i] * rightProducts[i]
	}
	return products
}
