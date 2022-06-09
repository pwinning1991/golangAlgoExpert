package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateDocument(t *testing.T) {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"
	expected := true
	actual := GenerateDocument(characters, document)
	require.Equal(t, expected, actual)
}
