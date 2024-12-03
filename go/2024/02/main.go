package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	safe := 0
	for fileScanner.Scan() {
		report := strings.Fields(fileScanner.Text())
		prevLevel, err := strconv.Atoi(report[0])
		if err != nil {
			fmt.Println(err)
		}
		next, err := strconv.Atoi(report[1])
		if err != nil {
			fmt.Println(err)
		}

		reportDecreasing := (prevLevel - next) > 0

		// fmt.Println("", report, prevLevel, next, reportDecreasing)

		reportSafe := true
		for idx := 1; idx < len(report); idx++ {
			current, err := strconv.Atoi(report[idx])
			if err != nil {
				fmt.Println(err)
			}
			prev, err := strconv.Atoi(report[idx - 1])
			if err != nil {
				fmt.Println(err)
			}

			diff := prev - current
			levelDecreasing := diff > 0
			absDiff := int(math.Abs(float64(diff)))
			// fmt.Println("report ", diff, levelDecreasing)
			if levelDecreasing != reportDecreasing {
				reportSafe = false
				break
			}
			if absDiff < 1 {
				reportSafe = false
				break
			}
			if absDiff > 3 {
				reportSafe = false
				break
			}
		}
		if reportSafe {
			safe += 1
		}
	}
	fmt.Println("Answer to Day 2 Part 1: ", safe)
}


func checkSafe(report []string) bool {
	prevLevel, err := strconv.Atoi(report[0])
	if err != nil {
		fmt.Println(err)
	}
	next, err := strconv.Atoi(report[1])
	if err != nil {
		fmt.Println(err)
	}

	reportDecreasing := (prevLevel - next) > 0

	// fmt.Println("", report, prevLevel, next, reportDecreasing)

	reportSafe := true
	for idx := 1; idx < len(report); idx++ {
		current, err := strconv.Atoi(report[idx])
		if err != nil {
			fmt.Println(err)
		}
		prev, err := strconv.Atoi(report[idx - 1])
		if err != nil {
			fmt.Println(err)
		}

		diff := prev - current
		levelDecreasing := diff > 0
		absDiff := int(math.Abs(float64(diff)))
		// fmt.Println("report ", diff, levelDecreasing)
		if levelDecreasing != reportDecreasing {
			reportSafe = false
			break
		}
		if absDiff < 1 {
			reportSafe = false
			break
		}
		if absDiff > 3 {
			reportSafe = false
			break
		}
	}

	return reportSafe
}


func partTwo() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	safe := 0
	for fileScanner.Scan() {
		report := strings.Fields(fileScanner.Text())

		reportSlice := make([]string, len(report))
		copy(reportSlice, report)
		isSafe := checkSafe(reportSlice)
		if isSafe {
			safe += 1
			continue
		}

		// fmt.Println("REPORT ", report)
		for idx := 0; idx < len(report); idx++ {
			reportSlice := make([]string, len(report))
			copy(reportSlice, report)
			reportSlice = append(reportSlice[:idx], reportSlice[idx + 1:]...)

			isSafe := checkSafe(reportSlice)
			// fmt.Println(idx, reportSlice, isSafe)
			if isSafe {
				safe += 1
				break
			}
		}
	}
	fmt.Println("Answer to Day 2 Part 2: ", safe)
}
