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

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	regA, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[2])
	fileScanner.Scan()
	regB, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[2])
	fileScanner.Scan()
	regC, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[2])
	fileScanner.Scan()
	fileScanner.Scan()
	programStrings := strings.Split(fileScanner.Text(), " ")[1]
	programNums := []int{}
	for _, numString := range strings.Split(programStrings, ",") {
		num, _ := strconv.Atoi(numString)
		programNums = append(programNums, num)
	}

	computer := ChronospatitalComputer{
		regA,
		regB,
		regC,
		0,
		programNums,
		"",
		false,
	}
	
	computer.compute()
	fmt.Println("Answer to Day 17 Part 1: ", computer.output)
}

func partTwo(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	// regA, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[2])
	fileScanner.Scan()
	regB, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[2])
	fileScanner.Scan()
	regC, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[2])
	fileScanner.Scan()
	fileScanner.Scan()
	programStrings := strings.Split(fileScanner.Text(), " ")[1]
	programNums := []int{}
	for _, numString := range strings.Split(programStrings, ",") {
		num, _ := strconv.Atoi(numString)
		programNums = append(programNums, num)
	}

	var seed int

	for i := len(programNums) - 1; i >= 0; i-- {
		seed <<= 3
		for {
			fmt.Println("Computing with register set to: ", seed)
			computer := ChronospatitalComputer{
				seed,
				regB,
				regC,
				0,
				programNums,
				"",
				true,
			}
			computer.compute()
			if computer.output != programStrings[i*2:] {
				seed++
			} else {
				break
			}
		}
	}
	fmt.Println("Answer to Day 17 Part 2: ", seed)
}


type ChronospatitalComputer struct {
	regA, regB, regC int
	instructionPointer int
	instructions [] int
	output string
	part2 bool
}

func (cc *ChronospatitalComputer) compute() {
	for cc.instructionPointer < len(cc.instructions){
		if cc.part2 && len(cc.output) > len(cc.instructions) * 2 {
			return
		}
		switch cc.instructions[cc.instructionPointer] {
		case 0:
			cc.adv()
		case 1:
			cc.bxl()
		case 2:
			cc.bst()
		case 3:
			cc.jnz()
		case 4:
			cc.bxc()
		case 5:
			cc.out()
		case 6:
			cc.bdv()
		case 7:
			cc.cdv()
		}
	}
}

func (cc *ChronospatitalComputer) literal() int {
	return cc.instructions[cc.instructionPointer + 1]
}

func (cc *ChronospatitalComputer) combo() int {
	switch cc.instructions[cc.instructionPointer + 1] {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return cc.regA
	case 5:
		return cc.regB
	case 6:
		return cc.regC
	case 7:
		panic("invalid program")
	}
	return -1
}

func (cc *ChronospatitalComputer) adv() {
	cc.regA = cc.regA / int(math.Exp2(float64(cc.combo())))
	cc.instructionPointer += 2
}

func (cc *ChronospatitalComputer) bxl() {
	cc.regB = cc.regB ^ cc.literal()
	cc.instructionPointer += 2
}

func (cc *ChronospatitalComputer) bst() {
	cc.regB = cc.combo() % 8
	cc.instructionPointer += 2
}

func (cc *ChronospatitalComputer) jnz() {
	if cc.regA == 0 {
		cc.instructionPointer += 2
		return
	}
	cc.instructionPointer = cc.literal()
}

func (cc *ChronospatitalComputer) bxc() {
	cc.regB = cc.regB ^ cc.regC
	cc.instructionPointer += 2
}

func (cc *ChronospatitalComputer) out() {
	if len(cc.output) > 0 {
		cc.output = fmt.Sprintf("%s,%d", cc.output, cc.combo() % 8)
	} else {
		cc.output = fmt.Sprintf("%d", cc.combo() % 8)
	}
	cc.instructionPointer += 2
}

func (cc *ChronospatitalComputer) bdv() {
	cc.regB = cc.regA / int(math.Exp2(float64(cc.combo())))
	cc.instructionPointer += 2
}

func (cc *ChronospatitalComputer) cdv() {
	cc.regC = cc.regA / int(math.Exp2(float64(cc.combo())))
	cc.instructionPointer += 2
}
