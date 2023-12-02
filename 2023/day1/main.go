package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	isNum := regexp.MustCompile(`\d`)
	// capture all digits left to right
	firstNum := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	// capture the right most digit
	// regexp does not support negative lookaheads
	lastNum := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|\d)`)

	file, err := os.ReadFile("./day1.txt")
	if err != nil {
		panic(err.Error())
	}

	calibrations := strings.Split(string(file), "\n")

	var sum int
	for _, calibration := range calibrations {
		first := firstNum.FindString(calibration)

		var last string
		if match := lastNum.FindStringSubmatch(calibration); len(match) > 0 {
			last = match[1]
		}

		// no initial match means empty string
		if first == "" {
			continue
		}

		// if the first/last digit is a word, convert it
		if word := isNum.FindString(first); word == "" {
			first = wordToNum(first)
		}
		if word := isNum.FindString(last); word == "" {
			last = wordToNum(last)
		}

		firstLast := first + last
		if last == "" {
			firstLast += first
		}

		calibrationVal, err := strconv.Atoi(firstLast)
		if err != nil {
			panic(err.Error())
		}
		sum += calibrationVal
	}

	fmt.Printf("Calibration Sum: %d\n", sum)
}

func wordToNum(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}
