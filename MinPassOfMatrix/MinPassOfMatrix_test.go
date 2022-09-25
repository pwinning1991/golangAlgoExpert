package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinPassOfMatrix(t *testing.T) {
	input := [][]int{
		{0, -1, -3, 2, 0},
		{1, -2, -5, -1, -3},
		{3, 0, 0, -4, -1},
	}
	expected := 3
	actual := MinPassOfMatrix(input)
	require.Equal(t, expected, actual)
}
