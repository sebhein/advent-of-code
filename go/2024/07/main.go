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
				calculated += opValues[j + 1]
			} else {
				calculated *= opValues[j + 1]
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
		for _, strNumber := range strings.Split(targetAndValues[1], " ") {
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


func calibrateEquationConcatenation(target int, opValues []int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	numOperators := len(opValues) - 1
	possibilities := 1
	for i := 0; i < numOperators; i++ {
		possibilities *= 3
	}

	for i := 0; i < possibilities; i++ {
		current := i
		calculated := opValues[0]
		for j := 0; j < numOperators; j++ {
			op := (current % 3) - 1
			switch op {
			case -1:
				calculated += opValues[j + 1]
			case 0:
				calculated *= opValues[j + 1]
			case 1:
				concatedStr := strconv.Itoa(calculated) + strconv.Itoa(opValues[j + 1])
				concatedInt, _ := strconv.Atoi(concatedStr)
				calculated = concatedInt
			}
			current /= 3
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
