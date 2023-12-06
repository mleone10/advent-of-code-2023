package day05

import (
	"math"
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type RangeFunc func(n int) int

type Almanac struct {
	Seeds []int
	Maps  map[string]RangeFunc
}

func NewAlmanac(in string) Almanac {
	a := Almanac{
		Maps: map[string]RangeFunc{},
	}

	for _, sec := range strings.Split(in, "\n\n") {
		ls := slice.TrimSplit(sec)
		headParts := strings.Split(ls[0], ":")
		switch headParts[0] {
		case "seeds":
			seedStrings := strings.Fields(headParts[1])
			a.Seeds = slice.Map(seedStrings, func(s string) int {
				return mth.Atoi(s)
			})
			break
		default:
			a.Maps[strings.Split(ls[0], " ")[0]] = NewRangeFunc(ls[1:])
			break
		}
	}

	return a
}

type mapRange struct {
	destStart, srcStart, max int
}

func NewRangeFunc(rs []string) RangeFunc {
	mrs := []mapRange{}
	for _, r := range rs {
		rParts := strings.Fields(r)
		destStart := mth.Atoi(rParts[0])
		srcStart := mth.Atoi(rParts[1])
		length := mth.Atoi(rParts[2])

		mrs = append(mrs, mapRange{destStart, srcStart, srcStart + length})
	}

	return func(n int) int {
		for _, mr := range mrs {
			if n >= mr.srcStart && n < mr.max {
				return mr.destStart + (n - mr.srcStart)
			}
		}
		return n
	}
}

func (a Almanac) Location(seed int) int {
	return a.Maps["humidity-to-location"](
		a.Maps["temperature-to-humidity"](
			a.Maps["light-to-temperature"](
				a.Maps["water-to-light"](
					a.Maps["fertilizer-to-water"](
						a.Maps["soil-to-fertilizer"](
							a.Maps["seed-to-soil"](seed),
						),
					),
				),
			),
		),
	)
}

func (a Almanac) ClosetLocation() int {
	return mth.Min(slice.Map(a.Seeds, a.Location)...)
}

func (a Almanac) ClosetLocationSeedRange() int {
	minLoc := math.MaxInt

	for i := 0; i < len(a.Seeds)/2; i++ {
		start := a.Seeds[i*2]
		length := a.Seeds[i*2+1]

		for s := 0; s < length; s++ {
			loc := a.Location(start + s)
			if loc < minLoc {
				minLoc = loc
			}
		}
	}

	return minLoc
}
