package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNonRepeatingCharacter(t *testing.T) {
	input := "aabc"
	output := 2
	got := NonRepeatingCharacter(input)
	require.Equal(t, output, got)
}
