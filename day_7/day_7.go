package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	puzzle            = "puzzle.txt"
	HIGH_CARD       = 1
	ONE_PAIR        = 2
	TWO_PAIR        = 3
	THREE_OF_A_KIND = 4
	FULL_HOUSE      = 5
	FOUR_OF_A_KIND  = 6
	FIVE_OF_A_KIND  = 7
)

type hand struct {
	cards    string
	bid      int
	bestHand int
}

var strength = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func GetBest(cards string) int {
	cardCounts := map[rune]int{}

	for _, card := range cards {
		cardCounts[card]++
	}

	switch len(cardCounts) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, count := range cardCounts {
			if count == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		for _, count := range cardCounts {
			if count == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 5:
		return HIGH_CARD
	default:
		return ONE_PAIR
	}
}

func wildJoke(h hand) int {
	cardCounts := map[rune]int{}

	for _, card := range h.cards {
		cardCounts[card] += 1
	}

	jokerCount := cardCounts['J']

	if jokerCount >= 4 {
		return FIVE_OF_A_KIND
	}

	if jokerCount == 3 {
		if len(cardCounts) == 2 {
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	}

	if jokerCount == 2 {
		if h.bestHand == TWO_PAIR {
			return FOUR_OF_A_KIND
		}

		if h.bestHand == ONE_PAIR {
			return THREE_OF_A_KIND
		}

		if h.bestHand == FULL_HOUSE {
			return FIVE_OF_A_KIND
		}
	}

	if jokerCount == 1 {
		if h.bestHand == THREE_OF_A_KIND {
			return FOUR_OF_A_KIND
		}

		if h.bestHand == TWO_PAIR {
			return FULL_HOUSE
		}

		if h.bestHand == ONE_PAIR {
			return THREE_OF_A_KIND
		}

		if h.bestHand == FOUR_OF_A_KIND {
			return FIVE_OF_A_KIND
		}
	}

	return ONE_PAIR
}

func main() {
	byteMap, _ := os.ReadFile(puzzle)
	puzzle := string(byteMap)
	lines := strings.Split(puzzle, "\n")
	hands := []hand{}
	score := 0

	score = calculateScore(lines, hands, false)	
	fmt.Printf("%#v\n", score)
	
	score = calculateScore(lines, hands, true)
	fmt.Printf("%#v\n", score)

}

func calculateScore(lines []string, hands []hand, isJokerWild bool) int {

	if isJokerWild {
		strength['J'] = -1
	}

	for _, line := range lines {
		values := strings.Split(line, " ")
		bid, _ := strconv.Atoi(values[1])
		h := hand{cards: values[0], bid: bid}
		h.bestHand = GetBest(h.cards)

		if isJokerWild && strings.Contains(h.cards, "J") {
			h.bestHand = wildJoke(h)
		}

		hands = append(hands, h)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].bestHand == hands[j].bestHand {
			for k := range hands[i].cards {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return strength[hands[i].cards[k]] < strength[hands[j].cards[k]]
			}
		}
		return hands[i].bestHand < hands[j].bestHand
	})

	score := 0
	for i, hand := range hands {
		score += (i + 1) * hand.bid
	}

	return score
}