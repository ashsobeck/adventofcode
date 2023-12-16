package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Node struct {
	Char    rune
	Visited bool
}
type Universe struct {
	Grid       [][]Node
	Galaxies   []GalaxyPair
	GalaxyList []int
}

type GalaxyPair struct {
	Start         int
	End           int
	Distance      int
	StartLocation GalaxyLocation
	EndLocation   GalaxyLocation
}

type GalaxyLocation struct {
	Number int
	I      int
	J      int
}

type SearchNode struct {
	N Node
	I int
	J int
}

func main() {
	file, _ := os.ReadFile("./day11.txt")
	u := strings.Fields(string(file))

	var universe Universe
	var numGalaxies int
	for _, row := range u {
		runes := CreateNodeList([]rune(row))
		universe.Grid = append(universe.Grid, runes)
		if slices.Contains(runes, Node{'#', false}) {
			for _, r := range runes {
				if r.Char == '#' {
					universe.GalaxyList = append(universe.GalaxyList, numGalaxies)
					numGalaxies++
				}
			}
		} else {
			universe.Grid = append(universe.Grid, runes)
		}
	}
	// for i := range universe.Grid {
	// 	fmt.Println(universe.Grid[i])
	// }
	fmt.Println(u[0])
	var colsUpdated int
	for i := 0; i < len(u[0]); i++ {
		var col []rune
		for _, val := range u {
			col = append(col, rune(val[i]))
		}
		if !slices.Contains(col, '#') {
			fmt.Println("no # in col", i)
			for j := range universe.Grid {
				universe.Grid[j] = slices.Insert(universe.Grid[j], i+colsUpdated, Node{'.', false})
			}
			colsUpdated++
		}
	}

	locations := universe.MarkGalaxies()
	for i, n := range universe.GalaxyList {
		for j := i + 1; j < len(universe.GalaxyList); j++ {
			if i == j {
				continue
			} else {
				pair := GalaxyPair{n, universe.GalaxyList[j], 0, locations[i], locations[j]}
				universe.Galaxies = append(universe.Galaxies, pair)
			}
		}
	}
	// for i := range universe.Grid {
	// 	fmt.Println(universe.Grid[i])
	// }

	// fmt.Println(len(universe.Galaxies))
	// for _, pair := range universe.Galaxies {
	// 	// fmt.Println(pair.Start, pair.End)
	// }
	// fmt.Println(universe.GalaxyList)

	var res int
	for i, pair := range universe.Galaxies {
		fmt.Println(pair.Start, pair.End)
		res += universe.ExploreGalaxyPair(pair.StartLocation.I, pair.StartLocation.J, i)
	}
	fmt.Println(res)
}

func CreateNodeList(runes []rune) []Node {
	var nodes []Node
	for _, r := range runes {
		nodes = append(nodes, Node{r, false})
	}
	return nodes
}

func (u *Universe) MarkGalaxies() []GalaxyLocation {
	var numGalaxies int
	var locations []GalaxyLocation
	for r := range u.Grid {
		for c := range u.Grid[r] {
			if u.Grid[r][c].Char == '#' {
				u.Grid[r][c].Char = rune(numGalaxies)
				locations = append(locations, GalaxyLocation{numGalaxies, r, c})
				numGalaxies++
			}
		}
	}
	return locations
}

func (u Universe) ExploreGalaxyPair(i, j, pair int) int {
	var q []SearchNode
	visited := make(map[SearchNode]bool)
	distance := make(map[SearchNode]int)
	q = append(q, SearchNode{u.Grid[i][j], i, j})

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.N.Char == rune(u.Galaxies[pair].End) {
			print(distance[node])
			return distance[node]
		}
		neighbors := u.GetAdjacentNodes(node.I, node.J)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				distance[neighbor] = distance[node] + 1
				q = append(q, neighbor)
			}
		}
	}
	fmt.Println(distance)
	return 0
}

func (u Universe) GetAdjacentNodes(i, j int) []SearchNode {
	var nodes []SearchNode
	if i-1 >= 0 && !u.Grid[i-1][j].Visited {
		nodes = append(nodes, SearchNode{u.Grid[i-1][j], i - 1, j})
	}
	if i+1 < len(u.Grid) && !u.Grid[i+1][j].Visited {
		nodes = append(nodes, SearchNode{u.Grid[i+1][j], i + 1, j})
	}
	if j-1 >= 0 && !u.Grid[i][j-1].Visited {
		nodes = append(nodes, SearchNode{u.Grid[i][j-1], i, j - 1})
	}
	if j+1 < len(u.Grid[i]) && !u.Grid[i][j+1].Visited {
		nodes = append(nodes, SearchNode{u.Grid[i][j+1], i, j + 1})
	}
	return nodes
}
