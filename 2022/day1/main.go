package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sum(s string) int {
	sSlice := strings.Split(s, "\n")
	var sum int
	for _, v := range sSlice {
		if v == "" {
			continue
		}
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sum += val

	}
	return sum
}

func sumNums(nums []int) int {

	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func findHighestCals(cals []string) int {
	var maxCal int

	for _, v := range cals {
		localMax := sum(v)
		maxCal = int(math.Max(float64(localMax), float64(maxCal)))
	}
	return maxCal
}

func findTopThreeCals(cals []string) int {
	sumSlice := make([]int, len(cals))

	for _, v := range cals {
		sumSlice = append(sumSlice, sum(v))
	}

	slices.SortFunc(sumSlice, func(i, j int) int {
		if i > j {
			return -1
		} else if i < j {
			return 1
		}
		return 0
	})

	return sumNums(sumSlice[:3])

}

func main() {
	cals, err := os.ReadFile("./day1.txt")
	if err != nil {
		fmt.Println("Error reading in file... exiting")
		panic(err)
	}

	calSlice := strings.Split(string(cals), "\n\n")
	hiCal := findHighestCals(calSlice)

	fmt.Println(hiCal)
	fmt.Println(findTopThreeCals(calSlice))
}
