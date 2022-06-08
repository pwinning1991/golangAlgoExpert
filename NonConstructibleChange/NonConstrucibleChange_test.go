package main

import "testing"

type coinData struct {
	coins  []int
	output int
}

func TestNonConstrucibleChange(t *testing.T) {
	coinDataList := []coinData{
		{
			coins:  []int{5, 7, 1, 1, 2, 3, 22},
			output: 20,
		},
		{
			coins:  []int{1, 1, 1, 1, 1},
			output: 6,
		},
	}
	for _, data := range coinDataList {
		got := NonConstrucibleChange(data.coins)
		if got != data.output {
			t.Errorf("Did not get right answer got %d wanted %d", got, data.output)

		}

	}

}
