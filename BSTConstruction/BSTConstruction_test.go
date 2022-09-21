package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func TestBSTConstruction(t *testing.T) {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	root.Insert(12)
	require.True(t, root.Right.Left.Left.Value == 12)

	root.Remove(10)
	require.True(t, root.Contains(10) == false)
	require.True(t, root.Value == 12)

	require.True(t, root.Contains(15))
}
