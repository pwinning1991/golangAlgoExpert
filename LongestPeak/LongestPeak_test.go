package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPeak(t *testing.T) {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	actual := LongestPeak(array)
	require.Equal(t, 6, actual)
}
