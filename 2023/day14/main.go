package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"regexp"
	"strings"
)

type Platform struct {
	Grid [][]rune
}

func main() {
	file, _ := os.ReadFile("./day14.txt")

	rows := strings.Fields(string(file))

	var platform Platform
	var platform2 Platform
	for i, r := range rows {
		platform.Grid = append(platform.Grid, []rune(r))
		platform2.Grid = append(platform2.Grid, []rune(r))
		for j, c := range r {
			if c == 'O' {
				for k := i - 1; k >= 0; k-- {
					if platform.Grid[k][j] == '.' {
						platform.Grid[k][j] = 'O'
						platform.Grid[k+1][j] = '.'
					} else if platform.Grid[k][j] == '#' || platform.Grid[k][j] == 'O' {
						break
					}
				}
			}
		}
	}
	res := platform.CalculateLoad()
	fmt.Println("Total Load:", res)

	hashes := make(map[uint64]int)
	cycles := make(map[int]int)
	var cycle []int
	var cycleStart int
	for i := 0; i < 1000; i++ {
		platform2.Cycle(false)
		if hash := platform2.Hash(); hashes[hash] > 0 {
			hashes[hash]++
			if hashes[hash] == 2 {
				cycle = append(cycle, platform2.CalculateLoad())
				cycles[i] = platform2.CalculateLoad()
			} else if hashes[hash] == 3 && cycleStart == 0 {
				cycleStart = i - len(cycle)
				fmt.Println(cycle, cycleStart, len(cycle))
			}
		} else {
			hashes[hash]++
		}
	}
	idx := (((1000000000 - cycleStart) % len(cycle)) - 1) + cycleStart

	fmt.Println("Total Load After 1,000,000,000 Cycles:", cycles[idx])
}

func (p Platform) Hash() uint64 {
	h := fnv.New64a()
	for _, row := range p.Grid {
		h.Write([]byte(string(row)))
	}
	return h.Sum64()
}
func (p Platform) CalculateLoad() int {
	var res int
	numO := regexp.MustCompile(`O`)
	numRows := len(p.Grid)
	for i, row := range p.Grid {
		roundRocks := len(numO.FindAllString(string(row), -1))
		res += roundRocks * (numRows - i)
	}
	return res
}

func (p *Platform) Cycle(print bool) {
	p.RotateNorth()
	if print {
		fmt.Println("N")
		for i := range p.Grid {
			fmt.Println(string(p.Grid[i]))
		}
	}
	p.RotateWest()
	if print {
		fmt.Println("W")
		for i := range p.Grid {
			fmt.Println(string(p.Grid[i]))
		}
	}
	p.RotateSouth()
	if print {
		fmt.Println("S")
		for i := range p.Grid {
			fmt.Println(string(p.Grid[i]))
		}
	}
	p.RotateEast()
	if print {
		fmt.Println("E")
		for i := range p.Grid {
			fmt.Println(string(p.Grid[i]))
		}
	}
}

func (p *Platform) RotateNorth() {
	for i, r := range p.Grid {
		for j, c := range r {
			if c == 'O' {
				for k := i - 1; k >= 0; k-- {
					if p.Grid[k][j] == '.' {
						p.Grid[k][j] = 'O'
						p.Grid[k+1][j] = '.'
					} else if p.Grid[k][j] == '#' || p.Grid[k][j] == 'O' {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) RotateWest() {
	for i, r := range p.Grid {
		for j, c := range r {
			if c == 'O' {
				for k := j - 1; k >= 0; k-- {
					if p.Grid[i][k] == '.' {
						p.Grid[i][k] = 'O'
						p.Grid[i][k+1] = '.'
					} else if p.Grid[i][k] == '#' || p.Grid[i][k] == 'O' {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) RotateSouth() {
	for i := len(p.Grid) - 1; i >= 0; i-- {
		r := p.Grid[i]
		for j, c := range r {
			if c == 'O' {
				for k := i + 1; k < len(p.Grid); k++ {
					if p.Grid[k][j] == '.' {
						p.Grid[k][j] = 'O'
						p.Grid[k-1][j] = '.'
					} else if p.Grid[k][j] == '#' || p.Grid[k][j] == 'O' {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) RotateEast() {
	for i, r := range p.Grid {
		for j := len(r) - 1; j >= 0; j-- {
			if r[j] == 'O' {
				for k := j + 1; k < len(r); k++ {
					if p.Grid[i][k] == '.' {
						p.Grid[i][k] = 'O'
						p.Grid[i][k-1] = '.'
					} else if p.Grid[i][k] == '#' || p.Grid[i][k] == 'O' {
						break
					}
				}
			}
		}
	}
}
