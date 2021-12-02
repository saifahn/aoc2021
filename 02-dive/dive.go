package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	lines, err := readLines("day2input")
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
