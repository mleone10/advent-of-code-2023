package day09

import (
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/set"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

func Next(l string) int {
	seq := parseReadings(l)

	return seq[len(seq)-1] + rightDiff(seq)
}

func Prev(l string) int {
	seq := parseReadings(l)

	return seq[0] - leftDiff(seq)
}

func parseReadings(r string) []int {
	return slice.Map(strings.Fields(r), func(s string) int {
		return mth.Atoi(s)
	})
}

func rightDiff(seq []int) int {
	diffs := []int{}
	for i := 0; i < len(seq)-1; i++ {
		diffs = append(diffs, seq[i+1]-seq[i])
	}

	uniqDiffs := set.Set[int]{}
	uniqDiffs.Add(diffs...)

	if len(uniqDiffs) == 1 {
		return diffs[0]
	}

	return diffs[len(diffs)-1] + rightDiff(diffs)
}

func leftDiff(seq []int) int {
	diffs := []int{}
	for i := 0; i < len(seq)-1; i++ {
		diffs = append(diffs, seq[i+1]-seq[i])
	}

	uniqDiffs := set.Set[int]{}
	uniqDiffs.Add(diffs...)

	if len(uniqDiffs) == 1 {
		return diffs[0]
	}

	return diffs[0] - leftDiff(diffs)
}
