package main

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Slice(redShirtHeights, func(i, j int) bool {
		return redShirtHeights[i] > redShirtHeights[j]
	})
	sort.Slice(blueShirtHeights, func(i, j int) bool {
		return blueShirtHeights[i] > blueShirtHeights[j]
	})

	shirtColorInFirstRow := "BLUE"
	if redShirtHeights[0] < blueShirtHeights[0] {
		shirtColorInFirstRow = "RED"
	}

	for idx := range redShirtHeights {
		redShirtHeight := redShirtHeights[idx]
		blueShirtHeight := blueShirtHeights[idx]

		if shirtColorInFirstRow == "RED" {
			if redShirtHeight >= blueShirtHeight {
				return false
			}
		} else {
			if blueShirtHeight >= redShirtHeight {
				return false
			}
		}

	}

	return true

}
