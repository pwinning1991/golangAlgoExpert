package main

import (
	"reflect"
	"testing"
)

type squaredata struct {
	inputs []int
	result []int
}

func TestSortedSquaredArray(t *testing.T) {
	dataToTest := []squaredata{
		{
			inputs: []int{1, 2, 3, 4},
			result: []int{1, 4, 9, 16},
		},
		{
			inputs: []int{3, 4, 5},
			result: []int{9, 16, 25},
		},
	}
	for _, data := range dataToTest {
		got := SortedSquaredArray(data.inputs)
		want := data.result
		reflect.DeepEqual(got, want)
	}

}
