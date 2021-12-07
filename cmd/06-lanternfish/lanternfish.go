package main

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

type Day = int

type Lanternfish = map[Day]int

func parseLanternfish(lines []string) Lanternfish {
	if len(lines) != 1 {
		log.Fatal("expected just one line")
	}
	fish := strings.Split(lines[0], ",")
	lanternfish := Lanternfish{}
	for _, f := range fish {
		lf, _ := strconv.Atoi(f)
		lanternfish[lf] += 1
	}
	return lanternfish
}

func simulate(days int, lf Lanternfish) Lanternfish {
	for i := 0; i < days; i++ {
		// take the fish at 0 days, set reproduce
		reproducingFish := lf[0]
		// go through each key of the lanternfish, reduce them
		for j := 1; j < 9; j++ {
			lf[j-1] = lf[j]
		}
		// add the reproducingFish
		lf[6] += reproducingFish
		// replace the 8th because we didn't do it above
		lf[8] = reproducingFish
	}
	return lf
}

func main() {
	lines, err := utils.ReadLines("data/06-real")
	if err != nil {
		log.Fatalf("could not read lines: %s", err)
	}

	lanternfish := parseLanternfish(lines)
	end := simulate(256, lanternfish)
	total := 0
	for k := 0; k < 9; k++ {
		total += end[k]
	}
	// fmt.Printf("%v", end)
	// fmt.Printf("%v", end)
	println(total)
}
