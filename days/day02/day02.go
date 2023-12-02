package day02

import (
	"strconv"
	"strings"
)

type cubeColor string

const (
	colorBlue  cubeColor = "blue"
	colorGreen cubeColor = "green"
	colorRed   cubeColor = "red"
)

var cubeSet = map[cubeColor]int{
	colorBlue:  14,
	colorGreen: 13,
	colorRed:   12,
}

type set map[cubeColor]int

type game struct {
	id   int
	sets []set
}

func SumPossibleGameIds(ls []string) int {
	sum := 0

	gs := parseGames(ls)
	for _, g := range gs {
		if isPossible(g) {
			sum += g.id
		}
	}

	return sum
}

func SumMinimumCubePower(ls []string) int {
	sum := 0

	gs := parseGames(ls)
	for _, g := range gs {
		sum += minimumPower(g)
	}

	return sum
}

func parseGames(ls []string) []game {
	gs := []game{}
	for _, l := range ls {
		gs = append(gs, parseGame(l))
	}
	return gs
}

func parseGame(l string) game {
	g := game{sets: []set{}}

	gParts := strings.Split(l, ":")
	g.id, _ = strconv.Atoi(strings.Fields(gParts[0])[1])

	sets := strings.Split(gParts[1], ";")
	for _, s := range sets {
		set := set{}
		colors := strings.Split(s, ",")
		for _, c := range colors {
			cParts := strings.Fields(c)
			color := cubeColor(cParts[1])
			num, _ := strconv.Atoi(cParts[0])
			set[color] = num
		}
		g.sets = append(g.sets, set)
	}

	return g
}

func isPossible(g game) bool {
	for _, s := range g.sets {
		for color, num := range s {
			if num > cubeSet[color] {
				return false
			}
		}
	}
	return true
}

func minimumPower(g game) int {
	colorMins := map[cubeColor]int{
		colorBlue:  0,
		colorGreen: 0,
		colorRed:   0,
	}

	for _, s := range g.sets {
		for color, num := range s {
			if num > colorMins[color] {
				colorMins[color] = num
			}
		}
	}

	ret := 1
	for _, min := range colorMins {
		ret *= min
	}

	return ret
}
