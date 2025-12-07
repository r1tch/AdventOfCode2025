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

func smax(s string) (int, int) {
	maxVal := 0
	maxInd := -1
	for idx, c := range s {
		num := int(c - '0')
		if num > maxVal {
			maxVal = num
			maxInd = idx
		}
	}
	return maxInd, maxVal
}

// the logic for joltLegth of 12: 
// - we must use up 12 digits (or else it can't be a maximum number anyways)
// - first digit must come from the first length-11 digits, its index will be maxInd
// - next digit must skip that, so we must start at maxInd+1
//   example input: 97 865 432 198 765 
//   --> first digit must come from can only use 978 (we need remaining 11 for the rest) -> we'll use 9
//   --> second digit can come from 786 --> we'll use 8
//   --> third digit comes from 65 ... etc
func getJolt(bank string, joltLength int) int {
	fmt.Println("getJolt", bank, joltLength)

	if len(bank) < joltLength {
		return 0
	}

	joltValue := 0
	maxInd := -1

	for digit := range(joltLength) {
		joltValue *= 10
		usableDigits := bank[maxInd+1:len(bank)-(joltLength - digit - 1)]
		newMaxInd, firstMax := smax(usableDigits)
		maxInd += newMaxInd + 1
		joltValue += firstMax
		fmt.Println("usableDigits", usableDigits, "maxInd:", maxInd, "max:", firstMax, "jolt:", joltValue)
	}
	return joltValue

}

func runTests() {
	if idx, val := smax("87697"); idx != 3 || val != 9 {
		panic(fmt.Sprintf("err idx:%d val:%d", idx, val))
	}

	if jolt := getJolt("87", 3); jolt != 0 {
		panic(jolt)
	}

	if jolt := getJolt("87", 2); jolt != 87 {
		panic(jolt)
	}

	if jolt := getJolt("873329", 2); jolt != 89 {
		panic(jolt)
	}

	if jolt := getJolt("3311", 2); jolt != 33 {
		panic(jolt)
	}

	if jolt := getJolt("111111", 2); jolt != 11 {
		panic(jolt)
	}
}

func main() {
	runTests()

	sum := 0
	sum2 := 0
	processLine := func(s string) {
		sum += getJolt(s, 2)
		sum2 += getJolt(s, 12)
	}

	processLines("input.txt", processLine)
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2) // 170849145841479 - too high, more tests needed
}
