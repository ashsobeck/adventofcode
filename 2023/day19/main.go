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

func main() {
	file, _ := os.ReadFile("./day19.txt")
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
}

func (p Part) RatingSum() int {
	return p.X + p.M + p.A + p.S
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
