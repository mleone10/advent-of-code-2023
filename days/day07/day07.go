package day07

import (
	"sort"
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mp"
	"github.com/mleone10/advent-of-code-2023/internal/mth"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type handType int

const (
	typeFiveOfAKind handType = iota
	typeFourOfAKind
	typeFullHouse
	typeThreeOfAKind
	typeTwoPair
	typeOnePair
	typeHighCard
)

const (
	cardOrderNoWilds    = "AKQJT98765432"
	cardOrderJokersWild = "AKQT98765432J"
	cardJoker           = "J"
)

var (
	jokersWild = false
	cardOrder  = cardOrderNoWilds
)

type hand struct {
	cards string
	bid   int
	hType handType
}

func TotalWinnings(ls []string) int {
	hs := parseHands(ls)

	sort.Slice(hs, func(i, j int) bool {
		return !hs[i].strongerThan(hs[j])
	})

	winnings := 0
	for i := 0; i < len(hs); i++ {
		winnings += (i + 1) * hs[i].bid
	}

	return winnings
}

func JokersWild() {
	cardOrder = cardOrderJokersWild
	jokersWild = true
}

func parseHands(ls []string) []hand {
	return slice.Reduce(ls, []hand{}, func(l string, hs []hand) []hand {
		return append(hs, parseHand(l))
	})
}

func parseHand(l string) hand {
	hParts := strings.Fields(l)
	return hand{
		cards: hParts[0],
		bid:   mth.Atoi(hParts[1]),
		hType: getHandType(hParts[0]),
	}
}

func getHandType(h string) handType {
	rs := mp.RuneCount(h)
	cardCounts := mp.Values(rs)
	maxMatch := mth.Max(cardCounts...)
	js := jokersWild && strings.Contains(h, cardJoker)

	if maxMatch == 5 {
		// Five of a kind
		return typeFiveOfAKind
	}
	if maxMatch == 4 && js {
		// Four of a kind with joker
		return typeFiveOfAKind
	}
	if maxMatch == 4 {
		// Four of a kind
		return typeFourOfAKind
	}
	if len(rs) == 2 && maxMatch == 3 && js {
		// Full house with jokers
		return typeFiveOfAKind
	}
	if len(rs) == 2 && maxMatch == 3 {
		// Full house
		return typeFullHouse
	}
	if maxMatch == 3 && js {
		// Three of a kind with two jokers
		return typeFourOfAKind
	}
	if maxMatch == 3 {
		return typeThreeOfAKind
	}

	twoPair := len(slice.Filter(cardCounts, func(n int) bool { return n == 2 })) == 2
	if twoPair && js && rs[rune(cardJoker[0])] == 2 {
		// Two pair, one of which is a pair of jokers
		return typeFourOfAKind
	}
	if twoPair && js {
		// Two pair, plus a spare joker
		return typeFullHouse
	}
	if twoPair {
		return typeTwoPair
	}

	onePair := maxMatch == 2
	if onePair && js {
		return typeThreeOfAKind
	}
	if onePair || js {
		return typeOnePair
	}
	return typeHighCard
}

func (h hand) strongerThan(s hand) bool {
	if h.hType != s.hType {
		// If the hand types differ, return the stronger one
		return h.hType < s.hType
	}

	// Otherwise, the hand types are the same and we have to go by the card sort order
	for i := 0; i < len(h.cards); i++ {
		hCardIdx := strings.Index(cardOrder, string(h.cards[i]))
		sCardIdx := strings.Index(cardOrder, string(s.cards[i]))
		if hCardIdx < sCardIdx {
			return true
		}
		if hCardIdx > sCardIdx {
			return false
		}
	}
	return true
}
