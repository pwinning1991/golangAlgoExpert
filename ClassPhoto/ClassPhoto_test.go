package main

import "testing"

func TestClassPhotos(t *testing.T) {
	redShirtHeights := []int{5, 8, 1, 3, 4}
	blueShirtHeights := []int{6, 9, 2, 4, 5}
	expected := true
	actual := ClassPhotos(redShirtHeights, blueShirtHeights)
	if expected != actual {
		t.Errorf("Got %v Expected %v\n", actual, expected)
	}

}
