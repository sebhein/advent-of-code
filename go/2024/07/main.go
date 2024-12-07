package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

	resultChan := make(chan int)
	doneChan := make(chan int)
	var wg sync.WaitGroup

	go processCalculations(resultChan, doneChan)

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		targetAndValues := strings.Split(line, ":")
		target, _ := strconv.Atoi(targetAndValues[0])
		var values []int
		for _, strNumber := range strings.Split(targetAndValues[1], " ") {
			number, _ := strconv.Atoi(strNumber)
			values = append(values, number)
		}
		wg.Add(1)
		calibrateEquation(target, values, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	fmt.Println("Answer to Day 7 Part 1: ", <-doneChan)
}

func processCalculations(resultChan <-chan int, done chan<- int) {
	calibrationResult := 0
	for result := range resultChan {
		calibrationResult += result
	}
	done <- calibrationResult
}

func calibrateEquation(target int, opValues []int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	numOperators := len(opValues) - 1
	for i := 0; i < (1 << numOperators); i++ {
		calculated := opValues[0]
		for j := 0; j < numOperators; j++ {
			if (i & (1 << j)) != 0 {
				calculated += opValues[j+1]
			} else {
				calculated *= opValues[j+1]
			}
			if calculated > target {
				break
			}
		}
		if calculated == target {
			resultChan <- target
			return
		}
	}
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	resultChan := make(chan int)
	doneChan := make(chan int)
	var wg sync.WaitGroup

	go processCalculations(resultChan, doneChan)

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		targetAndValues := strings.Split(line, ":")
		target, _ := strconv.Atoi(targetAndValues[0])
		var values []int
		for _, strNumber := range strings.Split(strings.TrimSpace(targetAndValues[1]), " ") {
			number, _ := strconv.Atoi(strNumber)
			values = append(values, number)
		}
		wg.Add(1)
		calibrateEquationConcatenation(target, values, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	fmt.Println("Answer to Day 7 Part 2: ", <-doneChan)
}

func powerOfTen(num int) int {
	if num < 10 {
		return 10
	} else if num < 100 {
		return 100
	} else if num < 1000 {
		return 1000
	} else if num < 10000 {
		return 10000
	}
	return 0
}

func check(target int, vals []int) bool {
	checkVal := vals[len(vals)-1]
	rest := vals[:len(vals)-1]
	if len(rest) == 0 {
		return checkVal == target
	}

	if target-checkVal > 0 {
		if check(target-checkVal, rest) {
			return true
		}
	}

	if target%checkVal == 0 {
		if check(target/checkVal, rest) {
			return true
		}
	}

	if target%powerOfTen(checkVal) == checkVal {
		if check(target/powerOfTen(checkVal), rest) {
			return true
		}
	}

	return false
}

func calibrateEquationConcatenation(target int, opValues []int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	if check(target, opValues) {
		resultChan <- target
	}
}
