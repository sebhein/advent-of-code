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

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	tokens := 0
	var a1, b1, c1, a2, b2, c2 int
	fileScanner := bufio.NewScanner(readFile)
	for i := 1; fileScanner.Scan(); i++ {
		if i%4 == 0 {
			continue
		}
		line := fileScanner.Text()
		if i%4 == 1 {
			// first coefficient
			a1, _ = strconv.Atoi(line[11:14])
			a2, _ = strconv.Atoi(line[17:20])
			continue
		}
		if i%4 == 2 {
			// second coefficient
			b1, _ = strconv.Atoi(line[11:14])
			b2, _ = strconv.Atoi(line[17:20])
			continue
		}
		if i%4 == 3 {
			// equals
			splitLine := strings.Fields(line)
			c1, _ = strconv.Atoi(strings.Split(splitLine[1][:len(splitLine[1])-1], "=")[1])
			c2, _ = strconv.Atoi(strings.Split(splitLine[2], "=")[1])
			det := float64(a1*b2 - a2*b1)
			if det == 0.0 {
				// lines are parrallel
				continue
			}
			x := float64(c1*b2-c2*b1) / det
			y := float64(a1*c2-a2*c1) / det
			if x > 100 || x < 0 || y < 0 || y > 100 {
				continue
			}
			if x != float64(int(x)) || y != float64(int(y)) {
				// not whole numbers
				continue
			}
			tokens += int(3*x) + int(y)
			continue
		}
	}
	fmt.Println("Answer to Day 13 Part 1: ", tokens)
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	tokens := 0
	var a1, b1, c1, a2, b2, c2 int
	fileScanner := bufio.NewScanner(readFile)
	for i := 1; fileScanner.Scan(); i++ {
		if i%4 == 0 {
			continue
		}
		line := fileScanner.Text()
		if i%4 == 1 {
			// first coefficient
			a1, _ = strconv.Atoi(line[11:14])
			a2, _ = strconv.Atoi(line[17:20])
			continue
		}
		if i%4 == 2 {
			// second coefficient
			b1, _ = strconv.Atoi(line[11:14])
			b2, _ = strconv.Atoi(line[17:20])
			continue
		}
		if i%4 == 3 {
			// equals
			splitLine := strings.Fields(line)
			c1, _ = strconv.Atoi(strings.Split(splitLine[1][:len(splitLine[1])-1], "=")[1])
			c2, _ = strconv.Atoi(strings.Split(splitLine[2], "=")[1])
			c1 += 10000000000000
			c2 += 10000000000000
			det := float64(a1*b2 - a2*b1)
			if det == 0.0 {
				// lines are parrallel
				continue
			}
			x := float64(c1*b2-c2*b1) / det
			y := float64(a1*c2-a2*c1) / det
			if x < 0 || y < 0 {
				continue
			}
			if x != float64(int(x)) || y != float64(int(y)) {
				// not whole numbers
				continue
			}
			tokens += int(3*x) + int(y)
			continue
		}
	}
	fmt.Println("Answer to Day 13 Part 2: ", tokens)
}
