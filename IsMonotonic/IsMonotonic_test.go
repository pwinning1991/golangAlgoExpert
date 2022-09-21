package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsMonotonic(t *testing.T) {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	actual := IsMonotonic(array)
	require.True(t, actual)
}
