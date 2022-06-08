package main

import (
	"reflect"
	"testing"
)

func TestTwoNumSum(t *testing.T) {
	got := TwoNumberSum([]int{1, 2, 7, 9, 5}, 6)
	want := []int{1, 5}
	reflect.DeepEqual(got, want)

}
