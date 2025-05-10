package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"sync"
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

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var keys [][5]int
	var locks [][5]int
	var current [5]int
	fileScanner := bufio.NewScanner(readFile)
	for idx := 0; fileScanner.Scan(); idx++ {
		line := fileScanner.Text()
		if idx % 8 == 0 {
			// next key or lock
			current = [5]int{0, 0, 0, 0, 0}
			continue
		}
		if idx % 8 == 7 {
			// blank - don't care
			continue
		}
		if idx % 8 == 6 {
			// put in keys or locks
			if line[0] == '#' {
				keys = append(keys, current)
			} else {
				locks = append(locks, current)
			}
			continue
		}
		for si, c := range line {
			if c == '#' {
				current[si] += 1
			}
		}
	}

	resultChan := make(chan int)
	doneChan := make(chan int)
	var wg sync.WaitGroup
	go countFits(resultChan, doneChan)

	for _, key := range keys {
		wg.Add(1)
		checkFit(key, &locks, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	fmt.Println("Answer to Day 25 Part 1: ", <-doneChan)
}

func checkFit(key [5]int, locks *[][5]int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	totalUnique := 0
	for _, lock := range *locks {
		unique := true
		for idx, kp := range key {
			if kp + lock[idx] > 5 {
				unique = false
			}
		}
		if unique {
			totalUnique += 1
		}
	}
	result<-totalUnique
}

func countFits(resultChan <-chan int, done chan<- int) {
	total := 0
	for result := range resultChan {
		total += result
	}
	done <- total
}
