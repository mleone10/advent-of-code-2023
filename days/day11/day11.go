package day11

import (
	"github.com/mleone10/advent-of-code-2023/internal/geo"
	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

const galaxyRune = '#'

type galaxy int

type Universe struct {
	ExpansionFactor int
	field           geo.Grid[galaxy]
}

func NewUniverse(in string, expansionFactor int) Universe {
	u := Universe{
		ExpansionFactor: expansionFactor,
		field:           geo.Grid[galaxy]{},
	}

	u.loadField(in)
	u.expandField()

	return u
}

func (u Universe) SumShortestPaths() int {
	return geo.Reduce(u.field, 0, func(g geo.Grid[galaxy], x1 int, y1 int, gal1 galaxy, res int) int {
		return res + geo.Reduce(u.field, 0, func(h geo.Grid[galaxy], x2 int, y2 int, gal2 galaxy, galRes int) int {
			if gal1 <= gal2 {
				return galRes
			}
			return galRes + geo.TaxicabLength(geo.Line{A: geo.Point{X: x1, Y: y1}, B: geo.Point{X: x2, Y: y2}})
		})
	})
}

func (u *Universe) loadField(in string) {
	n := 1
	for y, r := range slice.TrimSplit(in) {
		for x, c := range r {
			if c != galaxyRune {
				continue
			}
			u.field.Set(x, y, galaxy(n))
			n++
		}
	}
}

// Given a point (X, Y) with R empty rows above it and C empty columns left of it, the point's new coordinates are (X+C, Y+R)
func (u *Universe) expandField() {
	gs := geo.Points(u.field)                                 // geo.Points of all galaxies in the universe
	xs := slice.Map(gs, func(g geo.Point) int { return g.X }) // X coordinates of all galaxies
	ys := slice.Map(gs, func(g geo.Point) int { return g.Y }) // Y coordinates of all galaxies
	exs := []int{}
	eys := []int{}
	for i := 0; i < u.field.Width(); i++ {
		if !slice.Contains(xs, i) {
			exs = append(exs, i)
		}
	}
	for i := 0; i < u.field.Height(); i++ {
		if !slice.Contains(ys, i) {
			eys = append(eys, i)
		}
	}

	u.field = geo.Reduce(u.field, geo.Grid[galaxy]{}, func(g geo.Grid[galaxy], x int, y int, v galaxy, res geo.Grid[galaxy]) geo.Grid[galaxy] {
		exp := mth.Max(u.ExpansionFactor-1, 1)
		c := len(slice.Filter(exs, func(ex int) bool { return ex < x }))
		r := len(slice.Filter(eys, func(ey int) bool { return ey < y }))
		res.Set(x+(c*exp)-1, y+(r*exp)-1, v)
		return res
	})
}
