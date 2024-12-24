package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	x int
	y int
}

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

type NumPad struct {
	currentPosition Coord
}

func (kp *NumPad) press(target rune) string {
	var targetPosition Coord
	switch target {
	// Row 0
	case '0': targetPosition = Coord{1, 0}
	case 'A': targetPosition = Coord{2, 0}
	// Row 1
	case '1': targetPosition = Coord{0, 1}
	case '2': targetPosition = Coord{1, 1}
	case '3': targetPosition = Coord{2, 1}
	// Row 2
	case '4': targetPosition = Coord{0, 2}
	case '5': targetPosition = Coord{1, 2}
	case '6': targetPosition = Coord{2, 2}
	// Row 3
	case '7': targetPosition = Coord{0, 3}
	case '8': targetPosition = Coord{1, 3}
	case '9': targetPosition = Coord{2, 3}
	}

	var sequence string
	last := -1
	for targetPosition != kp.currentPosition {
		if last == 1 {
			xDiff := targetPosition.x - kp.currentPosition.x
			if xDiff > 0 {
				kp.currentPosition.x += 1
				sequence += ">"
				last = 1
				continue
			}
			if xDiff < 0 && (kp.currentPosition.y != 0 || kp.currentPosition.x - 1 != 0){
				kp.currentPosition.x -= 1
				sequence += "<"
				last = 1
				continue
			}
			last = -1
		} else {
			yDiff := targetPosition.y - kp.currentPosition.y
			if yDiff > 0 {
				kp.currentPosition.y += 1
				sequence += "^"
				last = -1
				continue
			}
			if yDiff < 0 && (kp.currentPosition.x != 0 || kp.currentPosition.y - 1 !=0) {
				kp.currentPosition.y -= 1
				sequence += "v"
				last = -1
				continue
			}
			last = 1
		}
	}
	return sequence + "A"
}

func (kp *NumPad) enter(input string) string {
	var sequence string
	for _, char := range input {
		sequence += kp.press(char)
	}
	return sequence
}

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

type DirPad struct {
	currentPosition Coord
}

func (kp *DirPad) press(target rune) string {
	var targetPosition Coord
	switch target {
	// Row 0
	case '<': targetPosition = Coord{0, 0}
	case 'v': targetPosition = Coord{1, 0}
	case '>': targetPosition = Coord{2, 0}
	// Row 1
	case '^': targetPosition = Coord{1, 1}
	case 'A': targetPosition = Coord{2, 1}
	}

	var sequence string
	preferRepeat := 1
	badCoord := Coord{0, 1}
	for targetPosition != kp.currentPosition {
		if kp.currentPosition == badCoord {
			panic("No this direction position is not allowed!")
		}
		xDiff := targetPosition.x - kp.currentPosition.x
		yDiff := targetPosition.y - kp.currentPosition.y
		if preferRepeat == 1 {
			if xDiff > 0 {
				kp.currentPosition.x += 1
				sequence += ">"
				preferRepeat = 1
				continue
			}
			preferRepeat = 2
		}
		if preferRepeat == 2 {
			if xDiff < 0 && (kp.currentPosition.y != 1 || kp.currentPosition.x - 1 != 0){
				kp.currentPosition.x -= 1
				sequence += "<"
				preferRepeat = 2
				continue
			}
			preferRepeat = 3
		}
		if preferRepeat == 3 {
			if yDiff > 0 && (kp.currentPosition.x != 0 || kp.currentPosition.y + 1 !=1) {
				kp.currentPosition.y += 1
				sequence += "^"
				preferRepeat = 3
				continue
			}
			preferRepeat = 4
		}
		if preferRepeat == 4 {
			if yDiff < 0 {
				kp.currentPosition.y -= 1
				sequence += "v"
				preferRepeat = 4
				continue
			}
			preferRepeat = 1
		}
	}
	return sequence + "A"
}

func (kp *DirPad) enter(input string) string {
	var sequence string
	for _, char := range input {
		sequence += kp.press(char)
	}
	return sequence
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	complexity := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()

		np := NumPad{Coord{2, 0}}
		numSequence := np.enter(input)

		dpOne := DirPad{Coord{2, 1}}
		dirSequence := dpOne.enter(numSequence)

		dpTwo := DirPad{Coord{2, 1}}
		me := dpTwo.enter(dirSequence)

		num, _ := strconv.Atoi(input[:len(input)-1])
		complexity += num * len(me)
		fmt.Println(me)
		fmt.Println(dirSequence)
		fmt.Println(numSequence)
		fmt.Println(input)
		fmt.Printf("input: %v len shortest? %v my answer: %v\n", input, len(me), me)
	}

	fmt.Println("Answer to Day 21 Part 1: ", complexity)
}
