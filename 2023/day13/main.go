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
	file, _ := os.ReadFile("./day13.txt")

	split := strings.Split(string(file), "\n")

	var patterns []Pattern
	var p Pattern
	for _, s := range split {
		if s == "" && len(p.Rows) > 0 {
			patterns = append(patterns, p)
			p = Pattern{}
		} else if s != "" {
			p.Rows = append(p.Rows, s)
		}
	}

	var res int
	for _, p := range patterns {
		p.GetCols()
		res += p.FindReflection()
	}

	fmt.Println("Reflection result:", res)
}

func (p Pattern) FindReflection() int {
	var res int
	if pointOfRef := p.CheckHorizontal(); pointOfRef != 0 {
		res = 100 * pointOfRef
	} else {
		res = p.CheckVertical()
	}
	if res == 0 {
		fmt.Println("no reflection")
		for _, rows := range p.Rows {
			fmt.Println(rows)
		}
	}
	return res

}

func (p Pattern) CheckVertical() int {
	// check vertical
	var prev string
	var left, right, pointOfRef int

	for i, col := range p.Columns {
		if i == 0 {
			prev = col
			continue
		}

		// point of reflection
		if prev == col {
			left = i - 2
			right = i + 1
			pointOfRef = i
			if p.CheckVerticalReflection(left, right) {
				break
			} else {
				pointOfRef = 0
			}
		}
		prev = col
	}

	return pointOfRef
}
func (p Pattern) CheckVerticalReflection(left, right int) bool {
	for left != -1 && right != len(p.Columns) {
		if p.Columns[left] != p.Columns[right] {
			return false
		}
		left--
		right++
	}
	return true
}

func (p Pattern) CheckHorizontal() int {
	// check horizontal
	var prev string
	var left, right, pointOfRef int

	for i, row := range p.Rows {
		if i == 0 {
			prev = row
			continue
		}

		// point of reflection
		if prev == row {
			left = i - 2
			right = i + 1
			pointOfRef = i

			if p.CheckHorizontalReflection(left, right) {
				break
			} else {
				pointOfRef = 0
			}
		}
		prev = row
	}
	return pointOfRef
}

func (p Pattern) CheckHorizontalReflection(left, right int) bool {
	for left != -1 && right != len(p.Rows) {
		if p.Rows[left] != p.Rows[right] {
			return false
		}
		left--
		right++
	}
	return true
}

func (p *Pattern) GetCols() {
	for j := range p.Rows[0] {
		var cols string
		for i := range p.Rows {
			cols += string(p.Rows[i][j])
		}
		p.Columns = append(p.Columns, cols)
	}
}
