package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}

	for idx, comcompetition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := comcompetition[0], comcompetition[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}

	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}
