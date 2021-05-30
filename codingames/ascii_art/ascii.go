package main

import (
	"bufio"
	"fmt"
	"os"
)

type Letter string

type Alphabet map[string]Letter

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)

	scanner.Scan()
	T := scanner.Text()
	_ = T // to avoid unused error
	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		_ = ROW // to avoid unused error
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("answer") // Write answer to stdout
}
