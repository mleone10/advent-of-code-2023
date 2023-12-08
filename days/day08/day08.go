package day08

import (
	"strings"
	"unicode"

	"github.com/mleone10/advent-of-code-2023/internal/mp"
	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type StartFunc func(v string) bool
type EndFunc func(v string) bool

var (
	StartAtAAA     = func(v string) bool { return v == "AAA" }
	StartAtAPrefix = func(v string) bool { return v[len(v)-1] == 'A' }
	EndAtZZZ       = func(v string) bool { return v == "ZZZ" }
	EndAtZSuffix   = func(v string) bool { return v[len(v)-1] == 'Z' }
)

type Map struct {
	Input string

	path  string
	nodes map[string]map[rune]string
}

func ShortestTraversalDist(m Map, isStart StartFunc, isDone EndFunc) int {
	m.init()

	// Identify all starting nodes
	curs := slice.Filter(mp.Keys(m.nodes), func(v string) bool {
		return isStart(v)
	})

	// Compute all of their traversal distances
	ds := slice.Map(curs, func(n string) int {
		return computeTraversalDist(m, n, isDone)
	})

	// The shortest traversal simultaneous traversal distance is the least common multiple of the individual distances
	return mth.Lcm(ds...)
}

func computeTraversalDist(m Map, node string, isDone EndFunc) int {
	dist := 0
	for !isDone(node) {
		d := m.path[dist%len(m.path)]
		node = m.nodes[node][rune(d)]
		dist++
	}

	return dist
}

func (m *Map) init() {
	if m.nodes != nil {
		return
	}

	iParts := strings.Split(m.Input, "\n\n")
	m.path = iParts[0]

	m.nodes = slice.Reduce(slice.TrimSplit(iParts[1]), map[string]map[rune]string{}, func(l string, ns map[string]map[rune]string) map[string]map[rune]string {
		nParts := strings.FieldsFunc(l, func(r rune) bool {
			return unicode.IsSpace(r) || unicode.IsSymbol(r) || unicode.IsPunct(r)
		})
		ns[nParts[0]] = map[rune]string{'L': nParts[1], 'R': nParts[2]}
		return ns
	})
}
