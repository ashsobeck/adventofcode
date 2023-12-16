package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Sequence struct {
	Sequence string
}

type Lens struct {
	Label    Sequence
	FocalLen int
}

func main() {
	file, _ := os.ReadFile("./day15.txt")
	f := func(c rune) bool {
		return c == ',' || unicode.IsSpace(c)
	}

	seqs := strings.FieldsFunc(string(file), f)

	var res int
	for _, sequence := range seqs {
		seq := Sequence{sequence}
		res += seq.Hash()
	}
	fmt.Println("Hash sum:", res)

	boxes := make(map[int][]Lens)
	for _, sequence := range seqs {
		seqStop := strings.IndexFunc(sequence, func(c rune) bool {
			return c == '=' || c == '-'
		})
		lens := Lens{Sequence{string(sequence[:seqStop])}, 0}
		hash := lens.Label.Hash()
		if strings.Contains(sequence, "=") {
			lens.FocalLen, _ = strconv.Atoi(string(sequence[len(sequence)-1]))
			exists := slices.IndexFunc(boxes[hash], func(e Lens) bool {
				return e.Label.Sequence == lens.Label.Sequence
			})
			if exists != -1 {
				boxes[hash][exists] = lens
			} else {
				boxes[hash] = append(boxes[hash], lens)
			}
		} else {
			boxes[hash] = slices.DeleteFunc(boxes[hash], func(e Lens) bool {
				return e.Label.Sequence == lens.Label.Sequence
			})
		}
	}

	var resP2 int
	for key, lenses := range boxes {
		for i, lens := range lenses {
			resP2 += (key + 1) * (i + 1) * lens.FocalLen
		}
	}
	fmt.Println("Focusing Power:", resP2)
}

func (s Sequence) Hash() int {
	var res int
	for _, r := range s.Sequence {
		res += int(r)
		res *= 17
		res = res % 256
	}
	return res
}
