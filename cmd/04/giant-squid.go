package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func markBoard(number string, board [][]string) {
out:
	for _, row := range board {
		for i, char := range row {
			if char == number {
				row[i] = "X"
				break out
			}
		}
	}
}

func checkBingoSingle(board [][]string) bool {
	for i, row := range board {
		// fmt.Printf("length of the row %v", len(row))
		if i == 0 {
			for j := range row {
				bingo := checkColumn(board, j)
				if bingo {
					return true
				}
			}
		}
		bingo := checkRow(row)
		if bingo {
			return true
		}
	}
	return false
}

func checkColumn(board [][]string, index int) bool {
	// fmt.Printf("%q\n\n", index)
	for _, row := range board {
		if row[index] != "X" {
			return false
		}
	}
	return true
}

func checkRow(row []string) bool {
	for _, char := range row {
		if char != "X" {
			return false
		}
	}
	return true
}

func intContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func calculateUnmarked(board [][]string) int {
	var total int
	for _, row := range board {
		for _, char := range row {
			if char != "X" {
				num, err := strconv.Atoi(char)
				if err != nil {
					log.Fatalf("failed to convert char, %q", err)
				}
				total += num
			}
		}
	}
	return total
}

func main() {
	file, err := os.Open("data/04")
	if err != nil {
		log.Fatalf("couldn't open file: %q", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	combined := strings.Join(lines, "\n")
	sections := strings.Split(combined, "\n\n")

	numbers := sections[0]
	boardsRaw := sections[1:]
	splitNumbers := strings.Split(numbers, ",")

	var boards [][][]string
	// each board is just a string first
	for _, b := range boardsRaw {
		// split the boards into rows
		boardRows := strings.Split(b, "\n")
		var board [][]string
		for _, br := range boardRows {
			// split the rows into a slice of strings
			splitRow := strings.Fields(br)
			board = append(board, splitRow)
		}
		boards = append(boards, board)
	}

	var completedBoardIndexes []int
	lastBoardState := make([][]string, len(boards[0]))
	var lastWonNumber int

	for _, number := range splitNumbers {
		for i, board := range boards {
			if intContains(completedBoardIndexes, i) {
				continue
			}
			markBoard(number, board)
			bingo := checkBingoSingle(board)
			if bingo {
				completedBoardIndexes = append(completedBoardIndexes, i)
				copy(lastBoardState, board)
				lastWonNumber, _ = strconv.Atoi(number)
			}
		}
	}
	boardSum := calculateUnmarked(lastBoardState)

	fmt.Printf("boardSum, %v, \n", lastBoardState)
	fmt.Printf("boardSum, %v, \n", boardSum)
	fmt.Printf("lastWonNumber, %v, \n", lastWonNumber)
	fmt.Printf("total, %v, \n", boardSum*lastWonNumber)
}
