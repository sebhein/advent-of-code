package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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

	network := make(map[string][]string)
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), "-")
		network[line[0]] = append(network[line[0]], line[1])
		network[line[1]] = append(network[line[1]], line[0])
	}

	var triads []string
	for node, connections := range network {
		for _, c := range connections {
			for _, nc := range network[c] {
				if slices.Contains(connections, nc) && slices.Contains(network[nc], node) && slices.Contains(network[nc], c) {
					if nc[0] != 't' && node[0] != 't' && c[0] != 't' {
						continue
					}
					triad := []string{node, c, nc}
					slices.Sort(triad)
					t := strings.Join(triad, "-")
					if !slices.Contains(triads, t) {
						triads = append(triads, t)
					}
				}
			}
		}
	}

	fmt.Println("Answer to Day 23 Part 1: ", len(triads))
	// for _, t := range triads {
	// 	fmt.Println(t)
	// }
}
