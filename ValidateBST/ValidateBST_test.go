package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func TestValidateBST(t *testing.T) {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	output := root.ValidateBst()
	require.True(t, output)

}
