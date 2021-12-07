package main

import (
	"aoc/utils"
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

func simulate(days int, lf []Lanternfish) (finalTotal int) {
	for i := 0; i < days; i++ {
		newFish := 0
		// decrease timer on each fish
		for j := range lf {
			lf[j] -= 1
			if lf[j] < 0 {
				lf[j] = 6
				newFish++
			}
		}
		// add the new fish
		for k := 0; k < newFish; k++ {
			lf = append(lf, 8)
		}
	}
	return len(lf)
}

func main() {
	lines, err := utils.ReadLines("data/06-test")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	lanternfish := parseLanternfish(lines)
	total := simulate(18, lanternfish)
	println(total)
}
