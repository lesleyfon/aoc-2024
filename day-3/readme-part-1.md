Day 3: Mull It Over
 ---

"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?

```go

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

```