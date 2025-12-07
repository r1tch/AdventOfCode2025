package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pair struct {
	begin int
	end int
}

type coord struct {
	row int
	col int
}

func processLines(filename string, f func(string)) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
}

func runTests() {
}

func printField(field []string) {
	fmt.Println("---")
	for _, s := range field {
		fmt.Println(s)
	}
}

func set(field []string, row int, col int, newRune rune) {
	runes := []rune(field[row])
	runes[col] = newRune
	field[row] = string(runes)
}

func get(field []string, row int, col int) rune {
	if row < 0 || col < 0 || row >= len(field) || col >= len(field[row]) {
		return '\000'
	}

	return rune(field[row][col])
}

func numNeighborsOccupied(field []string, row int, col int) int {
	num := 0
	if get(field, row-1, col-1) == '@' {
		num++
	}
	if get(field, row-1, col) == '@' {
		num++
	}
	if get(field, row-1, col+1) == '@' {
		num++
	}
	if get(field, row, col-1) == '@' {
		num++
	}
	if get(field, row, col+1) == '@' {
		num++
	}
	if get(field, row+1, col-1) == '@' {
		num++
	}
	if get(field, row+1, col) == '@' {
		num++
	}
	if get(field, row+1, col+1) == '@' {
		num++
	}
	return num
}

func isAccessible(field []string, row int, col int) bool {
	return numNeighborsOccupied(field, row, col) < 4
}


func numAccessible(field []string, rwField []string) int {
	num := 0
	width := len(field[0])
	accessible := make([]coord, 0)
	for row := 0; row < len(field); row++ {
		for col := 0; col < width; col++ {
			if get(field, row, col) == '@' && isAccessible(field, row, col) {
				num++
				set(rwField, row, col, 'x')
				accessible = append(accessible, coord{row, col})
			}
		}
	}

	for _, coord := range accessible {
		set(field, coord.row, coord.col, '.')
	}
	return num
}

func main() {
	field := make([]string, 0)
	runTests()

	sum := 0
	sum2 := 0
	processLine := func(s string) {
		field = append(field, s)
	}

	processLines("input.txt", processLine)
	fieldCopy := make([]string, len(field))
	copy(fieldCopy, field)
	sum = numAccessible(field, fieldCopy)
	printField(fieldCopy)

	currAccessible := sum
	sum2 = currAccessible

	for currAccessible != 0 {
		currAccessible = numAccessible(field, fieldCopy)
		fmt.Println("removed ", currAccessible, "sum2", sum2)
		printField(fieldCopy)
		sum2 += currAccessible
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
