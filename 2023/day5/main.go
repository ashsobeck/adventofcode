package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type AlmanacMap struct {
	DestStart   uint64
	SourceStart uint64
	RangeLength uint64
}

func main() {
	file, _ := os.ReadFile("./day5_test.txt")

	getNum := regexp.MustCompile(`\d+`)

	newLine := func(c rune) bool {
		return c == '\n' || c == ':'
	}

	almanac := strings.FieldsFunc(string(file), newLine)
	fmt.Println(almanac)
	var seeds []string
	var currentMap int
	almanacMaps := make(map[int][]AlmanacMap)

	for _, info := range almanac {
		nums := getNum.FindAllString(info, -1)
		if nums == nil {
			currentMap++
			continue
		}
		if currentMap == 1 {
			seeds = append(seeds, nums...)
		} else {
			var vals []uint64
			for _, val := range nums {
				temp, _ := strconv.ParseUint(val, 10, 64)
				vals = append(vals, temp)
			}

			almanacMaps[currentMap] = append(almanacMaps[currentMap], AlmanacMap{vals[0], vals[1], vals[2]})
		}
	}
	fmt.Println(seeds)
	fmt.Println(almanacMaps)

	var convertedVal uint64
	for _, seed := range seeds {
		for _, convs := range almanacMaps {
			var inRange bool

		}
	}
}
