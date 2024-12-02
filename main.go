package main

// https://www.scaler.com/topics/golang/golang-read-file-line-by-line/
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func partOne() {

	filePath := "data-part-1.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	var listOne []int
	var listTwo []int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		trimmedStr := strings.Join(strings.Fields(line), " ")
		splitStr := strings.Split(trimmedStr, " ")
		firstInt, err1 := strconv.Atoi(splitStr[0])
		twoInt, err2 := strconv.Atoi(splitStr[1])

		if err1 != nil || err2 != nil {
			fmt.Println(err1.Error())
			fmt.Println(err2.Error())
		}

		listOne = append(listOne, firstInt)
		listTwo = append(listTwo, twoInt)

		fileLines = append(fileLines, line)
	}
	slices.Sort(listOne)
	slices.Sort(listTwo)

	n := len(listOne)
	total := 0
	for i := 0; i < n; i += 1 {
		diff := listOne[i] - listTwo[i]
		abs := math.Abs(float64(diff))
		total = total + int(abs)
	}
	// 1222801
	fmt.Println(total)
}

func main() {
	partOne()
}
