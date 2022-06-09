package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeLargest(t *testing.T) {
	expected := []int{18, 141, 541}
	output := ThreeLargest([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7})
	require.Equal(t, expected, output)
}
