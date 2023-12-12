package day10

import (
	"github.com/mleone10/advent-of-code-2023/internal/grid"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type PipeField struct {
	input        string
	start        grid.Point
	field        grid.Grid[tile]
	loop         []grid.Point
	loopSegments []grid.Line
	prevCorner   grid.Point
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

func (p PipeField) StepsFarthestFromStart() int {
	return len(p.loop) / 2
}

func (p PipeField) TilesEnclosedByLoop() int {
	return grid.Reduce(p.field, 0, func(g grid.Grid[tile], x, y int, v tile, ret int) int {
		if isWithinLoop(p, grid.Point{X: x, Y: y}) {
			return ret + 1
		}
		return ret
	})
}

func NewPipeField(in string) PipeField {
	p := PipeField{
		input:        in,
		field:        grid.Grid[tile]{},
		loop:         []grid.Point{},
		loopSegments: []grid.Line{},
	}

	p.loadGrid(in)
	p.traverseLoop(p.start)

	return p
}

func (p *PipeField) loadGrid(in string) {
	for y, r := range slice.TrimSplit(in) {
		for x, c := range r {
			p.field.Set(x, y, tile{tType: c})
			if c == 'S' {
				p.start = grid.Point{X: x, Y: y}
			}
		}
	}
}

func (p *PipeField) traverseLoop(cur grid.Point) {
	// If the loop already contains this point, do nothing.
	if loopContains(p.loop, cur) {

		return
	}

	// Otherwise, store the new point in the loop, record a new loop segment, and traverse to the next neighbor.
	p.loop = append(p.loop, cur)
	ns := validNeighbors(*p, cur)

	if isStart(*p, cur) {
		// We only want to traverse the loop in one direction, so we only recurse in one direction from the start point.
		p.prevCorner = cur
		p.traverseLoop(ns[0])
		p.loopSegments = append(p.loopSegments, grid.Line{A: p.prevCorner, B: cur})
	} else {
		// If the current non-start point is a corner, record the end of a line segment
		if isCorner(*p, cur) {
			p.loopSegments = append(p.loopSegments, grid.Line{A: p.prevCorner, B: cur})
			p.prevCorner = cur
		}
		// Then recurse into all neighbors.  We only expect one to actually proceed with the loop.
		for _, n := range ns {
			p.traverseLoop(n)
		}
	}
}

func loopContains(ps []grid.Point, q grid.Point) bool {
	for _, p := range ps {
		if p.Equals(q) {
			return true
		}
	}
	return false
}

func validNeighbors(p PipeField, cur grid.Point) []grid.Point {
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
	return ns
}

func isStart(p PipeField, cur grid.Point) bool {
	t, _ := p.field.GetPoint(cur)
	return t.tType == 'S'
}

func isCorner(p PipeField, cur grid.Point) bool {
	t, _ := p.field.GetPoint(cur)
	return slice.Contains([]rune{'F', 'J', 'L', '7'}, t.tType)
}

func isWithinLoop(p PipeField, t grid.Point) bool {
	return true
}
