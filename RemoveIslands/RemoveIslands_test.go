package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveIslands(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}
	expected := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}
	actual := RemoveIslands(input)
	require.Equal(t, expected, actual)

}
