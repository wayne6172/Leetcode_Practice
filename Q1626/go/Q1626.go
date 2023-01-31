package Q1626

import "sort"

type Player struct {
	score int
	age   int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bestTeamScore(scores []int, ages []int) int {
	playerNum := len(scores)
	players := make([]Player, playerNum)
	for i := 0; i < playerNum; i++ {
		players[i] = Player{
			score: scores[i],
			age:   ages[i],
		}
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].age > players[j].age
	})

	ans := 0
	dp := make([]int, playerNum)
	for i := 0; i < len(players); i++ {
		dp[i] = players[i].score
		for j := 0; j < i; j++ {
			if players[i].score <= players[j].score {
				dp[i] = max(dp[i], dp[j]+players[i].score)
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}
