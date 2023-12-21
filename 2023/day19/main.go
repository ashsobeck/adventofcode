package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Left     string
	Operator string
	Right    int
	Next     string
}

type Workflow struct {
	Name  string
	Rules []Rule
}

type Part struct {
	X, M, A, S int
}

type Range struct {
	Upper int
	Lower int
}

type PartRange struct {
	X, M, A, S Range
}

type Schematic struct {
	Workflows map[string]Workflow
}

func main() {
	file, _ := os.ReadFile("./day19_test.txt")
	nums := regexp.MustCompile(`\d+`)
	getNameAndRules := func(c rune) bool {
		return c == '{' || c == '}'
	}

	plan := strings.Fields(string(file))

	workflows := make(map[string]Workflow)
	var parts []Part
	for _, row := range plan {
		if row[0] == '{' {
			digits := nums.FindAllString(row, -1)
			var x, m, a, s int

			x, _ = strconv.Atoi(digits[0])
			m, _ = strconv.Atoi(digits[1])
			a, _ = strconv.Atoi(digits[2])
			s, _ = strconv.Atoi(digits[3])

			parts = append(parts, Part{x, m, a, s})
		} else {
			nameRule := strings.FieldsFunc(row, getNameAndRules)
			r := nameRule[1]
			var rules []Rule
			for _, rule := range strings.Split(r, ",") {
				rules = append(rules, GetRule(rule))
			}
			workflows[nameRule[0]] = Workflow{nameRule[0], rules}
		}
	}

	var res int
	for _, p := range parts {
		flow := workflows["in"]
		var completed bool
		for !completed {
			for _, rule := range flow.Rules {
				if rule.Operator == "" {
					if rule.Left == "A" {
						completed = true
						res += p.RatingSum()
						break
					} else if rule.Left == "R" {
						completed = true
						break
					}
					flow = workflows[rule.Left]
					break
				}

				if pass := EvaluateRule(rule, p); pass {
					if rule.Next == "A" {
						completed = true
						res += p.RatingSum()
						break
					} else if rule.Next == "R" {
						completed = true
						break
					}
					flow = workflows[rule.Next]
					break
				}
			}
		}
	}
	fmt.Println("Rating sum of all accepted parts:", res)
	ranges := PartRange{Range{1, 4000}, Range{1, 4000}, Range{1, 4000}, Range{1, 4000}}
	flow := workflows["in"]
	// schema := Schematic{workflows}
	// var validRanges []PartRange
	fmt.Println(EvaluateRange(flow.Rules[0], ranges))
}

func (s *Schematic) WalkRanges(name string, parts PartRange) []PartRange {
	for _, rule := range s.Workflows[name].Rules {
		fmt.Println(rule)
		// TODO: go down left and right path (true and false ranges of rules)
	}
	var list []PartRange
	return list
}

func (p Part) RatingSum() int {
	return p.X + p.M + p.A + p.S
}

func EvaluateRange(r Rule, p PartRange) (*PartRange, *PartRange, bool) {
	switch r.Left {
	case "A":
		return &p, nil, true
	case "R":
		return nil, nil, true
	}

	tRange := PartRange{p.X, p.M, p.A, p.S}
	fRange := PartRange{p.X, p.M, p.A, p.S}
	switch r.Operator {
	case ">":
		switch r.Left {
		case "x":
			tRange.X.Lower = r.Right
			tRange.X.Upper = p.X.Upper

			fRange.X.Upper = r.Right - 1
			fRange.X.Lower = p.X.Lower
		case "m":
			tRange.M.Lower = r.Right
			tRange.M.Upper = p.M.Upper

			fRange.M.Upper = r.Right - 1
			fRange.M.Lower = p.M.Lower
		case "a":
			tRange.A.Lower = r.Right
			tRange.A.Upper = p.A.Upper

			fRange.A.Upper = r.Right - 1
			fRange.A.Lower = p.A.Lower
		case "s":
			tRange.S.Lower = r.Right
			tRange.S.Upper = p.S.Upper

			fRange.S.Upper = r.Right - 1
			fRange.S.Lower = p.S.Lower
		}

	case "<":
		switch r.Left {
		case "x":
			tRange.X.Upper = r.Right - 1
			tRange.X.Lower = p.X.Lower

			fRange.X.Lower = r.Right
			fRange.X.Upper = p.X.Upper
		case "m":
			tRange.M.Upper = r.Right - 1
			tRange.M.Lower = p.M.Lower

			fRange.M.Lower = r.Right
			fRange.M.Upper = p.M.Upper
		case "a":
			tRange.A.Upper = r.Right - 1
			tRange.A.Lower = p.A.Lower

			fRange.A.Lower = r.Right
			fRange.A.Upper = p.A.Upper
		case "s":
			tRange.S.Upper = r.Right - 1
			tRange.S.Lower = p.S.Lower

			fRange.S.Lower = r.Right
			fRange.S.Upper = p.S.Upper
		}
	}
	return &tRange, &fRange, false
}

func EvaluateRule(r Rule, p Part) bool {
	var left int
	switch r.Left {
	case "x":
		left = p.X
	case "m":
		left = p.M
	case "a":
		left = p.A
	case "s":
		left = p.S
	default:
		return true
	}

	switch r.Operator {
	case ">":
		return left > r.Right
	case "<":
		return left < r.Right
	}

	return false
}

func GetRule(s string) Rule {
	operator := func(c rune) bool {
		return c == '>' || c == '<'
	}

	leftRight := strings.FieldsFunc(s, operator)
	opIdx := strings.IndexFunc(s, operator)
	if opIdx == -1 {
		return Rule{leftRight[0], "", 0, ""}
	}
	rightNext := strings.Split(leftRight[1], ":")

	right, _ := strconv.Atoi(rightNext[0])

	next := rightNext[1]

	return Rule{leftRight[0], string(s[opIdx]), right, next}
}
