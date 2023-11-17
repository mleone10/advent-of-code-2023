package day19

import (
	"math/rand"
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mp"
	"github.com/mleone10/advent-of-code-2023/internal/set"
	"github.com/mleone10/advent-of-code-2023/internal/str"
)

type Replacements map[string][]string

func NewReplacements(input string) Replacements {
	r := Replacements{}

	for _, l := range strings.Split(strings.TrimSpace(input), "\n") {
		from, to, _ := strings.Cut(l, " => ")
		if _, ok := r[from]; !ok {
			r[from] = []string{}
		}
		r[from] = append(r[from], to)
	}

	return r
}

func CalibrationSum(init string, repl Replacements) int {
	uniqMols := set.Set[string]{}
	for src, opts := range repl {
		for _, o := range opts {
			uniqMols.Add(str.ReplaceVariants(init, src, o)...)
		}
	}
	return uniqMols.Size()
}

func FabricationLength(init string, repl Replacements) int {
	rev := reverseReplacements(repl)
	dests := mp.Keys(rev)

	steps := 0
	s := init
	for s != "e" {
		changeMade := false
		for _, d := range dests {
			reps := strings.Count(s, d)
			if reps == 0 {
				continue
			}
			s = strings.Replace(s, d, rev[d], 1)
			steps++
			changeMade = true
			break
		}
		if !changeMade {
			rand.Shuffle(len(dests), func(i, j int) { dests[i], dests[j] = dests[j], dests[i] })
			steps = 0
			s = init
		}
	}
	return steps
}

func reverseReplacements(repl Replacements) map[string]string {
	rev := map[string]string{}
	for k, v := range repl {
		for _, r := range v {
			rev[r] = k
		}
	}
	return rev
}
