package day18

import (
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/grid"
)

type Light struct {
	on    bool
	stuck bool
}

type LightGrid struct {
	grid          grid.Grid[*Light]
	width, height int
}

func NewLightGrid(input string) *LightGrid {
	g := grid.Grid[*Light]{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, l := range lines {
		for x, c := range l {
			g.Set(x, y, &Light{c == '#', false})
		}
	}

	return &LightGrid{
		grid:   g,
		width:  len(lines[0]),
		height: len(lines),
	}
}

func (lg *LightGrid) NumOn() int {
	return grid.Reduce(lg.grid, 0, func(g grid.Grid[*Light], x, y int, v *Light, res int) int {
		if v.on {
			return res + 1
		}
		return res
	})
}

func (lg *LightGrid) Step() {
	lg.grid = grid.Map(lg.grid, lightIsOn)
}

func (lg *LightGrid) StepN(n int) {
	for i := 0; i < n; i++ {
		lg.grid = grid.Map(lg.grid, lightIsOn)
	}
}

func (lg *LightGrid) CornersStuckOn() {
	topLeft, _ := lg.grid.Get(0, 0)
	topRight, _ := lg.grid.Get(lg.width-1, 0)
	bottomLeft, _ := lg.grid.Get(0, lg.height-1)
	bottomRight, _ := lg.grid.Get(lg.width-1, lg.height-1)

	topLeft.on = true
	topLeft.stuck = true
	topRight.on = true
	topRight.stuck = true
	bottomLeft.on = true
	bottomLeft.stuck = true
	bottomRight.on = true
	bottomRight.stuck = true
}

func lightIsOn(g grid.Grid[*Light], x, y int, v *Light) *Light {
	if v, ok := g.Get(x, y); ok && v.stuck {
		return &Light{v.on, v.stuck}
	}

	neighborsOn := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if l, ok := g.Get(x+i, y+j); ok && l.on && !(i == 0 && j == 0) {
				neighborsOn++
			}
		}
	}

	return &Light{
		on:    (v.on && (neighborsOn == 2 || neighborsOn == 3)) || (!v.on && neighborsOn == 3),
		stuck: v.stuck,
	}
}
