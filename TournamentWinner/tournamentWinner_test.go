package main

import "testing"

func TestTournamentWinner(t *testing.T) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	got := TournamentWinner(competitions, results)
	want := "Python"
	if got != want {
		t.Errorf("Got %s Wanted %s", got, want)

	}

}
