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

/*
*
 1. Import the file
 2. split the file at mul
 3. Iterate through the split list.
    a. if the starting element in the split char does not begin with brackets `(`
    - continue
    b. if it begins with an opening brackets `(`
    - check if the is a closing bracket in the current substr
    - if not:
    - continue
    else:
    i. create a stack,
    - create a temp str, create two bools to check if we are currently in the bracket and set it to false
    one for open and one for close
    - loop trough every char in the substr
    - if current char is not a number or an open or close bracket or a comma`,`
    -- break from the loop
    - if the current char is a number and open is false
    -- break from the loop
    break
    if the current char is a comma and open is false
    break
    if current char is `(` and open is false
    set open to true and continue
    if open is true and close is false and the current char is a number?
    append it to temp str
    if the current char is a comma and the open is true and close is false
    push temp to stack
    set temp to empty str
    continue
    if it is a closing char
    break
    multiply both int and add it total

________________________________________________________________________

regexToMatchPattern=	^\(\d{1,3},\d{1,3}\)$
*/
func partOne() {
	filePath := "data-part-1.txt"
	data := readFileAndReturnData(filePath)

	listOfStr := strings.Split(data, "mul")
	total := 0
	N := len(listOfStr)

	for i := 0; i < N; i += 1 {
		char := listOfStr[i]

		startingChar := char[:1]
		if startingChar != "(" {
			continue
		}
		if !strings.Contains(char, ")") {
			continue
		}
		stack := []string{}
		open := false
		temp := ""
		digits := "1234567890"
		for _, r := range char {
			currChar := string(r)
			if currChar != "(" && currChar != ")" && currChar != "," && !strings.Contains(digits, currChar) {
				break
			}
			if isNumber(currChar) && !open {
				break
			}
			if currChar == "," && !open {
				break
			}
			if currChar == "(" && open {
				break
			}
			if currChar == "(" {
				open = true
				continue
			}
			if isNumber(currChar) {
				temp = temp + currChar
				continue
			}
			if currChar == "," {
				stack = append(stack, temp)
				temp = ""
			}
			if currChar == ")" {
				stack = append(stack, temp)
				break
			}
		}
		if len(stack) == 2 {
			firstDigit, _ := strconv.Atoi(stack[0])
			secondDigit, _ := strconv.Atoi(stack[1])

			multi := firstDigit * secondDigit
			total += multi

		}
	}
	fmt.Println(total)

}

func partTwo() {}

func main() {
	partOne()
	partTwo()
}

// ^\(\d{1,3},\d{1,3}\)$
