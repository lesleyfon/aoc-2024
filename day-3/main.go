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
var (
	numberRegex     = regexp.MustCompile(`^\d+$`)
	mulPatternRegex = regexp.MustCompile(`^mul\(\d{1,3},\d{1,3}\)$`)
)

func readFileAndReturnData(filePath string) string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error opening file: %v", err.Error())
		return ""
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	builder := strings.Builder{}
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		builder.WriteString(line)
	}

	return builder.String()
}

func isNumber(digit string) bool {
	return numberRegex.MatchString(digit)
}

func matchesPattern(str string) bool {
	return mulPatternRegex.MatchString(str)
}

func partTwo() {
	// Variable
	filePath := "data-part-2.txt"
	data := readFileAndReturnData(filePath)
	// Determines if we should multiply the next equation
	shouldMultiply := true
	total := 0
	i := 0
	n := len(data)
	const (
		doLength   = 4 // length of "do()"
		dontLength = 7 // length of "don't()"
		mulLength  = 3 // length of "mul"
	)

	for i < n {
		strStartsWithD := string(data[i]) == "d"
		strStartsWithM := string(data[i]) == "m"

		// Check if we are out of bounds, id not
		if i <= n-doLength && strStartsWithD {
			do := string(data[i : i+doLength])
			if do == "do()" {
				shouldMultiply = true
				i += 1
				continue
			}
		}
		if i <= n-dontLength && strStartsWithD {
			dont := string(data[i : i+dontLength])
			if dont == "don't()" {

				shouldMultiply = false
				i += 1
				continue
			}
		}
		if !shouldMultiply {
			i += 1
			continue
		}
		if i <= n-mulLength && strStartsWithM {

			mul := string(data[i : i+mulLength])

			if mul == "mul" {
				// Find the closing bracket index
				closingBracketIdx := strings.Index(string(data[i:]), ")") + i
				// Get the chunk
				chunk := string(data[i : closingBracketIdx+1])
				// Check if the chunk matches the mul pattern
				matches := matchesPattern(chunk)

				if matches {

					chunkSplit := strings.Split(chunk, "mul(")
					secondSplit := strings.Split(chunkSplit[1], ",")
					firstInt, err := strconv.Atoi(secondSplit[0])

					if err != nil {
						fmt.Printf("error occurred while converting first split to int: %v", err)
						i += 1
						continue
					}

					sliceTo := len(secondSplit[1]) - 1

					secondSplitInt, err := strconv.Atoi((secondSplit[1][0:sliceTo]))
					if err != nil {
						fmt.Printf("error occurred while converting second split to int: %v", err)
						i += 1
						continue
					}

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
