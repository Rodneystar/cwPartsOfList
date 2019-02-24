package cwnbateams

import (
	"fmt"
	"strconv"
	"strings"
)

func nba_cup(results string, team string) string {
	return ""
}

func parseGame(str string) map[string]int {
	var scores = make(map[string]int)

	return scores
}

func getGameScores(str string) [2]string {
	var indScores [2]string
	i := 0
	tempName := []string{}
	for _, token := range strings.Split(str, " ") {
		tempName = append(tempName, token)
		_, err := strconv.Atoi(token)
		if err == nil {
			indScores[i] = strings.Join(tempName, " ")
			tempName = nil
			i++
		}
	}
	return indScores
}

func parseResults(result string) ([]string, []int) {
	var teams, scores = make([]string, 2), make([]int, 2)

	fmt.Sscanf(result+" \n", "%s %d %s %d", &teams[0], &scores[0], &teams[1], &scores[1])
	return teams, scores
}
