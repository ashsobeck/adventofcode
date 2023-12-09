package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	file, _ := os.ReadFile("./day8.txt")

	getLeftRight := regexp.MustCompile(`\w+`)

	f := func(r rune) bool {
		return r == '\n'
	}
	directions := strings.FieldsFunc(string(file), f)

	instructions := []rune(directions[0])

	nodes := make(map[string]Node)
	for _, n := range directions[1:] {
		key := strings.Split(n, " ")[0]
		leftRight := getLeftRight.FindAllString(n, -1)
		node := Node{Left: leftRight[1], Right: leftRight[2]}
		// TODO; part 2 with suffixes
		// strings.HasSuffix()
		nodes[key] = node
	}
	fmt.Println(instructions)
	fmt.Println(nodes)

	var steps int
	var current string
	goal := "ZZZ"
	step := "AAA"

	for current != goal {
		for _, direction := range instructions {
			fmt.Println(direction)

			steps++
			current = nodes[step].Step(direction)
			step = current
			fmt.Println(current)
			if current == goal {
				break
			}
		}
	}

	fmt.Println("Number of steps taken: ", steps)
}

func (n Node) Step(dir rune) string {
	if dir == 'L' {
		return n.Left

	} else if dir == 'R' {
		return n.Right
	}

	return n.Right
}
