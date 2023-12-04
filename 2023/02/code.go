package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}
func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func gameId(game string) int {
	splitLine := strings.Fields(game)
	idAsString := strings.TrimSuffix(splitLine[1], ":")
	id, _ := strconv.Atoi(idAsString)
	return id
}
func sets(game string) []string {
	m1 := regexp.MustCompile(`Game \d: `)
	return strings.Split(m1.ReplaceAllString(game, ""), "; ")
}
func isSetPossible(set string) bool {
	redRuleBroken := regexp.MustCompile(`(1[3-9]|[2-9][0-9]) red`)
	greenRuleBroken := regexp.MustCompile(`(1[4-9]|[2-9][0-9]) green`)
	blueRuleBroken := regexp.MustCompile(`(1[5-9]|[2-9][0-9]) blue`)
	if redRuleBroken.MatchString(set) || greenRuleBroken.MatchString(set) || blueRuleBroken.MatchString(set) {
		return false
	}
	return true
}
func idsSumOfPossibleGame(input string) int {
	separatedGames := strings.Split(input, "\n")
	var idsOfPossibleGame []int
	for _, game := range separatedGames {
		if len(game) == 0 {
			continue
		}
		isGamePossible := true
		for _, set := range sets(game) {
			if !isSetPossible(set) {
				isGamePossible = false
				break
			}
		}
		if isGamePossible {
			idsOfPossibleGame = append(idsOfPossibleGame, gameId(game))
		}
	}
	return sum(idsOfPossibleGame)
}
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return idsSumOfPossibleGame(input)
}
