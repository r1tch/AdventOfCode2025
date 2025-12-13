package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	begin int
	end int
}

type point struct {
	x, y, z int
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

func distance(a, b point) float64 {
	return math.Sqrt(float64(
		(a.x - b.x) ^ 2 +
		(a.y - b.y) ^ 2 +
		(a.z - b.z) ^ 2))
}

func runTests() {
}

func main() {
	runTests()

	sum := 0
	sum2 := 0
	points := make([]point, 0)
	processLines("inputt.txt", func(s string) {
		coordsStr := strings.Split(s, ",")
		x, _ := strconv.Atoi(coordsStr[0])
		y, _ := strconv.Atoi(coordsStr[1])
		z, _ := strconv.Atoi(coordsStr[2])
		points = append(points, point{x, y, z})
	})
	fmt.Println(points)

	// make X shortest connections
	// connected ones become a circuit
	//   --> array of sets of points?

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
