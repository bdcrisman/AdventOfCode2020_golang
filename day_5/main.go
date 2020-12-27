package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var (
	maxRows = 127
	maxCols = 7
	magic   = 8
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	lines := strings.Split(string(input), "\n")
	seatIDs := make([]int, 0)
	highestSeatID := 0

	for _, line := range lines {
		rowID := getSeatID(string(line[0:7]), 0, 0, maxRows)
		colID := getSeatID(string(line[7:]), 0, 0, maxCols)
		seatID := rowID*magic + colID

		seatIDs = append(seatIDs, seatID)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Printf("Part 1: Highest seat ID: %d", highestSeatID)

	mySeatID := getMySeatID(seatIDs)
	fmt.Printf("\nPart 2: My seat ID: %d", mySeatID)
}

func getSeatID(line string, i, lower, upper int) int {
	if i >= len(line) || lower >= upper || lower < 0 || upper < 0 {
		return lower
	}

	mid := (upper-lower)/2 + lower
	ch := string(line[i])
	if ch == "F" || ch == "L" {
		return getSeatID(line, i+1, lower, mid)
	}

	return getSeatID(line, i+1, mid+1, upper)
}

func getMySeatID(ids []int) int {
	if len(ids) <= 0 {
		return 0
	}

	sort.Ints(ids)
	seatID := 0
	max := len(ids) - 1

	for i := 0; i < max; i++ {
		if ids[i+1]-ids[i] == 2 {
			seatID = ids[i] + 1
			break
		}
	}

	return seatID
}
