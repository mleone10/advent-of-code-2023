package geo

import (
	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

// A Polygon represents a 2D collection of sequential points which define a bounded space.  Points have width and depth.
type Polygon struct {
	points  []Point
	vectors []Line
}

// Add adds a new Point to the Polygon's point set.  The call also clears the Vectors() memo.
func (p *Polygon) Add(pt Point) {
	p.points = append(p.points, pt)
	p.vectors = nil
}

// Vectors returns a slice of Lines which connect the individual, sequential Points of the Polygon.  Vectors are memoized across calls to Vectors().
func (p *Polygon) Vectors() []Line {
	if len(p.points) < 3 {
		return []Line{}
	}
	if p.vectors != nil {
		return p.vectors
	}

	ls := []Line{}
	for i := 0; i < len(p.points)-1; i++ {
		ls = append(ls, Line{A: p.points[i], B: p.points[i+1]})
	}
	ls = append(ls, Line{A: p.points[len(p.points)-1], B: p.points[0]})

	p.vectors = ls

	return ls
}

// Perimeter returns the number of points along the Polygon's outer edge.
func (p Polygon) Perimeter() int {
	return slice.Reduce(p.Vectors(), 0, func(l Line, ret int) int {
		return ret + TaxicabLength(l)
	})
}

// Contains determines if a Point is strictly within a Polygon.  Points located on the edge of the Polygon do not count as being within it.
func (p *Polygon) Contains(pt Point) bool {
	is := 0
	for _, seg := range p.Vectors() {
		if pointOnLine(seg, pt) {
			return false
		}
		if rayIntersects(seg, pt) {
			is++
		}
	}
	return is%2 == 1
}

func pointOnLine(l Line, p Point) bool {
	dxp := p.X - l.A.X
	dyp := p.Y - l.A.Y
	dx1 := l.B.X - l.A.X
	dy1 := l.B.Y - l.A.Y

	if (dxp*dy1 - dyp*dx1) != 0 {
		return false
	}

	if mth.Abs(dx1) >= mth.Abs(dy1) {
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

func rayIntersects(l Line, p Point) bool {
	return (p.Y < l.A.Y) != (p.Y < l.B.Y) &&
		p.X < l.A.X+((p.Y-l.A.Y)/(l.B.Y-l.A.Y))*(l.B.X-l.A.X)
}
