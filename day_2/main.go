package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ParsedInput is each line of the puzzle input
type ParsedInput struct {
	Min   int
	Max   int
	Ch    string
	Pword string
}

var validCount int

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("err reading input file")
	}

	partOne := partOneCheckHowManyPasswordsAreValid(string(input))
	validCount = 0
	partTwo := partTwoCheckHowManyPasswordsAreValid(string(input))

	fmt.Printf("Part 1: num correct passwords: %d", partOne)
	fmt.Printf("\nPart 2: num correct passwords: %d", partTwo)
}

func partOneCheckHowManyPasswordsAreValid(input string) int {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var p ParsedInput
		if p.ParseLine(line) && p.IsValidPassword() {
			validCount++
		}
	}
	return validCount
}

func partTwoCheckHowManyPasswordsAreValid(input string) int {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var p ParsedInput
		if p.ParseLine(line) && p.IsReallyValidPassword() {
			validCount++
		}
	}
	return validCount
}

func (p *ParsedInput) ParseLine(line string) bool {

	if line == "" {
		return false
	}

	splits := strings.Split(line, " ")
	if len(splits) != 3 {
		return false
	}

	minmax := strings.Split(splits[0], "-")
	if len(minmax) != 2 {
		return false
	}

	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return false
	}

	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return false
	}

	p.Min = min
	p.Max = max
	p.Ch = string(splits[1][0])
	p.Pword = splits[2]

	return true
}

func (p ParsedInput) IsValidPassword() bool {
	if p.Pword == "" {
		return false
	}

	count := strings.Count(p.Pword, p.Ch)
	return count >= p.Min && count <= p.Max
}

func (p ParsedInput) IsReallyValidPassword() bool {
	if p.Pword == "" {
		return false
	}

	c1 := string(p.Pword[p.Min-1]) == p.Ch
	c2 := string(p.Pword[p.Max-1]) == p.Ch

	return (c1 || c2) && c1 != c2
}
