package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parseCrabs(lines []string) []int {
	if len(lines) != 1 {
		log.Fatal("expected just one line")
	}
	crabs := strings.Split(lines[0], ",")
	var crabInts []int
	for _, c := range crabs {
		crab, _ := strconv.Atoi(c)
		crabInts = append(crabInts, crab)
	}
	return crabInts
}

func getRange(crabs []int) (min, max int) {
	// iterate through crabs
	min, max = crabs[0], crabs[0]
	for _, crab := range crabs {
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}
	}
	return min, max
}

func calculateLeastFuelNeeded(crabs []int, min, max int) int {
	leastFuel := 0
	for x := min; x <= max; x++ {
		distanceMap := map[int]int{}
		for _, crab := range crabs {
			distance := crab - x
			// get the absolute value
			if distance < 0 {
				distance = -distance
			}
			distanceMap[distance]++
		}
		total := 0
		// go through each entry in the distance map, get the total fuel for each distance
		for d, n := range distanceMap {
			// sum d
			fuel := 0
			for y := 1; y <= d; y++ {
				fuel += y
			}
			fuelForDistance := fuel * n
			total += fuelForDistance
		}
		if total < leastFuel {
			leastFuel = total
		}
		if leastFuel == 0 {
			leastFuel = total
		}
	}
	return leastFuel
}

func main() {
	lines, err := utils.ReadLines("data/07-real")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}
	crabs := parseCrabs(lines)
	min, max := getRange(crabs)
	leastFuel := calculateLeastFuelNeeded(crabs, min, max)
	// fmt.Printf("%v, %v", crabs)
	fmt.Printf("%v\n", leastFuel)
}
