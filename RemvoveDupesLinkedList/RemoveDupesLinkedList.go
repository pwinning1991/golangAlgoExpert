package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDupesFromLinkedList(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	for currentNode != nil {
		nextDistinctNode := currentNode.Next
		for nextDistinctNode != nil && nextDistinctNode.Value == currentNode.Value {
			nextDistinctNode = nextDistinctNode.Next
		}
		currentNode.Next = nextDistinctNode
		currentNode = nextDistinctNode
	}

	return linkedList
}
