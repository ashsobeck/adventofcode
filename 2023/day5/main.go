package main

import (
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./day5_test.txt")

	newLine := func(c rune) bool {
		return !(c == '\n')
	}

	almanac := strings.FieldsFunc(string(file), newLine)
	for _, 
}
