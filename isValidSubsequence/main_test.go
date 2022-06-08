package main

import "testing"

func TestIsValidSubsequence(t *testing.T) {
	got := IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})
	want := true
	if got != want {
		t.Errorf("Test failed got %v wanted %v\n", got, want)
	}

}
