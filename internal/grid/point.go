package grid

// A Point is a coordinate pair in a 2D plane.
type Point struct {
	X, Y int
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
	return []Point{
		p.Up(), p.Down(), p.Left(), p.Right(),
	}
}

func (p Point) Up() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) Down() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) Left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (p Point) Right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}
