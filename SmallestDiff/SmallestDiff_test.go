package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallestDiff(t *testing.T) {
	expected := []int{28, 26}
	output := SmallestDiff([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	require.Equal(t, expected, output)
}
