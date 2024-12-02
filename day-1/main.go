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

func readFileAndReturnData(filePath string) (
	[]int,
	[]int,
) {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println("error opening file: %w", err)
	}

	defer readFile.Close()

	var listOne, listTwo []int

	if err != nil {
		fmt.Println("Error reading file:", err)
		return listOne, listTwo
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		trimmedStr := strings.Join(strings.Fields(line), " ")
		splitStr := strings.Split(trimmedStr, " ")

		if len(splitStr) < 2 {
			fmt.Println("invalid line format: insufficient fields")
			break
		}
		firstInt, err1 := strconv.Atoi(splitStr[0])
		twoInt, err2 := strconv.Atoi(splitStr[1])

		if err1 != nil || err2 != nil {
			fmt.Println(err1.Error())
			fmt.Println(err2.Error())
		}

		listOne = append(listOne, firstInt)
		listTwo = append(listTwo, twoInt)

	}
	return listOne, listTwo
}

func partOne() {

	filePath := "data-part-1.txt"
	listOne, listTwo := readFileAndReturnData(filePath)

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

func partTwo() {
	filePathStr := "data-part-2.txt"

	listOne, listTwo := readFileAndReturnData(filePathStr)

	var listTwoElemCount = map[int]int{}
	var n = len(listOne)

	for i := 0; i < n; i += 1 {
		_, exist := listTwoElemCount[listTwo[i]]
		if exist {
			listTwoElemCount[listTwo[i]] += 1
		} else {
			listTwoElemCount[listTwo[i]] = 1
		}
	}
	total := 0
	for i := 0; i < n; i += 1 {
		firstInt := listOne[i]

		count, exist := listTwoElemCount[firstInt]
		if exist {
			total += (count * firstInt)
		}
	}
	// 22545250
	println(total)

}

func main() {
	partOne()
	partTwo()
}
