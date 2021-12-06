package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func markBoards(number string, boards [][][]string) {
	for _, board := range boards {
		// go through
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

func checkForBingo(boards [][][]string) (bool, [][]string) {
	// go through boards
	var bingo bool
	for _, board := range boards {
		// go over rows
		for i, row := range board {
			// fmt.Printf("length of the row %v", len(row))
			if i == 0 {
				for j := range row {
					bingo = checkColumn(board, j)
					if bingo {
						return true, board
					}
				}
			}
			bingo = checkRow(row)
			if bingo {
				return true, board
			}
		}
	}
	return false, nil
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

	// for _, s := range sections {
	// 	fmt.Printf("%s\n\n", s)
	// }

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

	// for _, board := range boards {
	// 	// fmt.Printf("%v", board)
	// }

	// iterate through numbers
	// checkForBingo
	for _, number := range splitNumbers {
		// markBoards
		markBoards(number, boards)
		// checkForBingo
		bingo, board := checkForBingo(boards)
		if bingo {
			// usedNumbers := number[0:i]
			boardSum := calculateUnmarked(board)
			currentNumber, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("failed to convert char, %q", err)
			}
			fmt.Println(boardSum)
			fmt.Println(currentNumber)
			fmt.Println(boardSum * currentNumber)
			break
		}
	}
}
