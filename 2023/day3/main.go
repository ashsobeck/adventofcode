package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isSymbol = regexp.MustCompile(`[^A-Za-z0-9.\d]`)
var isDigit = regexp.MustCompile(`\d`)

func main() {
	file, err := os.ReadFile("./day3.txt")
	if err != nil {
		panic(err.Error())
	}

	schematicRows := strings.Split(string(file), "\n")
	schematic := make([][]string, len(schematicRows))

	// build out schematic grid for DFS
	for i, row := range schematicRows {
		fmt.Println(row)
		schematic[i] = make([]string, len(row))
		schematic[i] = strings.Split(row, "")
	}

	var partSum int
	var gearRatioSum int
	for i, row := range schematic {
		for j, val := range row {
			if isSymbol.FindString(val) != "" {
				fmt.Println(val)
				parts, gears := searchForPart(i, j, schematic, val)
				partSum += parts
				gearRatioSum += gears
			}
		}
	}

	fmt.Println("part sum: ", partSum)
	fmt.Println("gear sum: ", gearRatioSum)
}

func searchForPart(i int, j int, schematic [][]string, symbol string) (int, int) {
	if i < 0 || j < 0 || i > len(schematic) || j > len(schematic[i]) {
		return 0, 0
	}

	var adjacentParts int
	var numAdjacent int
	gears := make(map[int]int)
	isGear := symbol == "*"
	fmt.Println("is gear", isGear)

	// top row
	var topLeft int
	if topLeft = findPartNums(i-1, j-1, schematic); topLeft != 0 {
		fmt.Println("top left")
		adjacentParts += topLeft
		numAdjacent++
		gears[numAdjacent] = topLeft
	}
	var above int
	if above = findPartNums(i-1, j, schematic); above != 0 && topLeft == 0 {
		fmt.Println("above")
		adjacentParts += above
		numAdjacent++
		gears[numAdjacent] = above
	}

	if topRight := findPartNums(i-1, j+1, schematic); topRight != 0 && above == 0 {
		fmt.Println("top right")
		adjacentParts += topRight
		numAdjacent++
		gears[numAdjacent] = topRight
	}

	// left
	if left := findPartNums(i, j-1, schematic); left != 0 {
		adjacentParts += left
		numAdjacent++
		gears[numAdjacent] = left
	}
	// right
	if right := findPartNums(i, j+1, schematic); right != 0 {
		adjacentParts += right
		numAdjacent++
		gears[numAdjacent] = right
	}
	// bottom row
	var bottomLeft int
	if bottomLeft = findPartNums(i+1, j-1, schematic); bottomLeft != 0 {
		fmt.Println("bottom left")
		adjacentParts += bottomLeft
		numAdjacent++
		gears[numAdjacent] = bottomLeft
	}

	var below int
	if below = findPartNums(i+1, j, schematic); below != 0 && bottomLeft == 0 {
		fmt.Println("bewlo")
		adjacentParts += below
		numAdjacent++
		gears[numAdjacent] = below
	}

	if bottomRight := findPartNums(i+1, j+1, schematic); bottomRight != 0 && below == 0 {
		fmt.Println("bottom right")
		adjacentParts += bottomRight
		numAdjacent++
		gears[numAdjacent] = bottomRight
	}

	fmt.Println("num adj: ", numAdjacent)
	var gearRatio int
	if isGear && numAdjacent == 2 {
		res := 1
		for _, num := range gears {
			res *= num
		}
		gearRatio = res
	}

	return adjacentParts, gearRatio
}

func findPartNums(i int, j int, schematic [][]string) int {
	if i < 0 || j < 0 || i > len(schematic) || j > len(schematic[i]) {
		return 0
	}

	var partNum string
	if partNum = isDigit.FindString(schematic[i][j]); partNum == "" {
		return 0
	}

	if fullPartNum, err := getPartNum(i, j, schematic, partNum); err == nil {
		return fullPartNum
	} else {
		panic(err.Error())
	}
}

func getPartNum(i int, j int, schematic [][]string, startingNum string) (int, error) {
	if i < 0 || j < 0 || i > len(schematic) || j > len(schematic[i]) {
		return 0, nil
	}
	fmt.Println(schematic[i])

	var leftPart string
	for left := j - 1; left >= 0; left-- {
		if num := isDigit.FindString(schematic[i][left]); num == "" {
			break
		} else {
			leftPart = num + leftPart
		}
	}

	var rightPart string
	for right := j + 1; right < len(schematic[i]); right++ {
		if num := isDigit.FindString(schematic[i][right]); num == "" {
			break
		} else {
			rightPart += num
		}
	}

	fmt.Println(leftPart + startingNum + rightPart)
	return strconv.Atoi(leftPart + startingNum + rightPart)
}
