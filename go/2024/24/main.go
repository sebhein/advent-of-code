package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

type Gate struct {
	w1, w2, op string
}

func compute(w1, w2 bool, op string) bool {
	switch op {
	case "AND":
		return w1 && w2
	case "OR":
		return w1 || w2
	case "XOR":
		return w1 != w2
	default:
		panic("invalid op")
	}
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	allWires := make(map[string]bool)
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ":")
		if len(line) == 1 {
			break
		}
		label := line[0]
		value := strings.TrimSpace(line[1]) == "1"
		allWires[label] = value
	}

	gates := make(map[string]Gate)

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		w1 := line[0]
		op := line[1]
		w2 := line[2]
		out := line[4]

		gates[out] = Gate{w1, w2, op}
	}

	for len(gates) > 0 {
		for out, gate := range gates {
			w1, w1Exists := allWires[gate.w1]
			w2, w2Exists := allWires[gate.w2]
			if w1Exists && w2Exists {
				allWires[out] = compute(w1, w2, gate.op)
				delete(gates, out)
			}
		}
	}

	number := 0
	for idx := 0; idx >= 0; idx++ {
		wire := fmt.Sprintf("z%02d", idx)
		v, exists := allWires[wire]
		if !exists {
			break
		}
		if v {
			number += int(math.Pow(2, float64(idx)))
		}
	}

	fmt.Println("Answer to Day 24 Part 1: ", number)
}
