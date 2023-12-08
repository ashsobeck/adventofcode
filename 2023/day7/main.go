package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Hand struct {
	Cards string
	Bid   int
	Rank  int
}

func main() {
	file, _ := os.ReadFile("./day7.txt")

	f := func(c rune) bool {
		return unicode.IsSpace(c)
	}

	game := strings.FieldsFunc(string(file), f)

	var hands []Hand
	for i, cardsBid := range game {
		// hand data
		if i%2 == 0 {
			hand := Hand{Cards: cardsBid}
			hands = append(hands, hand)
		} else {
			bid, _ := strconv.Atoi(cardsBid)
			hands[len(hands)-1].Bid = bid
		}
	}
	p1Hands := hands

	for i, hand := range p1Hands {
		pairs := FindHandType(hand, 0)

		p1Hands[i].Rank = GetRank(pairs, hand, len(p1Hands), 0)
	}

	sort.Slice(p1Hands, func(i, j int) bool {
		if p1Hands[i].Rank == p1Hands[j].Rank {
			return CompareHands(p1Hands[i], p1Hands[j], 0)

		}
		return p1Hands[i].Rank < p1Hands[j].Rank
	})

	var res int
	for i, hand := range p1Hands {
		res += hand.Bid * (i + 1)
	}
	fmt.Println("Total Winnings P1: ", res)

	for i, hand := range hands {
		pairs := FindHandType(hand, 1)

		hands[i].Rank = GetRank(pairs, hand, len(hands), 1)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Rank == hands[j].Rank {
			return CompareHands(hands[i], hands[j], 1)

		}
		return hands[i].Rank < hands[j].Rank
	})

	for _, val := range hands {
		fmt.Println(val)
	}

	var resP2 int
	for i, hand := range hands {
		resP2 += hand.Bid * (i + 1)
	}
	fmt.Println("Total Winnings P2: ", resP2)
}

func GetRank(pairs int, hand Hand, l int, t int) int {
	var rank int
	switch pairs {
	case 5:
		rank = 7
	case 4:
		rank = 6
	case 3:
		if IsFullHouse(hand, t) {
			rank = 5

		} else {
			rank = 4
		}
	case 2:
		numPairs := GetPairType(hand)
		if numPairs == 2 {
			rank = 3
		} else {
			rank = 2
		}
	case 1:
		rank = 1
	}
	return rank
}

func RunicTranform(hand Hand, t int) []rune {
	splitHand := strings.Map(func(r rune) rune {
		if r == 'A' {
			return 'Z'
		}
		if r == 'K' {
			return 'Y'
		}
		if r == 'Q' {
			return 'X'
		}
		if r == 'J' && t == 0 {
			return 'W'
		}
		if r == 'J' && t == 1 {
			return '0'
		}
		if r == 'T' {
			return 'V'
		}
		return r
	}, hand.Cards)

	return []rune(splitHand)
}

func CompareHands(i Hand, j Hand, t int) bool {
	left := RunicTranform(i, t)
	right := RunicTranform(j, t)
	for k := 0; k < len(left); k++ {
		if left[k] > right[k] {
			return false
		} else if left[k] < right[k] {
			return true
		}
	}
	return true
}

func FindHandType(hand Hand, t int) int {
	pairs := GetHandSet(hand)
	var most int
	for key, val := range pairs {
		if t == 0 && val > most {
			most = val
		} else if t == 1 && val > most && key != "J" {
			most = val
		}
	}
	if t == 0 {
		return most
	} else {
		return most + pairs["J"]
	}
}

func IsFullHouse(hand Hand, t int) bool {
	pairs := GetHandSet(hand)

	var foundTwo bool
	var foundThree bool
	var numTwos int
	for _, card := range pairs {
		switch card {
		case 2:
			foundTwo = !foundTwo
			numTwos++
		case 3:
			foundThree = !foundThree
		default:
			continue
		}
	}

	if t == 1 && numTwos == 2 && pairs["J"] == 1 {
		return true
	}

	return foundTwo && foundThree
}

func GetPairType(hand Hand) int {
	pairs := GetHandSet(hand)

	var numPairs int
	for _, pair := range pairs {
		if pair == 2 {
			numPairs++
		}
	}
	return numPairs
}

func GetHandSet(hand Hand) map[string]int {
	pairs := make(map[string]int)
	splitHand := strings.Split(hand.Cards, "")
	for _, card := range splitHand {
		pairs[card]++
	}
	return pairs
}
