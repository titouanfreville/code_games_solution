package main

import (
	"fmt"
	"math"
	"os"
)

func debug(val ...interface{}) {
	fmt.Fprintln(os.Stderr, val...)
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
type Bat struct {
	Pos     *Pos
	YDown   int
	XRight  int
	XLeft   int
	YUp     int
	NbRound int
	IsTower bool
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func (b *Bat) Up() {
	if b.IsTower {
		b.YDown = b.Pos.Y - 1
		candidate := b.YUp + (b.Pos.Y-b.YUp)/2
		if b.Pos.Y == candidate {
			b.Pos.Y = candidate + 1
		} else {
			b.Pos.Y = candidate
		}
		return
	}
	b.YDown = b.Pos.Y - 1
	candidate := b.YUp + max((b.YDown-b.YUp)/2, 1)
	if candidate == b.Pos.Y {
		b.Pos.Y--
	} else {
		b.Pos.Y = candidate
	}
}

func (b *Bat) Down() {
	if b.IsTower {
		b.YUp = b.Pos.Y + 1
		candidate := b.Pos.Y + (b.YDown-b.Pos.Y)/2
		if b.Pos.Y == candidate {
			b.Pos.Y = candidate + 1
		} else {
			b.Pos.Y = candidate
		}
		return
	}
	b.YUp = b.Pos.Y + 1
	candidate := b.YDown - max((b.YDown-b.YUp)/2, 1)
	if candidate == b.Pos.Y {
		b.Pos.Y++
	} else {
		b.Pos.Y = candidate
	}
}

func (b *Bat) Right() {
	b.XLeft = b.Pos.X + 1
	candidate := b.XRight - max((b.XRight-b.XLeft)/2, 1)
	if candidate == b.Pos.X {
		b.Pos.X++
	} else {
		b.Pos.X = candidate
	}
}

func (b *Bat) Left() {
	b.XRight = b.Pos.X - 1
	candidate := b.XLeft + max((b.XRight-b.XLeft)/2, 1)
	if candidate == b.Pos.X {
		b.Pos.X--
	} else {
		b.Pos.X = candidate
	}
}

type Pos struct {
	X int
	Y int
}

func main() {
	var (
		W, H, N, X0, Y0 int
	)

	fmt.Scan(&W, &H)
	fmt.Scan(&N)
	fmt.Scan(&X0, &Y0)
	pos := &Pos{X0, Y0}
	bat := &Bat{Pos: pos, YDown: H - 1, XRight: W - 1, IsTower: W == 1}

	for N != 0 {
		N--
		bat.NbRound = N

		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)

		switch bombDir {
		case "U":
			bat.Up()
		case "UR":
			bat.Up()
			bat.Right()
		case "R":
			bat.Right()
		case "DR":
			bat.Down()
			bat.Right()
		case "D":
			bat.Down()
		case "DL":
			bat.Down()
			bat.Left()
		case "L":
			bat.Left()
		case "UL":
			bat.Up()
			bat.Left()
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// the location of the next window Batman should jump to.
		fmt.Println(pos.X, pos.Y)
	}
}
