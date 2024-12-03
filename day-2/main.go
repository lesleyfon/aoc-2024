package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertStrToInt(str1 string, str2 string) (int, int, error) {
	firstChar, err1 := strconv.Atoi(str1)
	secondChar, err2 := strconv.Atoi(str2)
	if err1 != nil {
		return 0, 0, err1
	}

	if err2 != nil {
		return 0, 0, err1
	}
	return firstChar, secondChar, nil
}

func checkLevelSafety(level []string) bool {
	firstChar, secondCHar, err := convertStrToInt(level[0], level[1])

	if err != nil {
		return false
	}
	unsafeCounter := 0

	if firstChar == secondCHar {
		unsafeCounter += 1
	}

	N := len(level)

	numbers := make([]int, N)

	for i, str := range level {
		num, err := strconv.Atoi(str)
		if err != nil {
			return false
		}
		numbers[i] = num
	}

	for i := 0; i < N-1; i += 1 {
		curr, next, err := convertStrToInt(level[i], level[i+1])

		if err != nil {
			return false
		}
		if curr < next {
			unsafeCounter += 1
			continue
		}

		diff := curr - next
		abs := math.Abs(float64(diff))

		if abs < 1 || abs > 3 {
			unsafeCounter += 1
		}

	}
	if unsafeCounter > 1 {
		return false
	}

	return true
}

func main() {
	fileName := "data-part-2.txt"
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		trimmedLine := strings.Join(strings.Fields(line), " ")
		splitLine := strings.Split(trimmedLine, " ")

		isLevelSafe := checkLevelSafety(splitLine)

		if isLevelSafe {
			total += 1
		}
	}

	fmt.Println(total)

}
