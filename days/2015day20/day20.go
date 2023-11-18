package day20

import "github.com/mleone10/advent-of-code-2023/internal/maths"

// Today's puzzle seems to be all about prime numbers.  Specifically, the elves that visit a house are those whose numbers are the prime factors of the given house's number.

func FindHouseWithMinPresents(n int) int {
	for i := 0; ; i++ {
		if CalcPresents(i) >= n {
			return i
		}
	}
}

func CalcPresents(house int) int {
	fs := maths.Factors(house)

	p := 0
	// TODO: refactor this to use a slice.Reduce function
	for _, f := range fs {
		p += f * 10
	}

	return p
}
