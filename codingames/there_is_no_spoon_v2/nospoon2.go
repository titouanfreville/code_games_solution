package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	maxBytesToScan = 1000000
)

type Point struct {
	X int
	Y int
}

type Node struct {
	Weight   int
	SureLink map[Point]int
	Link     map[Point]int
}

type Graph [][]Node

func (g *Graph) AddNode(i, j, width, height int, weightStr string) *Graph {
	weight, _ := strconv.Atoi(weightStr)

	return g
}

func main() {
	var (
		width, height int

		cellMap Graph
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, maxBytesToScan), maxBytesToScan)

	// width: the number of cells on the X axis
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	// height: the number of cells on the Y axis
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)

	cellMap = make([][]Node, height)
	for i := 0; i < height; i++ {
		mapLine := make([]Node, width)

		scanner.Scan()
		line := scanner.Text()
		_ = line // to avoid unused error // width characters, each either 0 or .

		for j, cell := range line {
			mapLine[j] = isCell(string(cell))
		}

		cellMap[i] = mapLine
	}

	// Three coordinates: a node, its right neighbor, its bottom neighbor
	cellMap.printNeighboors(width, height)
}
