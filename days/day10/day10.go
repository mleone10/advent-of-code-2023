package day10

import (
	"github.com/mleone10/advent-of-code-2023/internal/geo"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type PipeField struct {
	input string
	start geo.Point
	field geo.Grid[rune]
	loop  []geo.Point

	poly geo.Polygon
}

var neighborDeltas = map[rune][]geo.Point{
	'-': {geo.DeltaLeft, geo.DeltaRight},
	'|': {geo.DeltaUp, geo.DeltaDown},
	'L': {geo.DeltaUp, geo.DeltaRight},
	'J': {geo.DeltaUp, geo.DeltaLeft},
	'7': {geo.DeltaDown, geo.DeltaLeft},
	'F': {geo.DeltaDown, geo.DeltaRight},
	'.': {},
	'S': {geo.DeltaUp, geo.DeltaDown, geo.DeltaLeft, geo.DeltaRight},
}

func (p PipeField) StepsFarthestFromStart() int {
	return p.poly.Perimeter() / 2
}

func (p PipeField) TilesEnclosedByLoop() int {
	return geo.Reduce(p.field, 0, func(g geo.Grid[rune], x, y int, v rune, ret int) int {
		if p.poly.Contains(geo.Point{X: x, Y: y}) {
			// if isWithinLoop(p, grid.Point{X: x, Y: y}) {
			return ret + 1
		}
		return ret
	})
}

func NewPipeField(in string) PipeField {
	p := PipeField{
		input: in,
		field: geo.Grid[rune]{},
		loop:  []geo.Point{},
		poly:  geo.Polygon{},
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
				p.start = geo.Point{X: x, Y: y}
			}
		}
	}
}

func (p *PipeField) traverseLoop(cur geo.Point) {
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

func loopContains(ps []geo.Point, pt geo.Point) bool {
	for _, p := range ps {
		if p.Equals(pt) {
			return true
		}
	}
	return false
}

func validNeighbors(p PipeField, cur geo.Point) []geo.Point {
	ns := []geo.Point{}
	c, _ := p.field.GetPoint(cur)
	for _, n := range slice.Map(neighborDeltas[c], func(delta geo.Point) geo.Point {
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
