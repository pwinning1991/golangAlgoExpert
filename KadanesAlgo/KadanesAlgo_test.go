package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKadandesAlgo(t *testing.T) {
	expected := 19
	output := KadanesAlgo([]int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4})
	require.Equal(t, expected, output)
}
