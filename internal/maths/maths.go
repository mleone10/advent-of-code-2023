package maths

import "github.com/mleone10/advent-of-code-2023/internal/set"

var memoFactors = map[int][]int{}

// Factors calculates all factors of `n`, returning them in a slice of ints.
func Factors(n int) []int {
	fs := set.Set[int]{}

	// Terminal case
	if n == 1 {
		return []int{1}
	}

	// Check for cached factors in memo
	if f, ok := memoFactors[n]; ok {
		return f
	}

	// Loop through integers.  When a factor is found, store it, then recurse to find all Factors of the quotient.  Add them to the map.
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			fs.Add(i)
			fs.Add(Factors(n / i)...)
			memoFactors[n] = append(memoFactors[n], fs.Slice()...)
		}
	}

	return fs.Slice()
}
