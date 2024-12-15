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

type robot struct {
	row int
	col int
	dr  int
	dc  int
}

func moveRobot(robo *robot, floorHeight, floorWidth int) {
	// fmt.Println("moving robot: ", robo)
	newRow := robo.row + robo.dr
	if newRow < 0 {
		newRow += floorHeight
	} else if newRow >= floorHeight {
		newRow = 0 + newRow - floorHeight
	}
	newCol := robo.col + robo.dc
	if newCol < 0 {
		newCol += floorWidth
	} else if newCol >= floorWidth {
		newCol = 0 + newCol - floorWidth
	}
	robo.row = newRow
	robo.col = newCol
	// fmt.Println("moved to: ", robo)
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var robots []robot
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		split := strings.Fields(fileScanner.Text())
		pos := strings.Split(split[0][2:], ",")
		vel := strings.Split(split[1][2:], ",")
		col, _ := strconv.Atoi(pos[0])
		row, _ := strconv.Atoi(pos[1])
		dc, _ := strconv.Atoi(vel[0])
		dr, _ := strconv.Atoi(vel[1])
		robots = append(robots, robot{row, col, dr, dc})
	}

	var floorHeight, floorWidth int
	if inputFile == "ex.txt" {
		floorHeight = 7
		floorWidth = 11
	} else {
		floorHeight = 103
		floorWidth = 101
	}

	quadrants := [4]int{}
	for _, robo := range robots {
		// fmt.Println("======================================================")
		for ts := 0; ts < 100; ts++ {
			moveRobot(&robo, floorHeight, floorWidth)
		}
		if robo.row < floorHeight/2 {
			if robo.col < floorWidth/2 {
				quadrants[0] += 1
			} else if robo.col > floorWidth/2 {
				quadrants[1] += 1
			}
		} else if robo.row > floorHeight/2 {
			if robo.col > floorWidth/2 {
				quadrants[2] += 1
			} else if robo.col < floorWidth/2 {
				quadrants[3] += 1
			}
		}
	}

	// fmt.Println(quadrants)
	// fmt.Println(robots)

	safetyFactor := 1
	for _, q := range quadrants {
		safetyFactor *= q
	}

	fmt.Println("Answer to Day 14 Part 1: ", safetyFactor)
}

type loc struct {
	row int
	col int
}

func partTwo(inputFile string) {
	if inputFile == "ex.txt" {
		fmt.Println("nah part 2 wont work with example input")
	}
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var robots []robot
	fileScanner := bufio.NewScanner(readFile)
	for i := 1; fileScanner.Scan(); i++ {
		split := strings.Fields(fileScanner.Text())
		pos := strings.Split(split[0][2:], ",")
		vel := strings.Split(split[1][2:], ",")
		col, _ := strconv.Atoi(pos[0])
		row, _ := strconv.Atoi(pos[1])
		dc, _ := strconv.Atoi(vel[0])
		dr, _ := strconv.Atoi(vel[1])
		robots = append(robots, robot{row, col, dr, dc})
	}

	floorHeight := 103
	floorWidth := 101

	ts := 0
	for {
		ts++
		// fmt.Println("Timestamp = ", ts)
		// time.Sleep(time.Second)
		robotNeighbors := make(map[loc]int)
		for i := 0; i < len(robots); i++ {
			moveRobot(&robots[i], floorHeight, floorWidth)
			robotNeighbors[loc{robots[i].row, robots[i].col}] = 0
		}

		// for x := 0; x < floorWidth; x++ {
		// 	for y := 0; y < floorWidth; y++ {
		// 		_, ok := robotNeighbors[loc{x, y}]
		// 		if ok {
		// 			fmt.Print("X")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Print("\n")
		// }

		for robo := range robotNeighbors {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					_, ok := robotNeighbors[loc{robo.row + dr, robo.col + dc}]
					if ok {
						robotNeighbors[loc{robo.row + dr, robo.col + dc}] += 1
					}
				}
			}
		}

		total := 0
		for _, numNeighbors := range robotNeighbors {
			if numNeighbors > 6 {
				total += 1
			}
		}
		// fmt.Println("total = ", total)
		if total >= 150 {
			break
		}
	}

	fmt.Println("Answer to Day 14 Part 2: ", ts)
}
