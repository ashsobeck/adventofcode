package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cube struct {
	Count int
	Color string
}

var maxRed, maxGreen, maxBlue = Cube{Count: 12, Color: "red"}, Cube{Count: 13, Color: "green"}, Cube{Count: 14, Color: "blue"}

func main() {
	file, err := os.ReadFile("./day2.txt")
	findId := regexp.MustCompile(`\d+`)
	if err != nil {
		panic(err.Error())
	}

	games := strings.Split(string(file), "\n")

	var idSum int
	var powerSum int
	var gameValid bool
	for _, game := range games {
		if game == "" {
			continue
		}

		var localRMax, localGMax, localBMax = 0, 0, 0
		gameValid = true

		fmt.Println(game)

		idSet := strings.Split(game, ":")
		id, err := strconv.Atoi(findId.FindString(idSet[0]))

		if err != nil {
			panic(err.Error())
		}

		sets := strings.Split(idSet[1], ";")

		for _, set := range sets {
			fmt.Println(set)
			cubes := strings.Split(set, ",")

			for _, cubeShown := range cubes {
				// first element without doing [1:] is " "
				countColor := strings.Split(cubeShown, " ")[1:]
				fmt.Println(countColor)
				count, err := strconv.Atoi(countColor[0])

				if err != nil {
					panic(err.Error())
				}

				cube := Cube{count, countColor[1]}
				switch cube.Color {
				case "red":
					localRMax = max(cube.Count, localRMax)

				case "green":
					localGMax = max(cube.Count, localGMax)
				case "blue":
					localBMax = max(cube.Count, localBMax)
				}

				if gameValid = isSetPossible(cube); !gameValid {
					fmt.Println(gameValid)
				}
			}

		}
		powerSum += localRMax * localBMax * localGMax

		if gameValid {
			idSum += id
		}
	}
	fmt.Println("sum: ", idSum)
	fmt.Println("power sum: ", powerSum)
}

func isSetPossible(cube Cube) bool {
	switch cube.Color {
	case "red":
		return maxRed.Count >= cube.Count
	case "green":
		return maxGreen.Count >= cube.Count
	case "blue":
		return maxBlue.Count >= cube.Count
	default:
		return true
	}
}
