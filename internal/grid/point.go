package grid

// A Point is a coordinate pair in a 2D plane.
type Point struct {
	X, Y int
}

var (
	DeltaUp    = Point{X: 0, Y: -1}
	DeltaDown  = Point{X: 0, Y: 1}
	DeltaLeft  = Point{X: -1, Y: 0}
	DeltaRight = Point{X: 1, Y: 0}

	CardinalDeltas = []Point{DeltaUp, DeltaDown, DeltaLeft, DeltaRight}
)

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

// The Neighbors of point `p` are all those points above, below, left, right, or diagonal from `p` with positive coordinates of their own.
func Neighbors(p Point) []Point {
	ps := []Point{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newX, newY := p.X+i, p.Y+j
			if newX == p.X && newY == p.Y {
				continue
			}
			if newX >= 0 && newY >= 0 {
				ps = append(ps, Point{newX, newY})
			}
		}
	}
	return ps
}

func CardinalNeighbors(p Point) []Point {
	ns := []Point{}
	for _, d := range CardinalDeltas {
		n := p.Add(d)
		if n.X >= 0 && n.Y >= 0 {
			ns = append(ns, n)
		}
	}
	return ns
}
