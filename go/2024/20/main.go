package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	inputFile := os.Args[1]

	start := time.Now()
	partOne(inputFile)
	fmt.Println("Solved Part 1 in: ", time.Since(start))
	// start = time.Now()
	// partTwo(inputFile, size, numFallen)
	// fmt.Println("Solved Part 2 in: ", time.Since(start))
}

type Position struct {
	row, col int
}

func (p Position) Add(other Position) Position {
	return Position{p.row + other.row, p.col + other.col}
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var raceTrack [][]bool

	var start, end Position
	fileScanner := bufio.NewScanner(readFile)
	for row := 0; fileScanner.Scan(); row++ {
		line := strings.Split(fileScanner.Text(), "")
		var nextRow []bool
		for col, char := range line {
			nextRow = append(nextRow, char == "." || char == "S" || char == "E")
			if char == "S" {
				start = Position{row, col}
			} else if char == "E" {
				end = Position{row, col}
			}
		}
		raceTrack = append(raceTrack, nextRow)
	}

	// drive the track
	var path []Position
	current := start
	for current != end {
		path = append(path, current)
		for _, next := range []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			nextPos := current.Add(next)
			if len(path) > 1 && nextPos == path[len(path)-2] {
				continue
			}
			if raceTrack[nextPos.row][nextPos.col] {
				current = nextPos
				break
			}
		}
	}
	path = append(path, current)

	shortcuts := make(map[int]int)

	for trackNumber, current := range path {
		for _, next := range []Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			wall := current.Add(next)
			if raceTrack[wall.row][wall.col] {
				// does not pass thru a wall, not a shortcut
				continue
			}

			shortcut := wall.Add(next)
			savings := slices.Index(path[trackNumber+1:], shortcut)
			if savings != -1 {
				shortcuts[savings-1] += 1
				// fmt.Printf("Possible shortcut at %v, saving: %v\n", shortcut, savings - 1)
			}
		}
	}

	total := 0
	for k, v := range shortcuts {
		if k >= 100 {
			total += v
		}
	}
	fmt.Println("Answer to Day 20 Part 1: ", total)
	// fmt.Println("======================================")
	//
	// keys := make([]int, 0, len(shortcuts))
	// for k := range shortcuts {
	// 		keys = append(keys, k)
	// }
	//
	// sort.Ints(keys)
	//
	// for _, k := range keys {
	// 	fmt.Printf("There are %v cheats that save %v picoseconds\n", shortcuts[k], k)
	// }
}