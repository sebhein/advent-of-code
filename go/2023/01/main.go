package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	fileScanner := bufio.NewScanner(readFile)

	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var first, last string
		revIdx := len(line) - 1
		for idx := 0; idx < len(line); idx++ {
			if len(first) > 0 && len(last) > 0 {
				break
			}
			if unicode.IsDigit(rune(line[idx])) && len(first) == 0 {
				first = string(line[idx])
			}
			if unicode.IsDigit(rune(line[revIdx-idx])) && len(last) == 0 {
				last = string(line[revIdx-idx])
			}
		}
		calibrationValue, _ := strconv.Atoi(first + last)
		total += calibrationValue
	}

	fmt.Println("Answer to Day 01 Part 1: ", total)
}

func convertToNumber(in string) string {
	switch in {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return in
	}
}

func allIndicies(toCheck, subStr string) []int {
	var indicies []int
	start := 0

	for {
		idx := strings.Index(toCheck[start:], subStr)
		if idx == -1 {
			break
		}
		indicies = append(indicies, start+idx)
		start += idx + 1
	}
	return indicies
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var first, last string
		bestFirst := 100
		bestLast := 0
		for _, number := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			positions := allIndicies(line, number)
			if len(positions) == 0 {
				continue
			}
			for _, pos := range positions {
				if pos <= bestFirst {
					first = number
					bestFirst = pos
				}
				if pos >= bestLast {
					last = number
					bestLast = pos
				}
			}
		}
		calibrationValue, _ := strconv.Atoi(convertToNumber(first) + convertToNumber(last))
		total += calibrationValue
	}

	fmt.Println("Answer to Day 01 Part 2: ", total)
}
