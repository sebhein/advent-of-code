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
	start = time.Now()
	partTwo(inputFile)
	fmt.Println("Solved Part 2 in: ", time.Since(start))
}

type MapLoc struct {
	row int
	col int
}

type Direction struct {
	Facing string
	RowD   int
	ColD   int
}

func (d *Direction) DoTurn() {
	switch d.Facing {
	case "north":
		d.Facing = "east"
		d.RowD = 0
		d.ColD = 1
	case "east":
		d.Facing = "south"
		d.RowD = 1
		d.ColD = 0
	case "south":
		d.Facing = "west"
		d.RowD = 0
		d.ColD = -1
	case "west":
		d.Facing = "north"
		d.RowD = -1
		d.ColD = 0
	default:
		panic("unknown direction!")
	}
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var labMap [][]string

	// parse input
	fileScanner := bufio.NewScanner(readFile)
	visited := make(map[MapLoc]int)
	row := 0
	var start MapLoc
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if start == (MapLoc{}) {
			startCol := strings.Index(line, "^")
			if startCol != -1 {
				start = MapLoc{row, startCol}
			}
		}
		labMap = append(labMap, strings.Split(line, ""))
		row += 1
	}

	// fmt.Println(labMap)

	currentLoc := start
	currentDir := Direction{"north", -1, 0}
	var nextRow, nextCol int
	for {
		// fmt.Println("current ", currentLoc)
		visited[currentLoc] += 1
		nextRow = currentLoc.row + currentDir.RowD
		nextCol = currentLoc.col + currentDir.ColD
		if nextRow < 0 || nextRow >= len(labMap) || nextCol < 0 || nextCol >= len(labMap[0]) {
			// fmt.Println("left map")
			break
		}

		// fmt.Println("item ", labMap[nextRow][nextCol])
		if labMap[nextRow][nextCol] == "#" {
			currentDir.DoTurn()
			// fmt.Println("turned ", currentDir)
			nextRow = currentLoc.row + currentDir.RowD
			nextCol = currentLoc.col + currentDir.ColD
		}

		currentLoc = MapLoc{nextRow, nextCol}
	}

	fmt.Println("Answer to Day 6 Part 1: ", len(visited))
}

func WhereDoIGo(direction string) [2]int {
	switch direction {
	case "north":
		return [2]int{-1, 0}
	case "east":
		return [2]int{0, 1}
	case "south":
		return [2]int{1, 0}
	case "west":
		return [2]int{0, -1}
	default:
		panic("unknown direction!")
	}
}

func PrintMap(labMap [][]string, visited map[MapLoc][]string) {
	for row := 0; row < len(labMap); row++ {
		for col := 0; col < len(labMap[0]); col++ {
			curr, ok := visited[MapLoc{row, col}]
			if !ok {
				fmt.Printf(labMap[row][col])
				fmt.Printf(" ")
				continue
			}

			var horizontal, vertical bool
			if slices.Contains(curr, "east") || slices.Contains(curr, "west") {
				horizontal = true
			}
			if slices.Contains(curr, "south") || slices.Contains(curr, "north") {
				vertical = true
			}

			if horizontal && vertical {
				fmt.Printf("+ ")
			} else if horizontal {
				fmt.Printf("- ")
			} else {
				fmt.Printf("| ")
			}
		}
		fmt.Printf("\n")
	}
}

func tracePath(labMap [][]string, startingPosition MapLoc, startingDirection Direction, visited map[MapLoc][]string) bool {
	currentLoc := startingPosition
	currentDir := startingDirection
	var nextRow, nextCol int
	for {
		// fmt.Println("current ", currentLoc)
		directions, ok := visited[currentLoc]
		if ok {
			if slices.Contains(directions, currentDir.Facing) {
				return true
			}
		}

		visited[currentLoc] = append(visited[currentLoc], currentDir.Facing)

		nextRow = currentLoc.row + currentDir.RowD
		nextCol = currentLoc.col + currentDir.ColD
		if nextRow < 0 || nextRow >= len(labMap) || nextCol < 0 || nextCol >= len(labMap[0]) {
			// fmt.Println("left map")
			break
		}

		// fmt.Println("item ", labMap[nextRow][nextCol])
		for {
			if labMap[nextRow][nextCol] == "#" || labMap[nextRow][nextCol] == "O" {
				currentDir.DoTurn()
				visited[currentLoc] = append(visited[currentLoc], currentDir.Facing)
				nextRow = currentLoc.row + currentDir.RowD
				nextCol = currentLoc.col + currentDir.ColD
			} else {
				break
			}
		}

		currentLoc = MapLoc{nextRow, nextCol}
	}
	return false
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var labMap [][]string

	// parse input
	fileScanner := bufio.NewScanner(readFile)
	visited := make(map[MapLoc][]string)
	row := 0
	var start MapLoc
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if start == (MapLoc{}) {
			startCol := strings.Index(line, "^")
			if startCol != -1 {
				start = MapLoc{row, startCol}
			}
		}
		labMap = append(labMap, strings.Split(line, ""))
		row += 1
	}

	// fmt.Println(labMap)

	cycles := make(map[MapLoc]int)
	tracePath(labMap, start, Direction{"north", -1, 0}, visited)

	for location, directions := range visited {
		for _, direction := range directions {
			toPlace := WhereDoIGo(direction)
			rPlace := location.row + toPlace[0]
			cPlace := location.col + toPlace[1]

			if rPlace < 0 || rPlace >= len(labMap) || cPlace < 0 || cPlace >= len(labMap[0]) || labMap[rPlace][cPlace] == "#" || (rPlace == start.row && cPlace == start.col) {
				continue
			}
			var labMapCopy [][]string
			for r := 0; r < len(labMap); r++ {
				var rowCopy []string
				for c := 0; c < len(labMap[0]); c++ {
					rowCopy = append(rowCopy, labMap[r][c])
				}
				labMapCopy = append(labMapCopy, rowCopy)
			}

			// fmt.Println("placing @ ", rPlace, cPlace)

			labMapCopy[rPlace][cPlace] = "O"
			newVisited := make(map[MapLoc][]string)
			// PrintMap(labMapCopy, newVisited)
			// fmt.Println("->->->->->->->->->->->->->->->->->->->->->->->->->->->->->->->->->")
			cycle := tracePath(labMapCopy, start, Direction{"north", -1, 0}, newVisited)
			if cycle {
				cycles[MapLoc{rPlace, cPlace}] += 1
				// fmt.Println("found cycle!", MapLoc{rPlace, cPlace})
			}
			// PrintMap(labMapCopy, newVisited)
			// fmt.Println("=========================================================================")
		}
	}

	_, ok := cycles[start]
	fmt.Println("start in cycle loc ", ok)

	// fmt.Println("cycles ", cycles)

	// PrintMap(labMap, visited)

	fmt.Println("Answer to Day 6 Part 2: ", len(cycles))
}
