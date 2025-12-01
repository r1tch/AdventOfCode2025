package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

//     dial      num  dial+num pw
//        0    1..99     1..99 +0
//        0    100..     100.. +(num/100)
//       99        1       100 +1
//       99        2       101 +1
//       99      101       200 +2
// ---                 newdial
//        0       -1        99  +1
//        0     -101        99  +2
//        0     -199         1  +2
//        0     -200         0  +3
//        0    100..     100.. +(num/100)
//       99       -1        98 0
//       99      -99         0 +1
//       99     -101         0 +1

func main() {
	dial := 50
	pw := 0
	processLine := func(s string) {
		num, _ := strconv.Atoi(s[1:])
		if s[0] == 'L' {
			num *= -1
		}
		fmt.Printf("dial at: %d, got line:%s %d (pw:%d)\n", dial, s, num, pw)

		if num > 0 {
			pw += (num + dial) / 100
		} else if num * -1 >= dial { // are we reaching or crossing zero?
			if dial != 0 { // if we are not on 0, we need to also count for crossing it
				pw++
			}
			pw += (num+dial) / -100
		}

		dial = (dial + num) % 100
		if dial < 0 {
			dial += 100
		}
	}

	processLines("input.txt", processLine)
	fmt.Printf("pw: %d\n", pw)
}
