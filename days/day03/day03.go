package day03

import (
	"strconv"
	"unicode"

	"github.com/mleone10/advent-of-code-2023/internal/geo"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type partNumber struct {
	val  int
	locs []geo.Point
}

func PartNumberSum(ls []string) int {
	ps := parsePartNumbers(ls)
	ss := parseSymbols(ls)

	validParts := slice.Filter(ps, func(pn partNumber) bool {
		adj := false
		for _, l := range pn.locs {
			for _, p := range geo.Neighbors(l) {
				if _, ok := ss.Get(p.X, p.Y); ok {
					adj = true
					break
				}
			}
		}
		return adj
	})

	return slice.Reduce(validParts, 0, func(p partNumber, res int) int {
		return res + p.val
	})
}

func GearRatioSum(ls []string) int {
	ps := parsePartNumbers(ls)
	gs := parseGears(ls)

	ratios := []int{}
	for _, g := range gs {
		adjParts := []partNumber{}
		ns := geo.Neighbors(g)
		for _, p := range ps {
			adj := false
			for _, l := range p.locs {
				for _, n := range ns {
					if n.X == l.X && n.Y == l.Y {
						adj = true
					}
				}
			}
			if adj {
				adjParts = append(adjParts, p)
			}
		}
		if len(adjParts) == 2 {
			ratios = append(ratios, adjParts[0].val*adjParts[1].val)
		}
	}

	return slice.Reduce(ratios, 0, func(i int, res int) int {
		return res + i
	})
}

func parsePartNumbers(ls []string) []partNumber {
	partNumbers := []partNumber{}

	// Identify all part numbers in the grid
	for y, l := range ls {
		numString := ""
		var p partNumber
		for x, r := range l {
			if unicode.IsDigit(r) && numString == "" {
				// Found start of new part number
				numString += string(r)
				p = partNumber{locs: []geo.Point{{X: x, Y: y}}}
			} else if unicode.IsDigit(r) {
				// Found new digit in existing part number
				numString += string(r)
				p.locs = append(p.locs, geo.Point{X: x, Y: y})
			}
			if numString != "" && (!unicode.IsDigit(r) || x == len(l)-1) {
				// Reached end of number or end of line
				p.val, _ = strconv.Atoi(numString)
				partNumbers = append(partNumbers, p)
				p = partNumber{}
				numString = ""
			}
		}
	}

	return partNumbers
}

func parseSymbols(ls []string) geo.Grid[bool] {
	g := geo.Grid[bool]{}

	for y, l := range ls {
		for x, r := range l {
			if r != '.' && !unicode.IsDigit(r) {
				g.Set(x, y, true)
			}
		}
	}

	return g
}

func parseGears(ls []string) []geo.Point {
	gs := []geo.Point{}

	for y, l := range ls {
		for x, r := range l {
			if r == '*' {
				gs = append(gs, geo.Point{X: x, Y: y})
			}
		}
	}

	return gs
}
