package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	begin int
	end int
}

func processLines(filename string, f func(string)) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
}

func getDigit(lines [][]rune, col int, row int) int {
	runes := lines[row]
	if len(runes) > col && runes[col] != ' ' {
		return int(runes[col] - '0')
	}

	return -1
}

func getNumInCol(lines [][]rune, col int) int {
	num := 0
	for row := 0; row < len(lines); row++ {
		if digit := getDigit(lines, col, row); digit != -1 {
			num = num * 10 + digit
		}
	}

	return num
}

func runTests() {
}

func main() {
	runTests()

	sum := 0
	sum2 := 0
	rows := make([][]int, 0)
	operations := make([]string, 0)
	operationsRunes := make([]rune, 0)
	lines := make([][]rune, 0)
	maxLen := 0
	processLine := func(s string) {
		row := make([]int, 0)
		// rowStr := strings.Split(s, " ")
		rowStr := strings.FieldsFunc(s, func(r rune) bool {
			return r == ' '
		})
		if rowStr[0] == "+" || rowStr[0] == "*" {
			operations = rowStr
		} else {
			for _, s := range rowStr {
				num, _ := strconv.Atoi(s)
				row = append(row, num)
			}
			fmt.Println(row)
			rows = append(rows, row)
		}
		if strings.HasPrefix(s, "*") || strings.HasPrefix(s, "+") {
			operationsRunes = []rune(s)
		} else {
			lines = append(lines, []rune(s))
			if maxLen < len(s) {
				maxLen = len(s)
			}
		}
	}

	processLines("input.txt", processLine)
	fmt.Println(operations)

	// part 1
	for col, operation := range operations {
		columnResult := 0
		if operation == "*" {
			columnResult = 1
		}
		for row := 0; row < len(rows); row++ {
			if operation == "+" {
				columnResult += rows[row][col]
			} else {
				columnResult *= rows[row][col]
			}
		}
		//fmt.Println(operation, "col", col, "result", columnResult)
		sum += columnResult
	}

	// part 2
	// determine max length
	//   -- func getNumInCol(lines, col) int --> lookup
	//   -- func getResultsFor(lines, fromCol, toCol, operation rune) int --> does operations
	// go over operationsRunes - using column value to get numbers from-to

	operation := ' '
	tmpSum := 0
	for col := 0; col < maxLen; col++ {
		if col < len(operationsRunes) && operationsRunes[col] != ' ' {
			operation = operationsRunes[col]
			sum2 += tmpSum
			fmt.Println("tmpSum", tmpSum)
			switch operation {
				case '*': tmpSum = 1
				case '+': tmpSum = 0
			}
		}
		num := getNumInCol(lines, col)
		if num == 0 {
			continue
		}
		fmt.Println(col, num)
		switch operation {
			case '*': tmpSum *= num
			case '+': tmpSum += num
		}
	}
	sum2 += tmpSum
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
