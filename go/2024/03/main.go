package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		re := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
		// fmt.Println(re.FindAllString(line, -1))
		for _, match := range re.FindAllString(line, -1) {
			// fmt.Println(match)
			splits := strings.Split(match, ",")
			left, err := strconv.Atoi(splits[0][4:])
			if err != nil {
				fmt.Println("not an integer")
			}
			right, err := strconv.Atoi(splits[1][:len(splits[1])-1])
			if err != nil {
				fmt.Println("not an integer")
			}
			// fmt.Println(left, right)
			sum += left * right
		}
	}
	fmt.Println("Answer to Day 3 Part 1: ", sum)
}


func partTwo() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	sum := 0
	enable := true
	for fileScanner.Scan() {
		line := fileScanner.Text()
		re := regexp.MustCompile(`(mul\(\d{1,3}\,\d{1,3}\))|(do\(\))|(don't\(\))`)
		// fmt.Println(re.FindAllString(line, -1))
		for _, match := range re.FindAllString(line, -1) {
			switch match {
			case "do()":
				// fmt.Println("enabling")
				enable = true
			case "don't()":
				// fmt.Println("disabling")
				enable = false
			default:
				// fmt.Println("we are enabled: ", enable)
				if !enable {
					continue
				}
				splits := strings.Split(match, ",")
				left, err := strconv.Atoi(splits[0][4:])
				if err != nil {
					fmt.Println("not an integer")
				}
				right, err := strconv.Atoi(splits[1][:len(splits[1])-1])
				if err != nil {
					fmt.Println("not an integer")
				}
				sum += left * right
			}
		}
	}
	fmt.Println("Answer to Day 3 Part 2: ", sum)
}
