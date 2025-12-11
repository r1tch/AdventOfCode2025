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

// observations:
// every 2nd line is empty - we can use those
// splitters always have room for beams next to them (e.g .^.), we can use those to mark a beam
func main() {
	runTests()

	sum := 0
	sum2 := 0
	// field := make([]string, 0)
	prev := make([]rune, 0)
	// for each line and each position, we store how many paths lead here. 
	// the total number of paths are simply the sum of these, eg:
	//    |
	//   |^|
	//  |^|^|
	//  1 2 1   <-- num of paths
	//  ^ ^ ^
	// | | | |   
	// 1 3 3 1  <-- num of paths
	prevNumPaths := make([]int, 0)
	processLine := func(s string) {
		runes := []rune(s)
		if len(prev) == 0 {
			prev = runes
			prevNumPaths = make([]int, len(prev))
			return
		}
		
		newNumPaths := make([]int, len(prev))
		sum2 = 0
		for idx, _ := range runes {
			switch prev[idx] {
			case 'S':
				runes[idx] = '|'
				newNumPaths[idx] = 1
			case '|':
				if runes[idx] == '^' {
					runes[idx-1] = '|'
					runes[idx+1] = '|'
					newNumPaths[idx-1] += prevNumPaths[idx]
					newNumPaths[idx+1] += prevNumPaths[idx]
					sum++
					sum2 += prevNumPaths[idx] * 2
				} else if runes[idx] == '.' || runes[idx] == '|' { // in previous if we might have set the current char to |
					runes[idx] = '|'
					newNumPaths[idx] += prevNumPaths[idx]
					sum2 += prevNumPaths[idx]
				}
			}
		}
		prev = runes
		prevNumPaths = newNumPaths
		fmt.Println(string(runes))
		fmt.Println(prevNumPaths)
	}

	processLines("input.txt", processLine)
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
