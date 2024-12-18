package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFile := os.Args[1]
	size, _ := strconv.Atoi(os.Args[2])
	numFallen, _ := strconv.Atoi(os.Args[3])

	start := time.Now()
	partOne(inputFile, size, numFallen)
	fmt.Println("Solved Part 1 in: ", time.Since(start))
	start = time.Now()
	partTwo(inputFile, size, numFallen)
	fmt.Println("Solved Part 2 in: ", time.Since(start))
}

func partOne(inputFile string, size, numFallen int) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}

	fileScanner := bufio.NewScanner(readFile)
	for bytesFallen := 0; fileScanner.Scan(); bytesFallen++ {
		if bytesFallen == numFallen {
			break
		}
		line := strings.Split(fileScanner.Text(), ",")
		col, _ := strconv.Atoi(line[0])
		row, _ := strconv.Atoi(line[1])
		grid[row][col] = true
	}
	start := Position{0, 0}
	end := Position{size - 1, size - 1}


	finalNode := aStar(grid, start, end)
	current := finalNode
	var path []Position
	for current != nil {
		path = append(path, current.pos)
		current = current.parent
	}

	for ri, row := range grid {
		for ci, obstacle := range row {
			toPrint := ". "
			if obstacle {
				toPrint = "# "
			}
			if slices.Contains(path, Position{ri, ci}) {
				toPrint = "O "
			}
			if ri == start.row && ci == start.col {
				toPrint = "S "
			}
			if ri == end.row && ci == end.col {
				toPrint = "E "
			}
			fmt.Print(toPrint)
		}
		fmt.Print("\n")
	}
	fmt.Println("Answer to Day 18 Part 1: ", len(path) - 1)
}


func partTwo(inputFile string, size, numFallen int) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}

	start := Position{0, 0}
	end := Position{size - 1, size - 1}

	fileScanner := bufio.NewScanner(readFile)
	for bytesFallen := 0; fileScanner.Scan(); bytesFallen++ {
		line := strings.Split(fileScanner.Text(), ",")
		col, _ := strconv.Atoi(line[0])
		row, _ := strconv.Atoi(line[1])
		grid[row][col] = true
		if bytesFallen > numFallen {
			finalNode := aStar(grid, start, end)
			if finalNode == nil {
				fmt.Println("Answer to Day 18 Part 2: ", line)
				return
			}
		}
	}
}

// Position represents a point in 2D space
type Position struct {
	row, col int
}

// Node represents a node in the pathfinding graph
type Node struct {
	pos    Position
	gScore float64 // Cost from start to current node
	fScore float64 // Estimated total cost (gScore + heuristic)
	parent *Node
	index  int // Required for heap interface
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].fScore < pq[j].fScore }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

func manhattanDistance(a, b Position) float64 {
	return float64(math.Abs(float64(a.row-b.row)) + math.Abs(float64(a.col-b.col)))
}

func aStar(grid [][]bool, start, end Position) *Node {
	height := len(grid)
	width := len(grid[0])
	openSet := &PriorityQueue{}
	heap.Init(openSet)

	startNode := &Node{
		pos:    start,
		gScore: 0,
		fScore: manhattanDistance(start, end),
	}
	heap.Push(openSet, startNode)

	cameFrom := make(map[Position]*Node)
	gScore := make(map[Position]float64)
	gScore[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)

		if current.pos == end {
			return current
		}

		// Check all neighbors
		for _, dir := range []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neighbor := Position{
				row: current.pos.row + dir.row,
				col: current.pos.col + dir.col,
			}

			if !isValid(neighbor, height, width) || grid[neighbor.row][neighbor.col] {
				continue
			}

			tentativeGScore := gScore[current.pos] + 1

			if existingScore, exists := gScore[neighbor]; !exists || tentativeGScore < existingScore {
				neighborNode := &Node{
					pos:    neighbor,
					gScore: tentativeGScore,
					fScore: tentativeGScore + manhattanDistance(neighbor, end),
					parent: current,
				}
				gScore[neighbor] = tentativeGScore
				cameFrom[neighbor] = current

				heap.Push(openSet, neighborNode)
			}
		}
	}

	return nil // No path found
}

// isValid checks if a position is within the grid bounds
func isValid(pos Position, height, width int) bool {
	return pos.row >= 0 && pos.row < height && pos.col >= 0 && pos.col < width
}
