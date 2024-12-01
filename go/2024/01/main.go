package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
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

	var left, right []int
	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		l, err := strconv.Atoi(line[0]);

		if err != nil {
			fmt.Println("not an integer")
		}

		r, err := strconv.Atoi(line[1]);
		if err != nil {
			fmt.Println("not an integer")
		}

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	diffSum := 0
	for idx := 0; idx < len(left); idx++ {
		diff := left[idx] - right[idx]
		if diff < 0 {
			diffSum -= diff
		} else {
			diffSum += diff
		}
	}

	fmt.Println("Answer to Day 1 Part 1: ", diffSum)
}


func partTwo() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	var left, right []int
	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		l, err := strconv.Atoi(line[0]);

		if err != nil {
			fmt.Println("not an integer")
		}

		r, err := strconv.Atoi(line[1]);
		if err != nil {
			fmt.Println("not an integer")
		}

		left = append(left, l)
		right = append(right, r)
	}

	sum := 0
	for i := 0; i < len(left); i++ {
		inner := left[i]
		for j := 0; j < len(right); j++ {
			if inner == right[j] {
				sum += inner
			}
		}
	}

	fmt.Println("Answer to Day 1 Part 2: ", sum)
}
