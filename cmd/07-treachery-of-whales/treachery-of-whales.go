package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strings"
)

func parseCrabs(lines []string) []string {
	if len(lines) != 1 {
		log.Fatal("expected just one line")
	}
	return strings.Split(lines[0], ",")
}

func main() {
	lines, err := utils.ReadLines("data/07-test")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}
	crabs := parseCrabs(lines)
	fmt.Printf("%v", crabs)
}
