package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Contraption struct {
	Grid [][]Node
}

type Node struct {
	Char      rune
	Energized bool
}

type Point struct {
	X, Y int
}

var points map[Point][]string

func main() {
	file, _ := os.ReadFile("./day16.txt")

	s := strings.Fields(string(file))
	points = make(map[Point][]string)

	var contraption Contraption
	for _, row := range s {
		var nodes []Node
		for _, c := range row {
			nodes = append(nodes, Node{c, false})
		}
		contraption.Grid = append(contraption.Grid, nodes)
	}
	contraption.BounceBeam(0, 0, "right")
	fmt.Println("Energized Cells:", contraption.GetEnergizedCells())
	contraption.ResetGrid()

	var maxEnergy int
	for i := range contraption.Grid {
		for j := range contraption.Grid[i] {
			if j == 0 {
				contraption.BounceBeam(i, j, "right")
				maxEnergy = int(math.Max(float64(maxEnergy), float64(contraption.GetEnergizedCells())))
				contraption.ResetGrid()
			} else if j == len(contraption.Grid[i])-1 {
				contraption.BounceBeam(i, j, "left")
				maxEnergy = int(math.Max(float64(maxEnergy), float64(contraption.GetEnergizedCells())))
				contraption.ResetGrid()
			}
			if i == 0 {
				contraption.BounceBeam(i, j, "down")
				maxEnergy = int(math.Max(float64(maxEnergy), float64(contraption.GetEnergizedCells())))
				contraption.ResetGrid()
			} else if i == len(contraption.Grid)-1 {
				contraption.BounceBeam(i, j, "up")
				maxEnergy = int(math.Max(float64(maxEnergy), float64(contraption.GetEnergizedCells())))
				contraption.ResetGrid()
			}

			contraption.ResetGrid()
		}
	}
	fmt.Println("Max Energized Cells:", maxEnergy)
}

func (c *Contraption) BounceBeam(i, j int, dist string) {
	if i < 0 || j < 0 || i >= len(c.Grid) || j >= len(c.Grid[i]) || slices.Contains(points[Point{i, j}], dist) {
		return
	}

	c.Grid[i][j].Energized = true
	points[Point{i, j}] = append(points[Point{i, j}], dist)

	if c.Grid[i][j].Char == '.' {
		ni, nj := GetCoords(i, j, dist)
		c.BounceBeam(ni, nj, dist)
	}

	if c.Grid[i][j].Char == '|' {
		if dist == "left" || dist == "right" {
			downI, downJ := GetCoords(i, j, "down")
			upI, upJ := GetCoords(i, j, "up")

			c.BounceBeam(downI, downJ, "down")
			c.BounceBeam(upI, upJ, "up")
		} else if dist == "down" || dist == "up" {
			ni, nj := GetCoords(i, j, dist)
			c.BounceBeam(ni, nj, dist)
		}
	}

	if c.Grid[i][j].Char == '-' {
		if dist == "down" || dist == "up" {
			rightI, rightJ := GetCoords(i, j, "right")
			leftI, leftJ := GetCoords(i, j, "left")

			c.BounceBeam(rightI, rightJ, "right")
			c.BounceBeam(leftI, leftJ, "left")
		} else if dist == "left" || dist == "right" {
			ni, nj := GetCoords(i, j, dist)
			c.BounceBeam(ni, nj, dist)
		}
	}

	if c.Grid[i][j].Char == '/' {
		switch dist {
		case "right":
			upI, upJ := GetCoords(i, j, "up")
			c.BounceBeam(upI, upJ, "up")
		case "left":
			downI, downJ := GetCoords(i, j, "down")
			c.BounceBeam(downI, downJ, "down")
		case "up":
			rightI, rightJ := GetCoords(i, j, "right")
			c.BounceBeam(rightI, rightJ, "right")
		case "down":
			leftI, leftJ := GetCoords(i, j, "left")
			c.BounceBeam(leftI, leftJ, "left")
		}
	}

	if c.Grid[i][j].Char == '\\' {
		switch dist {
		case "left":
			upI, upJ := GetCoords(i, j, "up")
			c.BounceBeam(upI, upJ, "up")
		case "right":
			downI, downJ := GetCoords(i, j, "down")
			c.BounceBeam(downI, downJ, "down")
		case "down":
			rightI, rightJ := GetCoords(i, j, "right")
			c.BounceBeam(rightI, rightJ, "right")
		case "up":
			leftI, leftJ := GetCoords(i, j, "left")
			c.BounceBeam(leftI, leftJ, "left")
		}
	}
}

func GetCoords(i, j int, dist string) (int, int) {
	switch dist {
	case "left":
		return i, j - 1
	case "right":
		return i, j + 1
	case "up":
		return i - 1, j
	case "down":
		return i + 1, j
	}
	return i, j
}

func (c Contraption) GetEnergizedCells() int {
	var n int
	for i := range c.Grid {
		for _, node := range c.Grid[i] {
			if node.Energized {
				n++
			}
		}
	}
	return n
}

func (c *Contraption) ResetGrid() {
	points = make(map[Point][]string)
	for i := range c.Grid {
		for j := range c.Grid[i] {
			c.Grid[i][j].Energized = false
		}
	}
}
