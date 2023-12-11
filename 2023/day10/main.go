package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"unicode"
)

type Maze struct {
	Grid  [][]Node
	Steps int
}

type Node struct {
	Val     string
	Visited bool
}

func IsPipe(c string) bool {
	return c != "."
}

func main() {
	file, _ := os.ReadFile("./day10.txt")
	f := func(r rune) bool {
		return unicode.IsSpace(r)
	}

	grid := strings.FieldsFunc(string(file), f)

	var maze Maze
	for _, row := range grid {
		r := strings.Split(row, "")
		var nodes []Node
		for _, val := range r {
			nodes = append(nodes, Node{val, false})
		}
		maze.Grid = append(maze.Grid, nodes)
	}

	var startFound bool
	for i := range maze.Grid {
		for j, col := range maze.Grid[i] {
			if col.Val == "S" {
				maze.FindLoop(i, j, "Start", col.Val)
				maze.Grid[i][j].Visited = true
				startFound = true
				break
			}
		}
		if startFound {
			break
		}
	}

	fmt.Println("Total steps in loop: ", maze.Steps)
	fmt.Println("Furthest point in loop from S: ", math.Floor(float64(maze.Steps)/2))
}

func (m *Maze) FindLoop(i, j int, dir, prev string) {
	if i < 0 || j < 0 || i >= len(m.Grid) || j >= len(m.Grid[i]) {
		return
	}

	current := m.Grid[i][j].Val
	canConnect := CanConnect(prev, current, dir)

	if (!IsPipe(current) || !canConnect) && dir != "Start" || m.Grid[i][j].Visited {
		return
	} else if canConnect {
		m.Grid[i][j].Visited = true
		m.Steps++
	}

	if current == "S" && dir != "Start" {
		return
	}

	m.FindLoop(i-1, j, "N", current)
	m.FindLoop(i, j+1, "E", current)
	m.FindLoop(i+1, j, "S", current)
	m.FindLoop(i, j-1, "W", current)
}

func CanConnect(p, connection, dir string) bool {
	west := []string{"L", "-", "F", "S"}
	east := []string{"-", "J", "7", "S"}
	south := []string{"J", "|", "L", "S"}
	north := []string{"F", "|", "7", "S"}
	switch p {
	case "S":
		switch dir {
		case "N":
			return slices.Contains(north, connection)
		case "E":
			return slices.Contains(east, connection)
		case "S":
			return slices.Contains(south, connection)
		case "W":
			return slices.Contains(west, connection)
		}
	case "|":
		switch dir {
		case "N":
			return slices.Contains(north, connection)

		case "S":
			return slices.Contains(south, connection)
		}
	case "-":
		switch dir {
		case "E":
			return slices.Contains(east, connection)

		case "W":
			return slices.Contains(west, connection)
		}
	case "L":
		switch dir {
		case "N":
			return slices.Contains(north, connection)

		case "E":
			return slices.Contains(east, connection)
		}
	case "J":
		switch dir {
		case "N":
			return slices.Contains(north, connection)

		case "W":
			return slices.Contains(west, connection)
		}
	case "7":
		switch dir {
		case "S":
			return slices.Contains(south, connection)

		case "W":
			return slices.Contains(west, connection)
		}
	case "F":
		switch dir {
		case "E":
			return slices.Contains(east, connection)

		case "S":
			return slices.Contains(south, connection)
		}
	}
	return false
}
