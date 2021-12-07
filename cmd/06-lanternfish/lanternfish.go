package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Lanternfish = int

func parseLanternfish(lines []string) []Lanternfish {
	if len(lines) != 1 {
		log.Fatal("expected just one line")
	}
	fish := strings.Split(lines[0], ",")
	var lanternfish []int
	for _, f := range fish {
		lf, _ := strconv.Atoi(f)
		lanternfish = append(lanternfish, lf)
	}
	return lanternfish
}

func main() {
	lines, err := utils.ReadLines("data/06-test")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	lanternfish := parseLanternfish(lines)
	fmt.Println(lanternfish)
}
