package day10

import (
	"github.com/mleone10/advent-of-code-2023/internal/grid"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type PipeField struct {
	input        string
	start        grid.Point
	field        grid.Grid[rune]
	loop         []grid.Point
	loopSegments []grid.Line
	prevCorner   grid.Point
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
	return grid.Reduce(p.field, 0, func(g grid.Grid[rune], x, y int, v rune, ret int) int {
		if isWithinLoop(p, grid.Point{X: x, Y: y}) {
			return ret + 1
		}
		return ret
	})
}

func NewPipeField(in string) PipeField {
	p := PipeField{
		input:        in,
		field:        grid.Grid[rune]{},
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

func isStart(p PipeField, cur grid.Point) bool {
	t, _ := p.field.GetPoint(cur)
	return t == 'S'
}

func isCorner(p PipeField, cur grid.Point) bool {
	t, _ := p.field.GetPoint(cur)
	return slice.Contains([]rune{'F', 'J', 'L', '7'}, t)
}

func isWithinLoop(p PipeField, t grid.Point) bool {
	intersections := 0
	for _, seg := range p.loopSegments {
		if pointOnLine(seg, t) {
			return false
		}
		if rayIntersects(seg, t) {
			intersections++
		}
	}
	return intersections%2 == 1
}

func rayIntersects(l grid.Line, p grid.Point) bool {
	return (p.Y < l.A.Y) != (p.Y < l.B.Y) &&
		p.X < l.A.X+((p.Y-l.A.Y)/(l.B.Y-l.A.Y))*(l.B.X-l.A.X)
}

func pointOnLine(l grid.Line, p grid.Point) bool {
	dxp := p.X - l.A.X
	dyp := p.Y - l.A.Y
	dx1 := l.B.X - l.A.X
	dy1 := l.B.Y - l.A.Y

	if (dxp*dy1 - dyp*dx1) != 0 {
		return false
	}

	if abs(dx1) >= abs(dy1) {
		if dx1 > 0 {
			return l.A.X <= p.X && p.X <= l.B.X
		}
		return l.B.X <= p.X && p.X <= l.A.X
	} else {
		if dy1 > 0 {
			return l.A.Y <= p.Y && p.Y <= l.B.Y
		}
		return l.B.Y <= p.Y && p.Y <= l.A.Y
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
