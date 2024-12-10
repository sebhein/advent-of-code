package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFile := os.Args[1]
	start := time.Now()
	partOne(inputFile)
	fmt.Println("Solved Part 1 in: ", time.Since(start))
	start = time.Now()
	partTwo(inputFile)
	fmt.Println("Solved Part 2 in: ", time.Since(start))
}

type Coord struct {
	row int
	col int
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var topMap [][]int
	var trailheads []Coord
	fileScanner := bufio.NewScanner(readFile)
	for row := 0; fileScanner.Scan(); row++ {
		line := strings.Split(fileScanner.Text(), "")
		var lineAsNums []int
		for col, char := range line {
			num, _ := strconv.Atoi(char)
			lineAsNums = append(lineAsNums, num)
			if num == 0 {
				trailheads = append(trailheads, Coord{row, col})
			}
		}
		topMap = append(topMap, lineAsNums)
	}

	height := len(topMap)
	width := len(topMap[0])

	total := 0
	for _, trailhead := range trailheads {
		trailheadsFound := make(map[Coord]int)
		followTrail(topMap, trailhead, height, width, trailheadsFound)
		total += len(trailheadsFound)
	}

	fmt.Println("Answer to Day 10 Part 1: ", total)
}

func inBounds(pos Coord, height, width int) bool {
	if pos.row < 0 || pos.row >= height {
		return false
	}
	if pos.col < 0 || pos.col >= width {
		return false
	}
	return true
}

var DIRECTIONS = [4]Coord{{-1, 0}, {1, 0},{0, -1},{0, 1}}

func followTrail(topMap [][]int, pos Coord, height, width int, trailheadsFound map[Coord]int) {
	if topMap[pos.row][pos.col] == 9 {
		trailheadsFound[pos] += 1
	}

	for _, d := range DIRECTIONS {
		next := Coord{pos.row + d.row, pos.col + d.col}
		if !inBounds(next, height, width) {
			continue
		}
		if topMap[next.row][next.col] == topMap[pos.row][pos.col] + 1 {
			followTrail(topMap, next, height, width, trailheadsFound)
		}
	}
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var topMap [][]int
	var trailheads []Coord
	fileScanner := bufio.NewScanner(readFile)
	for row := 0; fileScanner.Scan(); row++ {
		line := strings.Split(fileScanner.Text(), "")
		var lineAsNums []int
		for col, char := range line {
			num, _ := strconv.Atoi(char)
			lineAsNums = append(lineAsNums, num)
			if num == 0 {
				trailheads = append(trailheads, Coord{row, col})
			}
		}
		topMap = append(topMap, lineAsNums)
	}

	height := len(topMap)
	width := len(topMap[0])

	total := 0
	for _, trailhead := range trailheads {
		trailheadsFound := make(map[Coord]int)
		followTrail(topMap, trailhead, height, width, trailheadsFound)
		for _, paths := range trailheadsFound {
			total += paths
		}
	}

	fmt.Println("Answer to Day 10 Part 2: ", total)
}
