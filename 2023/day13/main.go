package main

import (
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	Rows    []string
	Columns []string
}

func main() {
	file, _ := os.ReadFile("./day13_test.txt")

	split := strings.Split(string(file), "\n")

	var patterns []Pattern
	var p Pattern
	for _, s := range split {
		if s == "" {
			patterns = append(patterns, p)
			p = Pattern{}
		}
		p.Rows = append(p.Rows, s)
	}
	for _, p := range patterns {
		p.GetCols()
	}

	for col := range patterns[0].Columns {
		fmt.Println(patterns[0].Columns[col])
	}
}

func (p Pattern) FindReflection() int {
	if pointOfRef := p.CheckHorizontal(); pointOfRef != -1 {
		return 100 * pointOfRef
	}

	return p.CheckVertical()
}

func (p Pattern) CheckHorizontal() int {
	// check horizontal
	var prev string
	var left, right, pointOfRef int

	for i, col := range p.Columns {
		if i == 0 {
			prev = col
			continue
		}

		// point of reflection
		if prev == col {
			left = i - 1
			right = i + 1
			pointOfRef = i
			break
		}
		prev = col
	}
	for left != -1 || right != len(p.Columns) {
		if p.Columns[left] != p.Columns[right] {
			return -1
		}
		left--
		right++
	}
	return pointOfRef
}

func (p Pattern) CheckVertical() int {
	// check vertical
	var prev string
	var left, right, pointOfRef int

	for i, row := range p.Rows {
		if i == 0 {
			prev = row
			continue
		}

		// point of reflection
		if prev == row {
			left = i - 1
			right = i + 1
			pointOfRef = i
			break
		}
		prev = row
	}
	for left != -1 || right != len(p.Rows) {
		if p.Rows[left] != p.Rows[right] {
			return -1
		}
		left--
		right++
	}
	return pointOfRef

}

func (p *Pattern) GetCols() {
	for j := range p.Rows[0] {
		var cols []string
		for i := range p.Rows {
			cols = append(cols, string(p.Rows[i][j]))
		}
		p.Columns = append(p.Columns, cols...)
	}
}
