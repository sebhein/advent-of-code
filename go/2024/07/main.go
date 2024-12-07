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

func check(target int, vals []int) bool {
	checkVal := vals[len(vals)-1]
	rest := vals[:len(vals)-1]
	if len(rest) == 0 {
		return checkVal == target
	}

	for _, op := range []string{"add", "mult", "concat"} {
		var newTarget int
		switch op {
		case "add":
			if target-checkVal < 0 {
				continue
			}
			newTarget = target - checkVal
		case "mult":
			if target%checkVal != 0 {
				continue
			}
			newTarget = target / checkVal
		case "concat":
			tarStr := strconv.Itoa(target)
			checkStr := strconv.Itoa(checkVal)
			if !strings.HasSuffix(tarStr, checkStr) {
				continue
			}
			newTarget, _ = strconv.Atoi(tarStr[:len(tarStr)-len(checkStr)])
		}
		if check(newTarget, rest) {
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
