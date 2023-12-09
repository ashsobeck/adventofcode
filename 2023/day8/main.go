package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
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

	var startNodes []string
	var goalNodes []string
	nodes := make(map[string]Node)
	for _, n := range directions[1:] {
		key := strings.Split(n, " ")[0]
		leftRight := getLeftRight.FindAllString(n, -1)
		node := Node{Left: leftRight[1], Right: leftRight[2]}

		if strings.HasSuffix(key, "A") {
			startNodes = append(startNodes, key)
		} else if strings.HasSuffix(key, "Z") {
			goalNodes = append(goalNodes, key)
		}

		nodes[key] = node
	}

	// fmt.Println(instructions)
	fmt.Println(nodes)
	fmt.Println(startNodes)
	fmt.Println(goalNodes)

	var nodeSteps []int
	var next string

	for i, step := range startNodes {
		var steps int
		next = step
		for !slices.Contains(goalNodes, startNodes[i]) {
			for _, direction := range instructions {
				steps++
				next = nodes[next].Step(direction)
				startNodes[i] = next

				if slices.Contains(goalNodes, startNodes[i]) {
					break
				}
			}
		}
		nodeSteps = append(nodeSteps, steps)
	}
	fmt.Println("Cycle for each node; ", nodeSteps)

	res := LCM(nodeSteps[0], nodeSteps[1], nodeSteps[2:]...)
	fmt.Println("Number of steps taken: ", res)
}

func (n Node) Step(dir rune) string {
	if dir == 'L' {
		return n.Left

	} else if dir == 'R' {
		return n.Right
	}

	return n.Right
}

// Greatest common denominator via Euclidean algorithm
// Retrieved from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, ints ...int) int {
	res := a * b / GCD(a, b)

	for i := 0; i < len(ints); i++ {
		res = LCM(res, ints[i])
	}

	return res
}
