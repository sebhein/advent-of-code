package main

import (
	"bufio"
	"fmt"
	"os"
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

type Coord struct {
	x int
	y int
}

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	antennas := make(map[string][]Coord)

	fileScanner := bufio.NewScanner(readFile)
	height := 0
	width := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for col, char := range strings.Split(line, "") {
			if char != "." {
				antennas[char] = append(antennas[char], Coord{height, col})
			}
			if height == 0 {
				width += 1
			}
		}
		height += 1
	}

	resultChan := make(chan []Coord)
	doneChan := make(chan int)
	var wg sync.WaitGroup

	go countAntinodes(resultChan, doneChan)

	for _, coords := range antennas {
		wg.Add(1)
		go findAntinodes(coords, height, width, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	fmt.Println("Answer to Day 8 Part 1: ", <-doneChan)
}

func countAntinodes(resultChan <-chan []Coord, done chan<- int) {
	foundAntinodes := make(map[Coord]int)
	for result := range resultChan {
		for _, antinode := range result {
			foundAntinodes[antinode] += 1
		}
	}
	done <- len(foundAntinodes)
}

func inBounds(loc Coord, height, width int) bool {
	if loc.x < 0 || loc.x >= height {
		return false
	}
	if loc.y < 0 || loc.y >= width {
		return false
	}
	return true
}

func findAntinodes(locations []Coord, height, width int, resultChan chan<- []Coord, wg *sync.WaitGroup) {
	defer wg.Done()
	var antinodes []Coord
	for first := 0; first < len(locations)-1; first++ {
		for second := first + 1; second < len(locations); second++ {
			a := locations[first]
			b := locations[second]
			dx := a.x - b.x
			dy := a.y - b.y

			if inBounds(Coord{a.x + dx, a.y + dy}, height, width) {
				antinodes = append(antinodes, Coord{a.x + dx, a.y + dy})
			}
			if inBounds(Coord{b.x - dx, b.y - dy}, height, width) {
				antinodes = append(antinodes, Coord{b.x - dx, b.y - dy})
			}
		}
	}
	resultChan <- antinodes
}

func findResonantAntinodes(locations []Coord, height, width int) []Coord {
	var antinodes []Coord
	if len(locations) <= 1 {
		return []Coord{}
	}
	for first := 0; first < len(locations)-1; first++ {
		for second := first + 1; second < len(locations); second++ {
			a := locations[first]
			b := locations[second]
			dx := a.x - b.x
			dy := a.y - b.y

			dist := 0
			for {
				if inBounds(Coord{a.x + dx*dist, a.y + dy*dist}, height, width) {
					antinodes = append(antinodes, Coord{a.x + dx*dist, a.y + dy*dist})
				} else {
					break
				}
				dist += 1
			}

			dist = 0
			for {
				if inBounds(Coord{b.x - dx*dist, b.y - dy*dist}, height, width) {
					antinodes = append(antinodes, Coord{b.x - dx*dist, b.y - dy*dist})
				} else {
					break
				}
				dist += 1
			}
		}
	}
	return antinodes
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	antennas := make(map[string][]Coord)

	fileScanner := bufio.NewScanner(readFile)
	height := 0
	width := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for col, char := range strings.Split(line, "") {
			if char != "." {
				antennas[char] = append(antennas[char], Coord{height, col})
			}
			if height == 0 {
				width += 1
			}
		}
		height += 1
	}

	foundAntinodes := make(map[Coord]int)
	for _, coords := range antennas {
		found := findResonantAntinodes(coords, height, width)
		for _, antinode := range found {
			foundAntinodes[antinode] += 1
		}
	}

	fmt.Println("Answer to Day 8 Part 2: ", len(foundAntinodes))
}
