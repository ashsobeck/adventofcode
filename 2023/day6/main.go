package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./day6_test.txt")

	getNums := regexp.MustCompile(`\d+`)

	f := func(c rune) bool {
		return c == '\n'
	}

	raceInfo := strings.FieldsFunc(string(file), f)

	var times []int
	var distances []int

	timeInfo := getNums.FindAllString(raceInfo[0], -1)
	for _, time := range timeInfo {
		t, _ := strconv.Atoi(time)
		times = append(times, t)
	}

	recordInfo := getNums.FindAllString(raceInfo[1], -1)
	for _, record := range recordInfo {
		d, _ := strconv.Atoi(record)
		distances = append(distances, d)
	}
	fmt.Println(times)
	fmt.Println(distances)

	var wins []int
	for i, time := range times {
		fmt.Println(time)
		fmt.Println(distances[i])
		for j := 0; j < time; j++ {
			if dist := (time - j) * j; dist > distances[i] {
				fmt.Println("win", dist)
				totalWays := (time - j) - 2
				fmt.Println(totalWays)
				wins = append(wins, totalWays)
				break
			}

		}
	}

	var res int = 1
	for _, win := range wins {
		res *= win
	}
	fmt.Println("Ways to win multiplied: ", res)
}
