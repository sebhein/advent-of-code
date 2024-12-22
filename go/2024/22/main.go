package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
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

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	resultChan := make(chan int)
	doneChan := make(chan int)
	var wg sync.WaitGroup
	go addSecrets(resultChan, doneChan)

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		seed, _ := strconv.Atoi(fileScanner.Text())
		wg.Add(1)
		go calculateSecret(seed, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	fmt.Println("Answer to Day 22 Part 1: ", <-doneChan)
}

func addSecrets(resultChan <-chan int, done chan<- int) {
	total := 0
	for result := range resultChan {
		total += result
	}
	done <- total
}

func calculateSecret(seed int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var result, secret int
	secret = seed
	for i := 0; i < 2000; i++ {
		result = secret

		result *= 64
		result ^= secret
		result %= 16777216

		secret = result

		result /= 32
		result ^= secret
		result %= 16777216

		secret = result

		result *= 2048
		result ^= secret
		result %= 16777216

		secret = result
	}
	// fmt.Printf("Seed %v produced %v\n", seed, secret)
	resultChan <- result
}
