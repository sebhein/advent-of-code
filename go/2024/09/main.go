package main

import (
	"bufio"
	"fmt"
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

	var diskMap []int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()
	fileID := 0
	for idx, char := range strings.Split(line, "") {
		num, _ := strconv.Atoi(char)
		if idx%2 == 0 {
			// this is a file
			for range num {
				diskMap = append(diskMap, fileID)
			}
			fileID += 1
		} else {
			// this is free space
			for range num {
				diskMap = append(diskMap, -1)
			}
		}
	}

	checksum := 0
	backwards := len(diskMap) - 1
	for idx, fid := range diskMap {
		// fmt.Println("idx, fid, backwards", idx, fid, backwards)
		if backwards < idx {
			break
		}
		if fid == -1 {
			for diskMap[backwards] == -1 {
				backwards--
			}
			diskMap[idx] = diskMap[backwards]
			diskMap[backwards] = -1
			backwards--
		}
		// fmt.Println("our map: ", diskMap)
		// fmt.Println("calculating: ", idx, diskMap[idx])
		checksum += idx * diskMap[idx]
	}

	fmt.Println("Answer to Day 9 Part 1: ", checksum)
}

type file struct {
	idx    int
	length int
	fileId int
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var diskMap []file
	var freeSpace []file

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()
	fileID := 0
	diskIdx := 0
	for idx, char := range strings.Split(line, "") {
		num, _ := strconv.Atoi(char)
		if idx%2 == 0 {
			// this is a file
			diskMap = append(diskMap, file{diskIdx, num, fileID})
			fileID += 1
		} else {
			// this is free space
			freeSpace = append(freeSpace, file{diskIdx, num, -1})
		}
		diskIdx += num
	}

	for rev := len(diskMap) - 1; rev >= 0; rev-- {
		file := diskMap[rev]
		for si, space := range freeSpace {
			if space.idx > file.idx {
				break
			}
			if space.length >= file.length {
				diskMap[rev].idx = space.idx
				freeSpace[si].length -= file.length
				freeSpace[si].idx += file.length
				break
			}
		}
	}

	// fmt.Println("hmmmmm: ", diskMap)

	checksum := 0
	for _, f := range diskMap {
		for i := 0; i < f.length; i++ {
			checksum += (f.idx + i) * f.fileId
		}
	}

	fmt.Println("Answer to Day 9 Part 2: ", checksum)
}
