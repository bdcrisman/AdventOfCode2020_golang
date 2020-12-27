package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("err reading file")
	}

	s := string(input)

	part1 := countTreesEncountered(s, 3, 1)
	fmt.Printf("Part 1: Num trees encountered: %d", part1)

	r1d1 := countTreesEncountered(s, 1, 1)
	r5d1 := countTreesEncountered(s, 5, 1)
	r7d1 := countTreesEncountered(s, 7, 1)
	r1d2 := countTreesEncountered(s, 1, 2)
	product := r1d1 * part1 * r5d1 * r7d1 * r1d2
	fmt.Printf("\nPart 2: Product of num trees encountered: %d*%d*%d*%d%d = %d", r1d1, part1, r5d1, r7d1, r1d2, product)
}

func countTreesEncountered(input string, right, down int) int {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return 0
	}

	lastIndex := len(lines) - 1
	rowIndex := 0
	colIndex := 0
	count := 0

	for rowIndex+down <= lastIndex {
		rowIndex += down
		colIndex += right

		row := strings.TrimSpace(lines[rowIndex])
		if colIndex >= len(row) {
			colIndex -= len(row)
		}

		if string(lines[rowIndex][colIndex]) == "#" {
			count++
		}
	}

	return count
}
