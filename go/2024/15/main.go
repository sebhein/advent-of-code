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

type position struct {
	row int
	col int
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var warehouse [][]string
	var robot position
	fileScanner := bufio.NewScanner(readFile)
	for i := 0; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		warehouse = append(warehouse, strings.Split(line, ""))
		robotCol := strings.Index(line, "@")
		if robotCol != -1 {
			robot = position{i, robotCol}
		}
		if line == strings.Repeat("#", len(line)) && len(warehouse) > 1 {
			// bottom wall
			break
		}
	}

	// blank line between map and instructions
	fileScanner.Scan()

	for fileScanner.Scan() {
		instructions := fileScanner.Text()
		for _, instruction := range strings.Split(instructions, "") {
			executeInstruction(warehouse, &robot, instruction)
			// for _, row := range warehouse {
			// 	fmt.Println(row)
			// }
		}
	}

	gps := 0
	for row := 1; row < len(warehouse)-1; row++ {
		for col := 1; col < len(warehouse[0])-1; col++ {
			if warehouse[row][col] == "O" {
				gps += (100 * row) + col
			}
		}
	}

	fmt.Println("Answer to Day 15 Part 1: ", gps)
}

func makeMove(warehouse [][]string, from, to position) bool {
	if warehouse[to.row][to.col] == "#" {
		return false
	}

	if warehouse[to.row][to.col] == "." {
		warehouse[to.row][to.col] = warehouse[from.row][from.col]
		warehouse[from.row][from.col] = "."
		return true
	}

	if warehouse[to.row][to.col] == "O" {
		dr := to.row - from.row
		dc := to.col - from.col
		if makeMove(warehouse, position{to.row, to.col}, position{to.row + dr, to.col + dc}) {
			return makeMove(warehouse, from, to)
		}
	}

	return false
}

func executeInstruction(warehouse [][]string, robot *position, instruction string) {
	var dr, dc int
	switch instruction {
	case "^":
		dr = -1
		dc = 0
	case ">":
		dr = 0
		dc = 1
	case "v":
		dr = 1
		dc = 0
	case "<":
		dr = 0
		dc = -1
	default:
		panic("invalid instruction")
	}

	if makeMove(warehouse, *robot, position{robot.row + dr, robot.col + dc}) {
		robot.row += dr
		robot.col += dc
	}
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var warehouse [][]string
	var robot position
	fileScanner := bufio.NewScanner(readFile)
	for i := 0; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		row := []string{}
		for _, col := range strings.Split(line, "") {
			var left, right string
			switch col {
			case "@":
				left = col
				right = "."
			case "O":
				left = "["
				right = "]"
			case ".":
				left = col
				right = col
			case "#":
				left = col
				right = col
			default:
				panic("unrecognized character")

			}
			row = append(row, left)
			row = append(row, right)
		}
		warehouse = append(warehouse, row)
		robotCol := strings.Index(line, "@")
		if robotCol != -1 {
			robot = position{i, robotCol * 2}
		}
		if line == strings.Repeat("#", len(line)) && len(warehouse) > 1 {
			// bottom wall
			break
		}
	}

	// blank line between map and instructions
	fileScanner.Scan()
	// for _, row := range warehouse {
	// 	fmt.Println(row)
	// }

	for fileScanner.Scan() {
		instructions := fileScanner.Text()
		for _, instruction := range strings.Split(instructions, "") {
			executeDoubleInstruction(warehouse, &robot, instruction)
			// fmt.Println(instruction)
			// for _, row := range warehouse {
			// 	fmt.Println(row)
			// }
		}
	}

	gps := 0
	for row := 1; row < len(warehouse)-1; row++ {
		for col := 1; col < len(warehouse[0])-1; col++ {
			if warehouse[row][col] == "[" {
				gps += (100 * row) + col
			}
		}
	}

	fmt.Println("Answer to Day 15 Part 1: ", gps)
}

func canMakeMove(warehouse [][]string, from, to position) bool {
	nextSpace := warehouse[to.row][to.col]
	if nextSpace == "#" {
		return false
	}

	if nextSpace == "." {
		return true
	}

	if to.row-from.row == 0 {
		// horizontal
		return canMakeMove(warehouse, to, position{to.row, to.col + to.col - from.col})
	} else {
		// vertical
		dr := to.row - from.row
		if nextSpace == "[" {
			left := canMakeMove(warehouse, position{to.row, to.col}, position{to.row + dr, to.col})
			right := canMakeMove(warehouse, position{to.row, to.col + 1}, position{to.row + dr, to.col + 1})
			return left && right
		} else {
			left := canMakeMove(warehouse, position{to.row, to.col}, position{to.row + dr, to.col})
			right := canMakeMove(warehouse, position{to.row, to.col - 1}, position{to.row + dr, to.col - 1})
			return left && right
		}
	}
}

func move(warehouse [][]string, from, to position) {
	nextSpace := warehouse[to.row][to.col]
	if nextSpace == "#" {
		return
	}

	if nextSpace == "." {
		warehouse[to.row][to.col] = warehouse[from.row][from.col]
		warehouse[from.row][from.col] = "."
		return
	}

	if to.row-from.row == 0 {
		// horizontal
		move(warehouse, to, position{to.row, to.col + to.col - from.col})
		warehouse[to.row][to.col] = warehouse[from.row][from.col]
		warehouse[from.row][from.col] = "."
		return
	} else {
		// vertical
		dr := to.row - from.row
		if nextSpace == "[" {
			move(warehouse, position{to.row, to.col}, position{to.row + dr, to.col})
			move(warehouse, position{to.row, to.col + 1}, position{to.row + dr, to.col + 1})
		}
		if nextSpace == "]" {
			move(warehouse, position{to.row, to.col}, position{to.row + dr, to.col})
			move(warehouse, position{to.row, to.col - 1}, position{to.row + dr, to.col - 1})
		}

		warehouse[to.row][to.col] = warehouse[from.row][from.col]
		warehouse[from.row][from.col] = "."
	}
}

func executeDoubleInstruction(warehouse [][]string, robot *position, instruction string) {
	var dr, dc int
	switch instruction {
	case "^":
		dr = -1
		dc = 0
	case ">":
		dr = 0
		dc = 1
	case "v":
		dr = 1
		dc = 0
	case "<":
		dr = 0
		dc = -1
	default:
		panic("invalid instruction")
	}

	if canMakeMove(warehouse, *robot, position{robot.row + dr, robot.col + dc}) {
		move(warehouse, *robot, position{robot.row + dr, robot.col + dc})
		robot.row += dr
		robot.col += dc
	}
}
