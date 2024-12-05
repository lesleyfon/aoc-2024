package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFileAndReturnData(filePath string) [][]string {
	readFile, err := os.Open(filePath)
	grid := [][]string{}
	if err != nil {

		return grid
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		row := strings.Split(line, "")

		grid = append(grid, row)
	}
	return grid
}

func getOccurrence(grid [][]string, row int, col int) int {
	M := len(grid)
	N := len(grid[0])
	count := 0

	// NORTH := [][]int{{-1, 0}, {-2, 0}, {-3, 0}}
	if row >= 3 {
		M := grid[row-1][col]
		A := grid[row-2][col]
		S := grid[row-3][col]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}

	// NORTH_EAST := [][]int{{-1, 1}, {-2, 2}, {-3, 3}}
	if row >= 3 && col < N-3 {
		M := grid[row-1][col+1]
		A := grid[row-2][col+2]
		S := grid[row-3][col+3]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	// EAST := [][]int{{0, 1}, {0, 2}, {0, 3}}
	if col < N-3 {
		M := grid[row][col+1]
		A := grid[row][col+2]
		S := grid[row][col+3]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	// SOUTH_EAST := [][]int{{1, 1}, {2, 2}, {3, 3}}
	if row < M-3 && col < N-3 {
		M := grid[row+1][col+1]
		A := grid[row+2][col+2]
		S := grid[row+3][col+3]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	// SOUTH := [][]int{{1, 0}, {2, 0}, {3, 0}}
	if row < M-3 {
		M := grid[row+1][col]
		A := grid[row+2][col]
		S := grid[row+3][col]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	// SOUTH_WEST := [][]int{{1, -1}, {2, -2}, {3, -3}}
	if row < M-3 && col >= 3 {
		M := grid[row+1][col-1]
		A := grid[row+2][col-2]
		S := grid[row+3][col-3]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	// WEST := [][]int{{0, -1}, {0, -2}, {0, -3}}
	if col >= 3 {
		M := grid[row][col-1]
		A := grid[row][col-2]
		S := grid[row][col-3]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	// NORTH_WEST := [][]int{{-1, -1}, {-2, -2}, {-3, -3}}
	if col >= 3 && row >= 3 {
		M := grid[row-1][col-1]
		A := grid[row-2][col-2]
		S := grid[row-3][col-3]

		COMBO := M + A + S
		if COMBO == "MAS" {
			count += 1
		}
	}
	return count
}
func partOne() {

	// NORTH_WEST := [][]int{{-1, -1}, {-2, -2}, {-3, -3}}

	filePath := "data-part-1.txt"
	grid := readFileAndReturnData(filePath)

	M := len(grid)
	N := len(grid[0])
	result := 0
	for row := 0; row < M; row += 1 {
		for col := 0; col < N; col += 1 {
			currElem := grid[row][col]
			if currElem == "X" {
				result += getOccurrence(grid, row, col)
			}
		}
	}
	fmt.Println(result)
}

func partTwo() {}

func main() {
	partOne()
	partTwo()
}
