package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://www.scaler.com/topics/golang/golang-read-file-line-by-line/

func readFileAndReturnData(filePath string) string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file: %w", err)
		return ""
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	data := ""
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		data = data + line
	}

	return data
}

func isNumber(digit string) bool {
	regex := regexp.MustCompile(`^\d+$`)
	return regex.MatchString(digit)
}

func matchesPattern(str string) bool {
	regex := regexp.MustCompile(`^mul\(\d{1,3},\d{1,3}\)$`)
	return regex.MatchString(str)
}

func partTwo() {
	filePath := "data-part-2.txt"
	conutDont := 0
	countDo := 0
	data := readFileAndReturnData(filePath)
	shouldMultiply := true
	total := 0
	i := 0
	n := len(data)
	for i < n {
		if i <= len(data)-4 {
			do := string(data[i : i+4])
			if do == "do()" {
				countDo += 1
				shouldMultiply = true
				i += 1
				continue
			}
		}
		if i <= len(data)-7 {
			dont := string(data[i : i+7])
			if dont == "don't()" {
				conutDont += 1
				shouldMultiply = false
				i += 1
				continue
			}
		}
		if !shouldMultiply {
			i += 1
			continue
		}
		if i <= len(data)-3 {
			mul := string(data[i : i+3])
			if mul == "mul" {
				closingBracketIdx := strings.Index(string(data[i:]), ")") + i
				chunk := string(data[i : closingBracketIdx+1])
				matches := matchesPattern(chunk)

				if matches {
					chunkSplit := strings.Split(chunk, "mul(")
					secondSplit := strings.Split(chunkSplit[1], ",")
					firstInt, _ := strconv.Atoi(secondSplit[0])
					sliceTo := len(secondSplit[1]) - 1

					secondSplitInt, _ := strconv.Atoi((secondSplit[1][0:sliceTo]))

					mul := firstInt * secondSplitInt

					total += mul

				}
			}
		}
		i += 1
	}
	println(total, "expected: ", 94455185)
}

func main() {
	partTwo()
}

// mul(234,456)
