package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type AlmanacMap struct {
	DestStart   uint64
	SourceStart uint64
	RangeLength uint64
}

func main() {
	file, _ := os.ReadFile("./day5.txt")

	getNum := regexp.MustCompile(`\d+`)
	getSeedRange := regexp.MustCompile(`\d+.\d+`)

	newLine := func(c rune) bool {
		return c == '\n' || c == ':'
	}

	almanac := strings.FieldsFunc(string(file), newLine)
	var seeds []uint64
	var currentMap int
	var almanacMaps [][]AlmanacMap

	var mapping []AlmanacMap
	for i, info := range almanac {
		if i == 0 || i == 1 {
			rangeSeeds := getSeedRange.FindAllString(info, -1)
			// fmt.Println(rangeSeeds)
			for _, seedRange := range rangeSeeds {
				s := strings.Split(seedRange, " ")
				start, _ := strconv.ParseUint(s[0], 10, 64)
				numSeeds, _ := strconv.ParseUint(s[1], 10, 64)
				for j := uint64(0); j < numSeeds; j++ {
					seeds = append(seeds, start+j)
				}
			}
		} else {
			nums := getNum.FindAllString(info, -1)
			if nums == nil {
				currentMap++
				if currentMap > 1 {
					almanacMaps = append(almanacMaps, mapping)
				}
				mapping = make([]AlmanacMap, 0)
				continue
			}
			var vals []uint64
			for _, val := range nums {
				temp, _ := strconv.ParseUint(val, 10, 64)
				vals = append(vals, temp)
			}

			mapping = append(mapping, AlmanacMap{vals[0], vals[1], vals[2]})
		}
	}

	almanacMaps = append(almanacMaps, mapping)
	// fmt.Println(seeds)
	// fmt.Println(almanacMaps)

	var convertedVal uint64
	var convertedSeeds []uint64
	for _, s := range seeds {

		for i, convs := range almanacMaps {
			// fmt.Println("almanacs", convs, i)
			for j, mapping := range convs {
				// fmt.Println("mapping: ", mapping, j)
				if i == 0 {
					// fmt.Println(mapping.DestStart, mapping.SourceStart, mapping.DestStart+mapping.RangeLength)
					if s >= mapping.SourceStart && s < mapping.SourceStart+mapping.RangeLength {
						diff := s - mapping.SourceStart
						// fmt.Println("diff; ", diff)
						convertedVal = mapping.DestStart + diff
						break
					} else if j == len(convs)-1 {
						convertedVal = s
					}
				} else {
					// fmt.Println(mapping.DestStart, mapping.SourceStart, mapping.DestStart+mapping.RangeLength)
					if convertedVal >= mapping.SourceStart && convertedVal < mapping.SourceStart+mapping.RangeLength {
						diff := convertedVal - mapping.SourceStart
						// fmt.Println("diff; ", diff)
						convertedVal = mapping.DestStart + diff
						break
					}
				}
			}
			// fmt.Println(s, " -> ", convertedVal)
		}
		convertedSeeds = append(convertedSeeds, convertedVal)
		// fmt.Println(convertedSeeds)
	}

	fmt.Println("Min: ", slices.Min(convertedSeeds))
}
