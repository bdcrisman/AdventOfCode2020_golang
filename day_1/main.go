package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	expenses, err := readfile("input.txt")
	if err != nil {
		log.Fatal("err reading input file")
	}

	e1, e2 := partOneProductOfTwoNumsSummingTotalSum(expenses, 2020)
	fmt.Printf("Part 1: entries that sum to %d: %d * %d = %d",
		2020, e1, e2, e1*e2)

	e1, e2, e3 := partTwoProductOfTwoNumsSummingTotalSum(expenses, 2020)
	fmt.Printf("\nPart 2: entries that sum to %d: %d * %d * %d = %d",
		2020, e1, e2, e3, e1*e2*e3)
}

func readfile(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]int, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			lines = append(lines, i)
		}
	}
	sort.Ints(lines)

	return lines, nil
}

func partOneProductOfTwoNumsSummingTotalSum(expenses []int, totalSum int) (int, int) {
	lastIndex := len(expenses) - 1
	iterations := 0
	entry1 := 0
	entry2 := 0
	found := false

	for !found && lastIndex > 0 {
		iterations++
		entry1 = expenses[lastIndex]
		for i := 0; i < len(expenses); i++ {
			sum := entry1 + expenses[i]
			if sum > totalSum {
				lastIndex--
				break
			}

			if sum == totalSum {
				found = true
				entry2 = expenses[i]
			}
		}
	}
	return entry1, entry2
}

func partTwoProductOfTwoNumsSummingTotalSum(expenses []int, totalSum int) (int, int, int) {
	lastIndex := len(expenses) - 1
	iterations := 0
	entry1 := 0
	entry2 := 0
	entry3 := 0
	found := false

	for !found && lastIndex > 0 {
		iterations++
		big := expenses[lastIndex]
		for i := 0; i < len(expenses); i++ {
			min := expenses[i]
			for j := i + 1; j < len(expenses); j++ {
				sum := min + big + expenses[j]
				if sum == totalSum {
					found = true
					entry1 = min
					entry2 = big
					entry3 = expenses[j]
				}
			}
		}
		lastIndex--
	}
	return entry1, entry2, entry3
}
