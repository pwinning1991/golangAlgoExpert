package main

import "testing"

func TestMinWaitingTime(t *testing.T) {
	queries := []int{3, 2, 1, 2, 6}
	expected := 17
	actual := MinWaitingTime(queries)
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

}
