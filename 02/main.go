package main

import (
	"bufio"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
}

func getPairsFrom(csv string) []pair {
	stringPairs := strings.Split(csv, ",")
	retval := make([]pair, len(stringPairs))
	for idx, dashpair := range stringPairs {
		fmt.Println("dashpair: ", dashpair)
		stringPair := strings.Split(dashpair, "-")
		begin, _ := strconv.Atoi(stringPair[0])
		end, _ := strconv.Atoi(stringPair[1])
		myPair := pair{begin, end}
		retval[idx] = myPair
	}

	return retval
}

func isInvalidPart1(num int) bool {
	return num >= 10 && num <= 99 && num % 11 == 0 || 
	       num >= 1000 && num <= 9999 && num % 101 == 0 || 
	       num >= 100000 && num <= 999999 && num % 1001 == 0 || 
	       num >= 10000000 && num <= 99999999 && num % 10001 == 0 || 
	       num >= 1000000000 && num <= 9999999999 && num % 100001 == 0 
}

// input.txt has 10-digit max --> 100010001
func isInvalidPart2(num int) bool {
	// one digit repeating:
	return num >= 100 && num <= 999 && num % 111 == 0 || 
	       num >= 1000 && num <= 9999 && num % 1111 == 0 || 
	       num >= 10000 && num <= 99999 && num % 11111 == 0 || 
	       num >= 100000 && num <= 999999 && num % 111111 == 0 || 
	       num >= 1000000 && num <= 9999999 && num % 1111111 == 0 || 
	       num >= 10000000 && num <= 99999999 && num % 11111111 == 0 || 
	       num >= 100000000 && num <= 999999999 && num % 111111111 == 0 || 
	       num >= 1000000000 && num <= 9999999999 && num % 1111111111 == 0 || 
	       //                          1234567890
	       // two digits repeating (212121, etc)
	       num >= 100000 && num <= 999999 && num % 10101 == 0 ||
	       num >= 10000000 && num <= 99999999 && num % 1010101 == 0 ||
	       num >= 1000000000 && num <= 9999999999 && num % 101010101 == 0 ||
	       // three digits repeating (321321321)
	       num >= 100000 && num <= 999999 && num % 1001 == 0 ||
	       num >= 100000000 && num <= 999999999 && num % 1001001 == 0 ||
	       // four digits repeating (43214321)
	       num >= 10000000 && num <= 99999999 && num % 10001 == 0 ||
	       // five digits repeating (5432154321)
	       num >= 1000000000 && num <= 9999999999 && num % 100001 == 0
}

func main() {
	for _, i := range []int{1, 3, 11, 22} {
		fmt.Println(i, " isInvalid:", isInvalidPart1(i))
	}

	sum := 0
	sum2 := 0
	processLine := func(s string) {
		pairs := getPairsFrom(s)
		for _, pair := range pairs {
			fmt.Println(pair)
			i := pair.begin
			for i <= pair.end {
				if isInvalidPart1(i) {
					sum += i
					sum2 += i
					fmt.Println("invalid: ", i)
				} else if isInvalidPart2(i) {
					sum2 += i
					fmt.Println("invalid2: ", i)
				}
				i++
			}
		}
	}

	processLines("input.txt", processLine)
	fmt.Println("Part 1: ", sum) // 21139440284
	fmt.Println("Part 2: ", sum2) // 41263441243 is too high....??? but works on the test input :( 
}
