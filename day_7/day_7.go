package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	day7            = "puzzle.txt"
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
	highCard int 
	bid      int
	bestHand int
}

func GetBest(cards string) int {
	m := map[rune]int{}

	for _, r := range cards {
		m[r] += 1
	}

	switch len(m) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, v := range m {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
	case 3:
		for _, v := range m {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 5:
		return HIGH_CARD
	default:
		return ONE_PAIR
	}
	return 0
}

func GetHighCard(cards string) int{
	max := 0
	for _, r := range cards {
		if int(r) > max{
			max = int(r)
		}
	}
	return max
}

func main() {
	byteMap, _ := os.ReadFile(day7)
	puzzle := string(byteMap)
	lines := strings.Split(puzzle, "\n")
	hands := []hand{}
	score := 0

	for _, line := range lines {
		values := strings.Split(line, " ")
		bid, _ := strconv.Atoi(values[1])
		h := hand{cards: values[0], bid: bid}
		h.bestHand = GetBest(h.cards)
		h.highCard = GetHighCard(h.cards)
		hands = append(hands, h)
	}
	fmt.Printf("%#v\n", hands)
	sort.Slice(hands, func(i, j int) bool{
		if hands[i].bestHand == hands[j].bestHand{
			return hands[i].highCard > hands[j].highCard 
		}
		return hands[i].bestHand < hands[j].bestHand
	})
	
	for i, hand := range hands{
		fmt.Printf("%#v\n", hand)
		score += ((i+1) * hand.bid)
	}

	fmt.Printf("%#v\n", score)

}
