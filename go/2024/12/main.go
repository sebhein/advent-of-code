package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inputFile := os.Args[1]
	start := time.Now()
	partOne(inputFile)
	fmt.Println("Solved Part 1 in: ", time.Since(start))
	// start = time.Now()
	// partTwo(inputFile)
	// fmt.Println("Solved Part 2 in: ", time.Since(start))
}

type Coord struct {
	row, col int
}

func measure(farm [][]string, visited [][]bool, plotLabel string, position Coord) [2]int {
	if position.row < 0 || position.row >= len(farm) || position.col < 0 || position.col >= len(farm[0]) {
		// out of bounds, put a fence on perimeter
		return [2]int{0, 1}
	}
	if farm[position.row][position.col] != plotLabel {
		// not part of this plot, put a fence on permieter
		return [2]int{0, 1}
	}
	if visited[position.row][position.col] {
		// already visited
		return [2]int{0, 0}
	}
	plotMeasures := [2]int{1, 0}
	visited[position.row][position.col] = true
	for _, offset := range []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		measurements := measure(farm, visited, plotLabel, Coord{position.row + offset.row, position.col + offset.col})
		plotMeasures[0] += measurements[0]
		plotMeasures[1] += measurements[1]
	}
	return plotMeasures
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var farm [][]string
	var visited [][]bool
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		farm = append(farm, strings.Split(fileScanner.Text(), ""))
		var r []bool
		for _ = range len(farm[0]) {
			r = append(r, false)
		}
		visited = append(visited, r)
	}
	
	totalPrice := 0
	for row := 0; row < len(farm); row++ {
		for col := 0; col < len(farm); col++ {
			m := measure(farm, visited, farm[row][col], Coord{row, col})
			totalPrice += m[0] * m[1]
		}
	}

	fmt.Println("Answer to Day 12 Part 1: ", totalPrice)
}
