package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	allOrdering := make(map[int][]int)
	var updates [][]int
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "|") {
			order := strings.Split(line, "|")
			before, _ := strconv.Atoi(order[0])
			after, _ := strconv.Atoi(order[1])
			allOrdering[before] = append(allOrdering[before], after)
		} else if strings.Contains(line, ",") {
			strUpdates := strings.Split(line, ",")
			intUpdates := []int{}
			for _, strNumber := range strUpdates {
				number, _ := strconv.Atoi(strNumber)
				intUpdates = append(intUpdates, number)
			}
			updates = append(updates, intUpdates)
		}
	}

	sum := 0
	for _, update := range updates {
		ordered := true
		for idx, page := range update {
			for searchIdx := idx + 1; searchIdx < len(update); searchIdx++ {
				if slices.Contains(allOrdering[update[searchIdx]], page) {
					ordered = false
					break
				}
			}
			if !ordered {
				break
			}
		}
		if ordered {
			sum += update[(len(update)-1)/2]
		}
	}

	fmt.Println("Answer to Day 5 Part 1: ", sum)
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	allOrdering := make(map[int][]int)
	var updates [][]int
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "|") {
			order := strings.Split(line, "|")
			before, _ := strconv.Atoi(order[0])
			after, _ := strconv.Atoi(order[1])
			allOrdering[before] = append(allOrdering[before], after)
		} else if strings.Contains(line, ",") {
			strUpdates := strings.Split(line, ",")
			intUpdates := []int{}
			for _, strNumber := range strUpdates {
				number, _ := strconv.Atoi(strNumber)
				intUpdates = append(intUpdates, number)
			}
			updates = append(updates, intUpdates)
		}
	}

	var unordered [][]int
	for _, update := range updates {
		ordered := true
		for idx, page := range update {
			for searchIdx := idx + 1; searchIdx < len(update); searchIdx++ {
				if slices.Contains(allOrdering[update[searchIdx]], page) {
					ordered = false
					break
				}
			}
			if !ordered {
				break
			}
		}
		if !ordered {
			unordered = append(unordered, update)
		}
	}

	sum := 0
	for _, bad := range unordered {
		slices.SortFunc(bad, func(a, b int) int {
			shouldComeBefore := slices.Contains(allOrdering[a], b)
			if shouldComeBefore {
				return -1
			}
			return 1
		})
		sum += bad[(len(bad)-1)/2]
	}

	fmt.Println("Answer to Day 5 Part 1: ", sum)
}
