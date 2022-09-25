package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRiverSizes(t *testing.T) {
	expected := []int{1, 2, 2, 2, 5}
	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}
