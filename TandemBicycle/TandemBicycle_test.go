package main

import "testing"

func TestTandemBicycle(t *testing.T) {
	redShirtSpeeds := []int{5, 5, 3, 9, 2}
	blueShirtSpeeds := []int{3, 6, 7, 2, 1}
	fastest := true
	expected := 32
	actual := TandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest)
	if expected != actual {
		t.Errorf("Got %d expected %d", actual, expected)
	}
}
