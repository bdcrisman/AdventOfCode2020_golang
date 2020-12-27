package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	optional = "cid"
	max      = 8
	min      = 7
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	count := validPassportCount(string(input), true)
	fmt.Printf("Part 1: Num valid passports: %d\n", count)

	count = validPassportCount(string(input), false)
	fmt.Printf("Part 2: Num valid passports: %d\n", count)
}

func validPassportCount(input string, isPartOne bool) int {
	if input == "" {
		return 0
	}

	lines := strings.Split(input, "\n")
	arrLines := make([]string, 0)
	count := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if isPassportValid(arrLines, isPartOne) {
				count++
			}
			arrLines = make([]string, 0)
		} else {
			arrLines = append(arrLines, strings.TrimSpace(line))
		}
	}

	if len(arrLines) > 0 && isPassportValid(arrLines, isPartOne) {
		count++
	}

	return count
}

func isPassportValid(lines []string, isPartOne bool) bool {
	m := parseMap(lines)
	if m == nil {
		return false
	}

	_, optionPresent := m[optional]
	result := len(m) == max || (len(m) == min && !optionPresent)
	if isPartOne || !result {
		return result
	}

	// part
	hcl, ok := m["hcl"]
	if !ok {
		return false
	}

	ecl, ok := m["ecl"]
	if !ok {
		return false
	}

	pid, ok := m["pid"]
	if !ok {
		return false
	}

	if _, ok = m["byr"]; !ok {
		return false
	}

	if _, ok = m["iyr"]; !ok {
		return false
	}

	if _, ok = m["eyr"]; !ok {
		return false
	}

	if _, ok = m["hgt"]; !ok {
		return false
	}

	byr, err := strconv.Atoi(m["byr"])
	if err != nil {
		return false
	}

	iyr, err := strconv.Atoi(m["iyr"])
	if err != nil {
		return false
	}

	eyr, err := strconv.Atoi(m["eyr"])
	if err != nil {
		return false
	}

	hgtCmStr := strings.Split(m["hgt"], "c")
	hgtCM := -1
	if len(hgtCmStr) > 1 {
		hgtCM, _ = strconv.Atoi(hgtCmStr[0])
	}

	hgtInStr := strings.Split(m["hgt"], "i")
	hgtIN := -1
	if len(hgtInStr) > 1 {
		hgtIN, _ = strconv.Atoi(hgtInStr[0])
	}

	if _, err := strconv.Atoi(m["pid"]); err != nil {
		return false
	}

	hasColors := make([]bool, 0)
	hasColors = append(hasColors, ecl == "amb")
	hasColors = append(hasColors, ecl == "blu")
	hasColors = append(hasColors, ecl == "brn")
	hasColors = append(hasColors, ecl == "gry")
	hasColors = append(hasColors, ecl == "grn")
	hasColors = append(hasColors, ecl == "hzl")
	hasColors = append(hasColors, ecl == "oth")
	colorCount := 0
	for _, res := range hasColors {
		if res {
			colorCount++
		}
	}

	return byr >= 1920 && byr <= 2002 &&
		iyr >= 2010 && iyr <= 2020 &&
		eyr >= 2020 && eyr <= 2030 &&
		((hgtCM >= 150 && hgtCM <= 193) ||
			(hgtIN >= 59 && hgtIN <= 76)) &&
		strings.HasPrefix(hcl, "#") && len(hcl) == 7 &&
		colorCount == 1 &&
		len(pid) == 9
}

func parseMap(lines []string) map[string]string {
	if len(lines) <= 0 {
		return nil
	}

	m := make(map[string]string)

	for _, line := range lines {
		kvps := strings.Split(line, " ")
		for _, k := range kvps {
			kvp := strings.Split(k, ":")
			if len(kvp) > 1 && kvp[0] != "" {
				m[kvp[0]] = kvp[1]
			}
		}
	}

	return m
}
