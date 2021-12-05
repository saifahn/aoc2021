package main

import (
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("data/02")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	aim, horizontal, depth := 0, 0, 0

	for _, line := range lines {
		// split by " "
		split := strings.Split(line, " ")
		// first is the command
		command := split[0]
		// convert the second into the number
		number, _ := strconv.Atoi(split[1])
		// switch on the command
		switch command {
		case "forward":
			horizontal += number
			depth += number * aim
		case "down":
			aim += number
		case "up":
			aim -= number
		}
	}

	println(horizontal * depth)
}
