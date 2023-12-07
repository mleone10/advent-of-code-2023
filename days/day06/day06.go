package day06

import (
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type race struct {
	dur, dist int
}

func MarginOfError(ls []string) int {
	rs := parseRaces(ls)

	return slice.Reduce(rs, 1, func(r race, ret int) int {
		return ret * len(winningHoldDurs(r))
	})
}

func SingleRaceWinningSolutions(ls []string) int {
	r := parseRace(ls)
	return len(winningHoldDurs(r))
}

func parseRaces(ls []string) []race {
	durs := strings.Fields(ls[0])[1:]
	dists := strings.Fields(ls[1])[1:]

	rs := []race{}
	for n, d := range durs {
		rs = append(rs, race{mth.Atoi(d), mth.Atoi(dists[n])})
	}

	return rs
}

func parseRace(ls []string) race {
	return race{
		mth.Atoi(strings.Join(strings.Fields(strings.Split(ls[0], ":")[1]), "")),
		mth.Atoi(strings.Join(strings.Fields(strings.Split(ls[1], ":")[1]), "")),
	}
}

func winningHoldDurs(r race) []int {
	ds := []int{}
	for t := 0; t < r.dur; t++ {
		if travelDist(r.dur, t) > r.dist {
			ds = append(ds, t)
		}
	}
	return ds
}

func travelDist(maxMs, holdMs int) int {
	return (maxMs - holdMs) * holdMs
}
