package main

import (
	"aoc/utils"
	"log"
	"strconv"
)

func getOxygen(lines []string, i int) string {
	var linesWithOne, linesWithZero, subsetLines []string
	for _, line := range lines {
		if line[i] == '0' {
			linesWithZero = append(linesWithZero, line)
		} else {
			linesWithOne = append(linesWithOne, line)
		}
	}

	if len(linesWithOne) >= len(linesWithZero) {
		subsetLines = linesWithOne
	} else {
		subsetLines = linesWithZero
	}

	if len(subsetLines) == 1 {
		return subsetLines[0]
	}

	return getOxygen(subsetLines, i+1)
}

func getCO2(lines []string, i int) string {
	var linesWithOne, linesWithZero, subsetLines []string
	for _, line := range lines {
		if line[i] == '0' {
			linesWithZero = append(linesWithZero, line)
		} else {
			linesWithOne = append(linesWithOne, line)
		}
	}

	if len(linesWithZero) <= len(linesWithOne) {
		subsetLines = linesWithZero
	} else {
		subsetLines = linesWithOne
	}

	if len(subsetLines) == 1 {
		return subsetLines[0]
	}

	return getCO2(subsetLines, i+1)
}

func main() {
	lines, err := utils.ReadLines("data/03")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	oxygen := getOxygen(lines, 0)
	co2 := getCO2(lines, 0)
	oxygenDecimal, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2, 2, 64)

	println(co2Decimal * oxygenDecimal)
}
