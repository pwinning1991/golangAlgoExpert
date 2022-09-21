package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayOfProdcuts(t *testing.T) {
	input := []int{5, 1, 4, 2}
	expected := []int{8, 40, 10, 20}
	actual := ArrayOfProducts(input)
	require.Equal(t, expected, actual)
}
