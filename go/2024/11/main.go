package main

import (
	"bufio"
	"fmt"
	"math"
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

	var stones []int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()
	for _, char := range strings.Fields(line) {
		num, _ := strconv.Atoi(char)
		stones = append(stones, num)
	}

	for iter := 0; iter < 25; iter++ {
		var nextArrangement []int
		for _, stone := range stones {
			if stone == 0 {
				nextArrangement = append(nextArrangement, 1)
				continue
			}
			length := countDigits(stone)
			if length%2 == 0 {
				half := length / 2
				powerOfTen := int(math.Pow(float64(10), float64(half)))
				nextArrangement = append(nextArrangement, stone/powerOfTen)
				nextArrangement = append(nextArrangement, stone%powerOfTen)
				continue
			}
			nextArrangement = append(nextArrangement, stone*2024)
		}
		stones = nextArrangement
	}

	fmt.Println("Answer to Day 11 Part 1: ", len(stones))
}

func countDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var stones []int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()
	for _, char := range strings.Fields(line) {
		num, _ := strconv.Atoi(char)
		stones = append(stones, num)
	}

	cache := make(map[int][]int)
	for iter := 0; iter < 75; iter++ {
		var nextArrangement []int
		for _, stone := range stones {
			_, ok := cache[stone]
			if ok {
				continue
			}

			if stone == 0 {
				cache[stone] = make([]int, 1)
				nextArrangement = append(nextArrangement, 1)
				continue
			}
			length := countDigits(stone)
			if length%2 == 0 {
				half := length / 2
				powerOfTen := int(math.Pow(float64(10), float64(half)))
				cache[stone] = []int{stone / powerOfTen, stone % powerOfTen}
				nextArrangement = append(nextArrangement, stone/powerOfTen)
				nextArrangement = append(nextArrangement, stone%powerOfTen)
				continue
			}
			cache[stone] = make([]int, stone*2024)
			nextArrangement = append(nextArrangement, stone*2024)
		}
		stones = nextArrangement
	}

	totalStones := 0
	for _, stone := range stones {
		for iter := 0; iter < 75; iter++ {
			var nextArrangement []int
			nextArrangement = append(nextArrangement, cache[stone]...)
		}
	}

	fmt.Println("Answer to Day 11 Part 2: ", totalStones)
}

func nextIteration(stone, currentIteration int) int {
	if currentIteration == 75 {
		return 1
	}

	if stone == 0 {
		return nextIteration(1, currentIteration+1)
	}
	length := countDigits(stone)
	if length%2 == 0 {
		half := length / 2
		powerOfTen := int(math.Pow(float64(10), float64(half)))
		num := nextIteration(stone/powerOfTen, currentIteration+1)
		return num + nextIteration(stone%powerOfTen, currentIteration+1)
	}

	return nextIteration(stone*2024, currentIteration+1)
}
