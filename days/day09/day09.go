package day09

import (
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

func Next(l string) int {
	seq := parseReadings(l)

	return seq[len(seq)-1] + diff(seq)
}

func Prev(l string) int {
	seq := parseReadings(l)

	return seq[0] + diff(slice.Reverse(seq))
}

func parseReadings(r string) []int {
	return slice.Map(strings.Fields(r), func(s string) int {
		return mth.Atoi(s)
	})
}

// Diffs recursively determines the next difference of a given sequence `seq`.
func diff(seq []int) int {
	// Compute the differences between elements of `seq`.
	diffs := []int{}
	for i := 0; i < len(seq)-1; i++ {
		diffs = append(diffs, seq[i+1]-seq[i])
	}

	// If all differences are zeroes, we've reached the bottom of the recursion.
	if len(slice.Filter(diffs, func(v int) bool { return v != 0 })) == 0 {
		return diffs[0]
	}

	// Otherwise, recurse to find the differences between the newly-computed sequence `diffs`.  Then return back to the caller with the computed next difference for the original `seq`.
	return diffs[len(diffs)-1] + diff(diffs)
}
