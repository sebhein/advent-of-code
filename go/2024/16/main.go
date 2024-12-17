package main

import (
	"bufio"
	"container/heap"
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

func partOne(inputFile string) {
	readFile, _ := os.Open(inputFile)
	defer readFile.Close()

	var maze [][]bool
	var start, end Position
	fileScanner := bufio.NewScanner(readFile)
	for rowIdx := 0; fileScanner.Scan(); rowIdx++ {
		line := fileScanner.Text()
		var row []bool
		for col, char := range strings.Split(line, "") {
			traversable := true
			if char == "#" {
				traversable = false
			}
			if char == "E" {
				end = Position{rowIdx, col}
			}
			if char == "S" {
				start = Position{rowIdx, col}
			}
			row = append(row, traversable)
		}
		maze = append(maze, row)
	}

	fmt.Println(start, end)
	for ri, row := range maze {
		for ci, traverse := range row {
			toPrint := "."
			if !traverse {
				toPrint = "#"
			}
			if ri == start.row && ci == start.col {
				toPrint = "S"
			}
			if ri == end.row && ci == end.col {
				toPrint = "E"
			}
			fmt.Print(toPrint)
		}
		fmt.Print("\n")
	}

	path := aStar(maze, start, end)
	fmt.Println(path)
}

// Position represents a point in 2D space
type Position struct {
	row, col int
}

// Node represents a node in the pathfinding graph
type Node struct {
	pos    Position
	facing string
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

func aStar(maze [][]bool, start, end Position) float64 {
	height := len(maze)
	width := len(maze[0])
	openSet := &PriorityQueue{}
	heap.Init(openSet)

	startNode := &Node{
		pos:    start,
		facing: "east",
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
			return current.gScore
		}

		var currentDir Position
		switch current.facing {
		case "north":
			currentDir = Position{-1, 0}
		case "east":
			currentDir = Position{0, 1}
		case "south":
			currentDir = Position{1, 0}
		case "west":
			currentDir = Position{0, -1}
		}

		// Check all neighbors
		for _, dir := range []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neighbor := Position{
				row: current.pos.row + dir.row,
				col: current.pos.col + dir.col,
			}

			if !isValid(neighbor, height, width) || !maze[neighbor.row][neighbor.col] {
				continue
			}

			cost := 1
			if dir != currentDir {
				turns := int(math.Abs(float64(currentDir.row-dir.row))) + int(math.Abs(float64(currentDir.col-dir.col))) - 1
				cost += turns * 1000
			}
			tentativeGScore := gScore[current.pos] + float64(cost)

			var facing string
			switch dir {
			case Position{-1, 0}:
				facing = "north"
			case Position{0, 1}:
				facing = "east"
			case Position{1, 0}:
				facing = "south"
			case Position{0, -1}:
				facing = "west"
			}

			if existingScore, exists := gScore[neighbor]; !exists || tentativeGScore < existingScore {
				neighborNode := &Node{
					pos:    neighbor,
					facing: facing,
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

	return 0 // No path found
}

// isValid checks if a position is within the grid bounds
func isValid(pos Position, height, width int) bool {
	return pos.row >= 0 && pos.row < height && pos.col >= 0 && pos.col < width
}
