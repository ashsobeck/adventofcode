package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Island struct {
	Grid        [][]Node
	HeatLoss    int
	Consecutive map[string]int
}

type Node struct {
	Val  int
	X, Y int
}

func main() {
	file, _ := os.ReadFile("./day17_test.txt")

	m := strings.Fields(string(file))

	var island Island
	for i, row := range m {
		var nodes []Node
		for j, c := range row {
			n, _ := strconv.Atoi(string(c))
			nodes = append(nodes, Node{n, i, j})

		}
		island.Grid = append(island.Grid, nodes)
	}
	island.Consecutive = make(map[string]int)
	fmt.Println(island.Grid[0][0])
	island.FindBestPath()
	fmt.Println("Total Heat Loss:", island.HeatLoss)

}

func (i *Island) FindBestPath() {
	var x, y int
	var dist string
	for {
		x, y, dist = i.NextNode(x, y, dist)
		fmt.Println(x, y, dist, i.Consecutive)
		node := i.Grid[x][y]
		i.HeatLoss += node.Val
		if node.X == len(i.Grid)-1 && node.Y == len(i.Grid[0])-1 {
			return
		}
		i.Consecutive[dist]++
		if x == 2 {
			return
		}
	}
}

func (is *Island) NextNode(i, j int, dist string) (int, int, string) {
	fmt.Println(i, j)
	if is.Consecutive[dist] == 3 {
		is.Consecutive[dist] = 0
		if dist == "left" || dist == "right" {
			up, down := -1, -1
			if i-1 >= 0 {
				up = int(is.Grid[i-1][j].Val)
			}
			if i+1 < len(is.Grid) {
				down = int(is.Grid[i+1][j].Val)
			}
			if (up != -1 && down-up >= 2) || down == -1 {
				return i - 1, j, "up"
			} else {
				return i + 1, j, "down"
			}
		} else if dist == "up" || dist == "down" {
			left, right := -1, -1
			if j-1 >= 0 {
				left = int(is.Grid[i][j-1].Val)
			}
			if j+1 < len(is.Grid) {
				right = int(is.Grid[i][j+1].Val)
			}
			if (left != -1 && right-left >= 2) || right == -1 {
				return i, j - 1, "left"
			} else {
				return i, j + 1, "right"
			}
		}
	} else {
		up, down, left, right := -1, -1, -1, -1
		if i-1 >= 0 {
			up = int(is.Grid[i-1][j].Val)
		}
		if i+1 < len(is.Grid) {
			down = int(is.Grid[i+1][j].Val)
		}
		if j-1 >= 0 {
			left = int(is.Grid[i][j-1].Val)
		}
		if j+1 < len(is.Grid[i]) {
			right = int(is.Grid[i][j+1].Val)
		}
		fmt.Println("up", up, "down", down, "left", left, "right", right)
		switch dist {
		case "left":
			if (down != -1 && left-down >= 2) || left == -1 {
				is.Consecutive[dist] = 0
				return i + 1, j, "down"
			} else if up != -1 && left-up >= 2 {
				is.Consecutive[dist] = 0
				return i - 1, j, "up"
			} else {
				return i, j - 1, "left"
			}

		case "right":
			if (down != -1 && right-down >= 2) || right == -1 {
				is.Consecutive[dist] = 0
				return i + 1, j, "down"
			} else if up != -1 && right-up >= 2 {
				is.Consecutive[dist] = 0
				return i - 1, j, "up"
			} else {
				return i, j + 1, "right"
			}
		case "up":
			if (right != -1 && up-right >= 2) || up == -1 {
				is.Consecutive[dist] = 0
				return i, j + 1, "right"
			} else if left != -1 && up-left >= 2 {
				is.Consecutive[dist] = 0
				return i, j - 1, "left"
			} else {
				return i - 1, j, "up"
			}
		case "down":
			if (right != -1 && down-right >= 2) || down == -1 {
				is.Consecutive[dist] = 0
				return i, j + 1, "right"
			} else if right == -1 && down-left >= 2 {
				is.Consecutive[dist] = 0
				return i, j - 1, "left"
			} else {
				return i + 1, j, "down"
			}
		}
	}

	return i, j + 1, "right"
}
