Day 2: Red-Nosed Reports 
---
Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

```
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
```
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

- The levels are either all increasing or all decreasing.
- Any two adjacent levels differ by at least one and at most three.

In the example above, the reports can be found safe or unsafe by checking those rules:

- `7 6 4 2 1`: Safe because the levels are all decreasing by 1 or 2.
- `1 2 7 8 9`: Unsafe because 2 7 is an increase of 5.
- `9 7 6 2 1`: Unsafe because 6 2 is a decrease of 4.
- `1 3 2 4 5`: Unsafe because 1 3 is increasing but 3 2 is decreasing.
- `8 6 4 4 1`: Unsafe because 4 4 is neither an increase or a decrease.
- `1 3 6 7 9`: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

#### Analyze the unusual data from the engineers. How many reports are safe?

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.scaler.com/topics/golang/golang-read-file-line-by-line/

/*
 Reports == rows
 levels == cols
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

func convertCharToInt(charOne string, charTwo string) (int, int, error) {

	firstChar, err1 := strconv.Atoi(charOne)
	secondCHar, err2 := strconv.Atoi(charTwo)

	if err1 != nil || err2 != nil {
		fmt.Println(err1.Error())
		fmt.Println(err2.Error())

		return 0, 0, fmt.Errorf("error converting char to int")
	}

	return firstChar, secondCHar, nil
}

func isLevelSafe(level []string) bool {
	firstChar, secondCHar, err := convertCharToInt(level[0], level[1])

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if firstChar == secondCHar {
		return false
	}

	var ASC bool = firstChar < secondCHar
	N := len(level)

	for i := 1; i < N; i += 1 {
		if ASC {
			current, prev, err := convertCharToInt(level[i], level[i-1])
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
			if current < prev {
				return false
			}
			diff := current - prev

			if diff < 1 || diff > 3 {
				return false
			}

		} else {

			current, prev, err := convertCharToInt(level[i], level[i-1])

			if err != nil {
				fmt.Println(err.Error())
				return false
			}
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

```