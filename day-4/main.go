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

	directions := []struct {
		name    string
		vectors [][2]int
		check   func(row int, col int) bool
	}{
		{"NORTH",      [][2]int{{-1, 0}, {-2, 0}, {-3, 0}},         func(r, c int) bool { return r >= 3 }},
		{"NORTH_EAST", [][2]int{{-1, 1}, {-2, 2}, {-3, 3}},         func(r, c int) bool { return r >= 3 && c < N-3 }},
		{"EAST",       [][2]int{{0, 1}, {0, 2}, {0, 3}},            func(r, c int) bool { return c < N-3 }},
		{"SOUTH_EAST", [][2]int{{1, 1}, {2, 2}, {3, 3}},            func(r, c int) bool { return r < M-3 && c < N-3 }},
		{"SOUTH",      [][2]int{{1, 0}, {2, 0}, {3, 0}},            func(r, c int) bool { return r < M-3 }},
		{"SOUTH_WEST", [][2]int{{1, -1}, {2, -2}, {3, -3}},         func(r, c int) bool { return r < M-3 && c >= 3 }},
		{"WEST",       [][2]int{{0, -1}, {0, -2}, {0, -3}},         func(r, c int) bool { return c >= 3 }},
		{"NORTH_WEST", [][2]int{{-1, -1}, {-2, -2}, {-3, -3}},      func(r, c int) bool { return r >= 3 && c >= 3 }},
	}

	for _, dir := range directions {
		if dir.check(row, col) {
			combo := ""
			for _, v := range dir.vectors {
				combo += grid[row+v[0]][col+v[1]]
			}
			if combo == "MAS" {
				count++
			}
		}
	}
	return count
}
func partOne() {

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
