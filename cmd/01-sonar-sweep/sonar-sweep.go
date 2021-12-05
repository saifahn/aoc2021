package main

import (
	"aoc/utils"
	"log"
	"strconv"
)

func main() {
	lines, err := utils.ReadLines("data/01")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	increased := 0
	prevSum := 0
	var currentSum int

	// make the first sum
	for i := 0; i < 3; i++ {
		current, _ := strconv.Atoi(lines[i])
		prevSum += current
	}

	for i := 3; i < len(lines); i++ {
		smallest, _ := strconv.Atoi(lines[i-3])
		current, _ := strconv.Atoi(lines[i])
		// add the latest, take away the last
		currentSum = prevSum - smallest + current
		if currentSum > prevSum {
			increased++
		}
		prevSum = currentSum
	}
	println(increased)
}
