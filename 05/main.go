package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pair struct {
	begin int
	end int
}

type coord struct {
	row int
	col int
}

type RangeSet struct {
	sortedRanges []pair
}

func rangeSort(a, b pair) int {
		switch {
		case a.begin < b.begin: return -1
		case a.begin > b.begin: return 1
		}
		return 0
}

func (rangeSet *RangeSet) add(pairs []pair) {
	for _, thePair := range pairs {
		// fmt.Println("Adding", thePair, "to", rangeSet.sortedRanges)
		if len(rangeSet.sortedRanges) == 0 {
			rangeSet.sortedRanges = []pair{thePair}
			continue
		}

		// We are merging the ranges. We store the "upcoming" addition in upcoming, 
		// which is not added immediately - but may be altered, depending on existing
		// ranges in the set.
		newRanges := make([]pair, 0)
		upcoming := thePair

		for _, existing := range rangeSet.sortedRanges {
			// existing     S...E 
			// upcoming S.E             --> add upcoming, upcoming := existing
			// upcoming           S.E   --> add existing, keep upcoming
			// upcoming   S...E         --> add nothing, upcoming := upcoming.begin, existing.end
			// upcoming        S...E    --> add nothing, upcoming := existing.begin, upcoming.end
			// upcoming      S.E        --> upcoming := existing
			// upcoming S..........E    --> do nothing, keep upcoming
			if existing.begin > upcoming.end {
				newRanges = append(newRanges, upcoming)
				upcoming = existing
			} else if existing.end < upcoming.begin {
				newRanges = append(newRanges, existing)
			} else if existing.begin >= upcoming.begin && existing.end >= upcoming.end {
				upcoming = pair{upcoming.begin, existing.end}
			} else if existing.begin <= upcoming.begin && existing.end <= upcoming.end {
				upcoming = pair{existing.begin, upcoming.end}
			} else if existing.begin <= upcoming.begin && existing.end >= upcoming.end {
				upcoming = existing
			// } else if existing.begin >= upcoming.begin && existing.end <= upcoming.end {
			// DO NOTHING
			}
			// fmt.Println(newRanges)
		}
		newRanges = append(newRanges, upcoming)
		rangeSet.sortedRanges = newRanges
	}
}

func (rangeSet *RangeSet) hasValue(value int) bool {

	// naive search solution - fast enough for us (can also use binary search...)
	for pos := 0; pos < len(rangeSet.sortedRanges) && rangeSet.sortedRanges[pos].begin <= value; pos++ {
		if rangeSet.sortedRanges[pos].begin <= value && rangeSet.sortedRanges[pos].end >= value {
			return true
		}
	}
	return false

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

func runTests() {
	var rangeSet RangeSet

	rangeSet.add([]pair{{1, 2}, {5,7}, {29, 45}, {40, 50}, {47, 47}, {60, 65}, {60, 60}, {62, 70}})

	// fmt.Println(rangeSet)

	if rangeSet.hasValue(0) != false {
		panic("")
	}

	if rangeSet.hasValue(1) != true {
		panic("")
	}
	if rangeSet.hasValue(2) != true {
		panic("")
	}
	if rangeSet.hasValue(3) != false {
		panic("")
	}
	if rangeSet.hasValue(5) != true {
		panic("")
	}
	if rangeSet.hasValue(6) != true {
		panic("")
	}
	if rangeSet.hasValue(7) != true {
		panic("")
	}
	if rangeSet.hasValue(8) != false {
		panic("")
	}
	if rangeSet.hasValue(47) != true {
		panic("")
	}
	if rangeSet.hasValue(50) != true {
		panic("")
	}
	if rangeSet.hasValue(60) != true {
		panic("")
	}
	if rangeSet.hasValue(61) != true {
		panic("")
	}
}

func main() {
	runTests()

	sum := 0
	sum2 := 0
	var rangeSet RangeSet
	pairs := make([]pair, 0)
	numbersToCheck := make([]int, 0)
	processLine := func(s string) {
		if len(s) == 0 {
			return
		}
		stringPair := strings.Split(s, "-")
		if len(stringPair) == 2 {
			begin, _ := strconv.Atoi(stringPair[0])
			end, _ := strconv.Atoi(stringPair[1])

			pairs = append(pairs, pair{begin, end})
		} else if len(stringPair) == 1 {
			num, _ := strconv.Atoi(stringPair[0])
			numbersToCheck = append(numbersToCheck, num)
		}
	}

	processLines("input.txt", processLine)

	rangeSet.add(pairs)
	// fmt.Println("pairs:", pairs)
	// fmt.Println("numbersToCheck:", numbersToCheck)

	slices.Sort(numbersToCheck)
	for _, num := range numbersToCheck {
		if rangeSet.hasValue(num) {
			// fmt.Println("fresh:", num)
			sum++
		}
	}

	for _, thePair := range rangeSet.sortedRanges {
		sum2 += thePair.end - thePair.begin + 1
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
