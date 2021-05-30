package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	maxBytesToScan = 1000000
)

type Map [][]bool

func (m Map) printNeighboors(width, heigh int) {
	for i, line := range m {
		for j, cell := range line {
			if !cell {
				continue
			}

			data := []int{j, i}

			hasRightNeighboor := false
			hasBotNeighboor := false

			if j+1 < width {
				for tmp := j + 1; tmp < width; tmp++ {
					hasRightNeighboor = line[tmp]
					if hasRightNeighboor {
						data = append(data, tmp, i)
						break
					}
				}

			}

			if !hasRightNeighboor {
				data = append(data, -1, -1)
			}

			if i+1 < heigh {
				for tmp := i + 1; tmp < heigh; tmp++ {
					hasBotNeighboor = m[tmp][j]
					if hasBotNeighboor {
						data = append(data, j, tmp)
						break
					}
				}

			}

			if !hasBotNeighboor {
				data = append(data, -1, -1)
			}

			fmt.Println(fmt.Sprintf("%d %d %d %d %d %d", data[0], data[1], data[2], data[3], data[4], data[5]))
		}
	}
}

func isCell(char string) bool {
	return char == "O" || char == "0"
}

func main() {
	var (
		width, height int

		cellMap Map
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, maxBytesToScan), maxBytesToScan)

	// width: the number of cells on the X axis
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	// height: the number of cells on the Y axis
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)

	cellMap = make([][]bool, height)
	for i := 0; i < height; i++ {
		mapLine := make([]bool, width)

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
