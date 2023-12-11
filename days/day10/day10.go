package day10

import (
	"github.com/mleone10/advent-of-code-2023/internal/grid"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type PipeField struct {
	input string
	start grid.Point
	field grid.Grid[tile]
	loop  []grid.Point
}

type tile struct {
	tType rune
}

var neighborDeltas = map[rune][]grid.Point{
	'-': {grid.DeltaLeft, grid.DeltaRight},
	'|': {grid.DeltaUp, grid.DeltaDown},
	'L': {grid.DeltaUp, grid.DeltaRight},
	'J': {grid.DeltaUp, grid.DeltaLeft},
	'7': {grid.DeltaDown, grid.DeltaLeft},
	'F': {grid.DeltaDown, grid.DeltaRight},
	'.': {},
	'S': {grid.DeltaUp, grid.DeltaDown, grid.DeltaLeft, grid.DeltaRight},
}

func NewPipeField(in string) PipeField {
	p := PipeField{
		input: in,
		field: grid.Grid[tile]{},
		loop:  []grid.Point{},
	}

	for y, r := range slice.TrimSplit(in) {
		for x, c := range r {
			p.field.Set(x, y, tile{tType: c})
			if c == 'S' {
				p.start = grid.Point{X: x, Y: y}
			}
		}
	}

	p.traverseLoop(p.start)

	return p
}

func (p *PipeField) traverseLoop(cur grid.Point) {
	// If the loop already contains this point, do nothing.
	if loopContains(p.loop, cur) {
		return
	}

	// Otherwise, store the new point in the loop.
	p.loop = append(p.loop, cur)

	// Then, recurse to the valid neighbors of the current point.
	ns := []grid.Point{}
	cTile, _ := p.field.GetPoint(cur)
	for _, n := range slice.Map(neighborDeltas[cTile.tType], func(delta grid.Point) grid.Point {
		return cur.Add(delta)
	}) {
		t, _ := p.field.GetPoint(n)
		for _, d := range neighborDeltas[t.tType] {
			if n.Add(d).Equals(cur) {
				ns = append(ns, n)
			}
		}
	}

	for _, n := range ns {
		p.traverseLoop(n)
	}
}

func loopContains(ps []grid.Point, q grid.Point) bool {
	for _, p := range ps {
		if p.X == q.X && p.Y == q.Y {
			return true
		}
	}
	return false
}

func (p PipeField) StepsFarthestFromStart() int {
	return len(p.loop) / 2
}
