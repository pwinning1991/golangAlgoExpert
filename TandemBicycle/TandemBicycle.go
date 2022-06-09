package main

import (
	"math"
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)
	idx_for_pairing := 0
	runningSum := 0
	if fastest {
		idx_for_pairing = len(redShirtSpeeds) - 1
	} else {
		idx_for_pairing = 0
	}

	for idx := range redShirtSpeeds {
		tandemMax := math.Max(float64(redShirtSpeeds[idx]), float64(blueShirtSpeeds[idx_for_pairing]))
		runningSum += int(tandemMax)
		if fastest {
			idx_for_pairing -= 1
		} else {
			idx_for_pairing += 1
		}

	}

	return runningSum

}
