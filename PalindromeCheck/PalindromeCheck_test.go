package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindromeCheck(t *testing.T) {
	expected := true
	output := PalindromeCheck("abcdcba")
	require.Equal(t, expected, output)

}
