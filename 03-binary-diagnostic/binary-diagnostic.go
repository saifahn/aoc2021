package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("day3input")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	// first, second, third, fourth, fifth := 0, 0, 0, 0, 0
	numLines := len(lines)
	lineLength := len(lines[0])
	digitSums := make([]int, lineLength)

	for _, line := range lines {
		for i, char := range line {
			// convert rune
			digitSums[i] += int(char - '0')
		}
	}

	gammaString, epsilonString := "", ""
	for _, sum := range digitSums {
		if sum > numLines/2 {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}
	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)
	println(gamma * epsilon)
}
