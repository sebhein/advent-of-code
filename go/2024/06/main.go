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
	RowD int
	ColD int
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

func (d Direction) PeekTurn() [2]int {
	switch d.Facing {
	case "north":
		return [2]int{0, 1}
	case "east":
		return [2]int{1, 0}
	case "south":
		return [2]int{0, -1}
	case "west":
		return [2]int{-1, 0}
	default:
		panic("unknown direction!")
	}
}


func partTwo(inputFile string) {
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
	loops := 0
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

		_, ok := visited[currentLoc]
		if ok {
			peek := currentDir.PeekTurn()
			_, ok := visited[MapLoc{peek[0], peek[1]}]
			if ok {
				loops += 1
			}
		}

	}

	fmt.Println("Answer to Day 6 Part 2: ", loops)
}
