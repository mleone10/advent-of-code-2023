package day10

import (
	"github.com/mleone10/advent-of-code-2023/internal/grid"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type PipeField struct {
	input    string
	start    grid.Point
	fullGrid grid.Grid[tile]
	loop     []grid.Point
}

type tile struct {
	tType   rune
	visited bool
}

var neighborDeltas = map[rune][]grid.Point{
	'-': []grid.Point{{X: -1, Y: 0}, {X: 1, Y: 0}},
	'|': []grid.Point{{X: 0, Y: -1}, {X: 0, Y: 1}},
	'L': []grid.Point{{X: 0, Y: -1}, {X: 1, Y: 0}},
	'J': []grid.Point{{X: 0, Y: -1}, {X: -1, Y: 0}},
	'7': []grid.Point{{X: -1, Y: 0}, {X: 0, Y: 1}},
	'F': []grid.Point{{X: 0, Y: 1}, {X: 1, Y: 0}},
	'.': []grid.Point{},
	'S': []grid.Point{},
}

func NewPipeField(in string) PipeField {
	p := PipeField{
		input:    in,
		fullGrid: grid.Grid[tile]{},
		loop:     []grid.Point{},
	}

	for y, r := range slice.TrimSplit(in) {
		for x, c := range r {
			p.fullGrid.Set(x, y, tile{tType: c})
			if c == 'S' {
				p.start = grid.Point{X: x, Y: y}
			}
		}
	}

	p.loop = traverseLoop(p.start, []grid.Point{}, p.fullGrid)

	return p
}

func (p PipeField) StepsFarthestFromStart() int {
	return 0
}

func traverseLoop(cur grid.Point, loop []grid.Point, g grid.Grid[tile]) []grid.Point {
	c, _ := g.GetPoint(cur)
	slice.Filter(neighborDeltas[c.tType], func(n grid.Point) bool {

	})
}

func canGoUp(cur grid.Point, g grid.Grid[tile]) bool {
	up, ok := g.GetPoint(cur.Up())
	return ok && slice.Contains([]rune{'F', '|', '7'}, up.tType)
}
