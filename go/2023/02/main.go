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
	go sumPossibleIds(resultChan, doneChan)

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		wg.Add(1)
		go determinePossible(fileScanner.Text(), resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	fmt.Println("Answer to Day 02 Part 1: ", <-doneChan)
}

func sumPossibleIds(resultChan <-chan int, done chan<- int) {
	total := 0
	for result := range resultChan {
		total += result
	}
	done <- total
}

func determinePossible(game string, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	splitGame := strings.Split(game, ":")
	reveals := strings.Split(splitGame[1], ";")
	for _, reveal := range reveals {
		cubes := strings.Split(reveal, ",")
		for _, cube := range cubes {
			fields := strings.Fields(cube)
			num, _ := strconv.Atoi(fields[0])
			switch fields[1] {
			case "red":
				if num > 12 {
					return
				}
			case "green":
				if num > 13 {
					return
				}
			case "blue":
				if num > 14 {
					return
				}
			}
		}
	}

	id, _ := strconv.Atoi(strings.Fields(splitGame[0])[1])
	resultChan <- id
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	resultChan := make(chan int)
	doneChan := make(chan int)
	var wg sync.WaitGroup
	go sumOfPowers(resultChan, doneChan)

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		wg.Add(1)
		go leastCubes(fileScanner.Text(), resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	fmt.Println("Answer to Day 02 Part 2: ", <-doneChan)
}

func sumOfPowers(resultChan <-chan int, done chan<- int) {
	total := 0
	for result := range resultChan {
		total += result
	}
	done <- total
}

func leastCubes(game string, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	splitGame := strings.Split(game, ":")
	reveals := strings.Split(splitGame[1], ";")
	lowest := [3]int{0, 0, 0}
	for _, reveal := range reveals {
		cubes := strings.Split(reveal, ",")
		for _, cube := range cubes {
			fields := strings.Fields(cube)
			num, _ := strconv.Atoi(fields[0])
			switch fields[1] {
			case "red":
				if num > lowest[0] {
					lowest[0] = num
				}
			case "green":
				if num > lowest[1] {
					lowest[1] = num
				}
			case "blue":
				if num > lowest[2] {
					lowest[2] = num
				}
			}
		}
	}

	resultChan <- lowest[0] * lowest[1] * lowest[2]
}
