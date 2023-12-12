package day10

import (
	"github.com/mleone10/advent-of-code-2023/internal/grid"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type PipeField struct {
	input string
	start grid.Point
	field grid.Grid[rune]
	loop  []grid.Point

	poly grid.Polygon
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
	return p.poly.Perimeter() / 2
}

func (p PipeField) TilesEnclosedByLoop() int {
	return grid.Reduce(p.field, 0, func(g grid.Grid[rune], x, y int, v rune, ret int) int {
		if p.poly.Contains(grid.Point{X: x, Y: y}) {
			// if isWithinLoop(p, grid.Point{X: x, Y: y}) {
			return ret + 1
		}
		return ret
	})
}

func NewPipeField(in string) PipeField {
	p := PipeField{
		input: in,
		field: grid.Grid[rune]{},
		loop:  []grid.Point{},
		poly:  grid.Polygon{},
	}

	p.loadGrid(in)
	p.traverseLoop(p.start)

	return p
}

func (p *PipeField) loadGrid(in string) {
	for y, r := range slice.TrimSplit(in) {
		for x, c := range r {
			p.field.Set(x, y, c)
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

	// Otherwise, store the new point in the loop and traverse to the next neighbor.
	p.loop = append(p.loop, cur)
	ns := validNeighbors(*p, cur)
	t, _ := p.field.GetPoint(cur)

	if t == 'S' {
		// We only want to traverse the loop in one direction, so we only recurse in one direction from the start point.
		p.poly.Add(cur)
		p.traverseLoop(ns[0])
	} else {
		// If the current non-start point is a corner, record a new point on the polygon
		if slice.Contains([]rune{'F', 'J', 'L', '7'}, t) {
			p.poly.Add(cur)
		}
		// Then recurse into all neighbors.  We only expect one to actually proceed with the loop, as the other was just visited.
		for _, n := range ns {
			p.traverseLoop(n)
		}
	}
}

func loopContains(ps []grid.Point, pt grid.Point) bool {
	for _, p := range ps {
		if p.Equals(pt) {
			return true
		}
	}
	return false
}

func validNeighbors(p PipeField, cur grid.Point) []grid.Point {
	ns := []grid.Point{}
	c, _ := p.field.GetPoint(cur)
	for _, n := range slice.Map(neighborDeltas[c], func(delta grid.Point) grid.Point {
		return cur.Add(delta)
	}) {
		t, _ := p.field.GetPoint(n)
		for _, d := range neighborDeltas[t] {
			if n.Add(d).Equals(cur) {
				ns = append(ns, n)
			}
		}
	}
	return ns
}
