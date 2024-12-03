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
