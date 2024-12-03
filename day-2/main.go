package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.scaler.com/topics/golang/golang-read-file-line-by-line/
// Reports == rows
// levels == cols

/*
- Read the file and for each line,
	i. Call a function that checks if the level is safe i.e. isLevelSafe
		1. if safe, return true
		2. else: return false
	isLevelSafe fn
		- takes a row as an array
		- create a bool to check is it is ascending or descending.
		- if list at 0 is equal to list at one
			- return false
		- if list at 0 is greater than list at one
			- set descending to true
		- else
			- set ascending to true
		- create a total variable an initialize it with one
		- iterate over each item
			- if descending is true
				check that the current item is less than the prev item
				three
		return true if we were able to iterate over every Item in the list.
*/

func convertCharToInt(charOne string, charTwo string) (int, int) {

	firstChar, err1 := strconv.Atoi(charOne)
	secondCHar, err2 := strconv.Atoi(charTwo)

	if err1 != nil || err2 != nil {
		fmt.Println(err1.Error())
		fmt.Println(err2.Error())
		// TODO: Update this to return error
		return 0, 0
	}

	return firstChar, secondCHar
}

func isLevelSafe(level []string) bool {
	firstChar, secondCHar := convertCharToInt(level[0], level[1])

	// TODO: uncomment this back
	// if err1 != nil || err2 != nil {
	// 	fmt.Println(err1.Error())
	// 	fmt.Println(err2.Error())
	// 	return false
	// }
	if firstChar == secondCHar {
		return false
	}

	var ASC bool = firstChar < secondCHar
	N := len(level)

	for i := 1; i < N; i += 1 {
		if ASC {
			current, prev := convertCharToInt(level[i], level[i-1])
			if current < prev {
				return false
			}
			diff := current - prev

			if diff < 1 || diff > 3 {
				return false
			}

		} else {
			current, prev := convertCharToInt(level[i], level[i-1])
			if current > prev {
				return false
			}
			diff := prev - current

			if diff < 1 || diff > 3 {
				return false
			}
		}
	}

	return true
}
func partOne() {
	fileName := "data-part-1.txt"
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Error while reading file")
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	total := 0
	for fileScanner.Scan() {

		line := fileScanner.Text()
		trimmedRow := strings.Join(strings.Fields(line), " ")
		lineReport := strings.Split(trimmedRow, " ")

		levelIsSafe := isLevelSafe(lineReport)
		if levelIsSafe {
			total += 1
		}
	}
	// ANS = 356
	fmt.Println(total)
}

func main() {
	partOne()
}
