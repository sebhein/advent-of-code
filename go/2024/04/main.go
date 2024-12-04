package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	inputFile := os.Args[1]
	partOne(inputFile)
	partTwo(inputFile)
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var puzzle [][]string
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), "")
		puzzle = append(puzzle, line)
	}

	counts := 0
	maxY := len(puzzle)
	maxX := len(puzzle[0])
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// fmt.Println("current char: ", puzzle[y][x])
			if puzzle[y][x] != "X" {
				continue
			}
			counts += searchAreaPartOne(puzzle, y, x, maxY, maxX)
		}
	}

	fmt.Println("Answer to Day 4 Part 1: ", counts)
}


func searchAreaPartOne(puzzle [][]string, locY, locX, maxY, maxX int) int {
	inBounds := func(x, y int) bool {
		if y < 0 || y >= maxY {
			return false
		}
		if x < 0 || x >= maxX {
			return false
		}
		return true
	}

	// fmt.Println("searching around")

	targetWord := []string{"M", "A", "S"}
	found := 0
	for dy := -1; dy <= 1; dy++ {
		if locY + dy < 0 || locY + dy >= maxY {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			if locX + dx < 0 || locX + dx >= maxX {
				continue
			}
			fullWord := true
			for dist := 1; dist <= len(targetWord); dist++ {
				checkY := locY + dy * dist
				checkX := locX + dx * dist
				if !inBounds(checkX, checkY) {
					fullWord = false
					break
				}
				char := targetWord[dist-1]
				// fmt.Println("Y coords: ", locY, dy, checkY, dist)
				// fmt.Println("X coords: ", locX, dx, checkX, dist)
				// fmt.Println("checking this: ", dy, dx, dist, puzzle[checkY][checkX], char)
				if puzzle[checkY][checkX] != char {
					fullWord = false
					break
				}
			}
			if fullWord {
				found++
			}
		}
	}
	return found
}


func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var puzzle [][]string
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), "")
		puzzle = append(puzzle, line)
	}

	counts := 0
	maxY := len(puzzle)
	maxX := len(puzzle[0])
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// fmt.Println("current char: ", puzzle[y][x])
			if puzzle[y][x] != "A" {
				continue
			}
			counts += searchArea(puzzle, y, x, maxY, maxX)
		}
	}

	fmt.Println("Answer to Day 4 Part 2: ", counts)
}


func searchArea(puzzle [][]string, locY, locX, maxY, maxX int) int {
	inBounds := func(x, y int) bool {
		if y < 0 || y >= maxY {
			return false
		}
		if x < 0 || x >= maxX {
			return false
		}
		return true
	}

	// fmt.Println("searching around ", locY, locX)

	targetWord := []string{"M", "M", "S", "S"}
	for rotation := 0; rotation < 4; rotation++ {
		left := locX - 1
		right := locX + 1
		up := locY - 1
		down := locY + 1

		corners := [][]int {{left, up}, {right, up}, {right, down}, {left, down}}
		
		if !inBounds(left, up) || !inBounds(right, up) || !inBounds(right, down) || !inBounds(left, down) {
			break
		}

		// println("in rotation: ", rotation)

		allCorners := true
		for toCheck := 0; toCheck < 4; toCheck++ {
			charToCheck := toCheck + rotation
			if charToCheck > 3 {
				charToCheck = charToCheck - 4
			}
			checkX := corners[toCheck][0]
			checkY := corners[toCheck][1]
			// println("checking: ", toCheck, charToCheck, checkX, checkY)
			// println("comparing: ", targetWord[charToCheck], puzzle[checkY][checkX])
			if targetWord[charToCheck] != puzzle[checkY][checkX] {
				allCorners = false
				break
			}

		}

		if allCorners {
			return 1
		}
	}
	return 0
}
