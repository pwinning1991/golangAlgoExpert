package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleCycleCheck(t *testing.T) {
	input := []int{2, 3, 1, -4, -4, 2}
	output := HasSingleCycle(input)
	expected := true
	require.Equal(t, expected, output)
}
