package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day4.txt")
	if err != nil {
		panic(err.Error())
	}

	cards := strings.Split(string(file), "\n")

	var points int
	var cardMatches []int
	var cardTotals []int

	for _, card := range cards {
		if card == "" {
			continue
		}

		nums := strings.Split(card, ": ")[1]

		cardTotals = append(cardTotals, 1)

		split := strings.Split(nums, " | ")
		winList, actual := split[0], split[1]

		winningNums := strings.Split(winList, " ")
		cardNums := strings.Split(actual, " ")

		var matches int
		for _, num := range cardNums {
			if num == "" {
				continue
			}

			if slices.Contains(winningNums, num) {
				matches++
			}
		}
		if matches > 0 {
			points += 1 << (matches - 1)
		}
		cardMatches = append(cardMatches, matches)
	}

	fmt.Println("points: ", points)

	for i := 0; i < len(cardTotals); i++ {
		if cardMatches[i] > 0 {
			for j := i + 1; j < i+1+cardMatches[i]; j++ {
				cardTotals[j] += cardTotals[i]
			}
		}
	}

	var totalCards int
	for _, val := range cardTotals {
		totalCards += val
	}
	fmt.Println("Total Cards: ", totalCards)
}
