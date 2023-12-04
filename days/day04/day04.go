package day04

import (
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mp"
	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type game struct {
	id          int
	winningNums []int
	yourNums    []int
}

func PilePoints(ls []string) int {
	games := parseGames(ls)

	return slice.Reduce(games, 0, func(g game, ret int) int {
		return ret + pointValue(g)
	})
}

func NumCards(ls []string) int {
	// I actually had this solved with plain for-loops, but wanted to play with converting it to use slice.Reduce instead.  It's less readable, but I had fun with the exercise.
	return slice.Sum(mp.Values(slice.Reduce(parseGames(ls), map[int]int{}, func(g game, numCards map[int]int) map[int]int {
		if _, ok := numCards[g.id]; !ok {
			numCards[g.id] = 1
		} else {
			numCards[g.id] += 1
		}
		numWinners := numWinners(g)
		for i := 1; i <= numWinners; i++ {
			numCards[g.id+i] += numCards[g.id]
		}
		return numCards
	})))
}

func parseGames(ls []string) []game {
	return slice.Reduce(ls, []game{}, func(l string, gs []game) []game {
		g := game{}

		gParts := strings.Split(l, ":")
		g.id = mth.Atoi(strings.Fields(gParts[0])[1])
		numParts := strings.Split(gParts[1], "|")
		g.winningNums = parseNums(numParts[0])
		g.yourNums = parseNums(numParts[1])

		return append(gs, g)
	})
}

func parseNums(s string) []int {
	return slice.Map(strings.Fields(s), func(sNum string) int {
		return mth.Atoi(sNum)
	})
}

func pointValue(g game) int {
	ws := numWinners(g)

	if ws == 0 {
		return 0
	}
	return 1 * mth.Pow(2, ws-1)
}

func numWinners(g game) int {
	return slice.Reduce(g.winningNums, 0, func(wn int, wnRet int) int {
		winner := isWinningNumber(g, wn)
		if winner {
			return wnRet + 1
		}
		return wnRet
	})
}

func isWinningNumber(g game, n int) bool {
	return slice.Reduce(g.yourNums, false, func(yn int, winner bool) bool {
		if n != yn {
			return winner
		}
		return true
	})
}
