package main

import (
	"math"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	return tree.validateBstHelper(math.MinInt32, math.MaxInt32)

}

func (tree *BST) validateBstHelper(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !tree.Left.validateBstHelper(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validateBstHelper(tree.Value, max) {
		return false
	}
	return true

}
