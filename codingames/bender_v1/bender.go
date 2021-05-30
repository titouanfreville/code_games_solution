package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func debug(vals ...interface{}) {
	fmt.Fprintln(os.Stderr, vals...)
}

var TeleportersPos []Point

// Direction values
type Direction string

type Directions []Direction

const (
	North      Direction = "N"
	South      Direction = "S"
	West       Direction = "W"
	East       Direction = "E"
	Teleporter Direction = "T"
)

func (directions Directions) String() string {
	res := ""
	for _, dir := range directions {
		res += dir.Long() + "\n"
	}
	return strings.TrimSuffix(res, "\n")
}

func reverse(l []Direction) []Direction {
	base := append([]Direction{}, reverse(l[1:])...)
	base = append(base, l[0])
	return base
}

// Map elements other than Direction
type Elem string

const (
	Start    Elem = "@"
	Border   Elem = "#"
	Obstacle Elem = "X"
	Inverter Elem = "I"
	Beers    Elem = "B"
	End      Elem = "$"
	Empty    Elem = " "
)

func (e Elem) IsDirection() bool {
	return Direction(e).Long() != ""
}

func (d Direction) Long() string {
	switch d {
	case North:
		return "NORTH"
	case South:
		return "SOUTH"
	case West:
		return "WEST"
	case East:
		return "EAST"
	case Teleporter:
		return "TP"
	}
	return ""
}

// Blender instance
type Bender struct {
	Direction           Direction
	CurrentCell         Cell
	Moves               Directions
	Positions           []Point
	Looping             bool
	DirectionPriorities []Direction
	breakingMode        bool
	destroyed           bool
	stopped             bool
}

func NewBender() *Bender {
	return &Bender{
		Direction:           South,
		Looping:             false,
		DirectionPriorities: []Direction{South, East, North, West},
		breakingMode:        false,
	}
}

func (b *Bender) StartAt(cell Cell, point Point) *Bender {
	b.Positions = []Point{point}
	b.CurrentCell = cell
	return b
}

func (b *Bender) CanMoveTo(cell Cell) bool {
	switch cell.Value {
	case Border:
		return false
	case Obstacle:
		return b.breakingMode
	default:
		return true
	}
}

func (b *Bender) selectMove() *Bender {
	evaluateCandidate := func(dir Direction) (*Bender, bool) {
		if nextCandidate, ok := b.CurrentCell.NextCells[dir]; ok && b.CanMoveTo(nextCandidate) {
			debug("moving to next")
			b.CurrentCell = nextCandidate
			b.Moves = append(b.Moves, dir)
			b.Positions = append(b.Positions, b.CurrentCell.Pos)
			debug("New values", fmt.Sprintf("%#v", b))
			return b, true
		}

		return b, false
	}

	if candidate, ok := evaluateCandidate(b.Direction); ok {
		return candidate
	}

	for _, direction := range b.DirectionPriorities {
		if candidate, ok := evaluateCandidate(direction); ok {
			return candidate
		}
	}

	b.stopped = true
	return b
}

func (b *Bender) afterMove() *Bender {
	switch b.CurrentCell.Value {
	case End:
		debug("Got end ?")
		b.destroyed = true
	case Beers:
		b.breakingMode = !b.breakingMode
	case Inverter:
		b.DirectionPriorities = reverse(b.DirectionPriorities)
	}

	switch Direction(b.CurrentCell.Value) {
	case North, South, West, East:
		b.Direction = Direction(b.CurrentCell.Value)
	case Teleporter:
		b.CurrentCell = b.CurrentCell.NextCells[Teleporter]
		b.Positions = append(b.Positions, b.CurrentCell.Pos)
	}

	return b
}

func (b *Bender) looping() bool {
	return false
}

func (b *Bender) Move() *Bender {
	b.selectMove()
	b.afterMove()
	debug(b.destroyed)
	if !b.destroyed && !b.stopped {
		return b.Move()
	}
	return b
}

// Map management
type Cell struct {
	Pos       Point
	Value     Elem
	NextCells map[Direction]Cell
}

func NewCell(val string, pos Point) Cell {
	debug("Add cell", val, pos)
	var value Elem
	switch val {
	case string(Start):
		value = Start
	case string(Border):
		value = Border
	case string(Obstacle):
		value = Obstacle
	case string(Inverter):
		value = Inverter
	case string(Beers):
		value = Beers
	case string(North):
		value = Elem(North)
	case string(South):
		value = Elem(South)
	case string(West):
		value = Elem(West)
	case string(East):
		value = Elem(East)
	case string(Teleporter):
		value = Elem(Teleporter)
	case string(End):
		value = End
	default:
		value = Empty
	}
	return Cell{Pos: pos, Value: value}
}

func (c Cell) IsBorder() bool {
	return c.Value != Border
}

func (c Cell) IsTP() bool {
	return Direction(c.Value) == Teleporter
}

type Point struct {
	X int
	Y int
}

type Map map[Point]Cell

func NewMap() Map {
	return make(Map)
}

func (m Map) AddCell(pos Point, cell Cell) Map {
	debug("Adding cell", fmt.Sprintf("%#v, %#v", pos, cell))
	m[pos] = cell

	if cell.IsTP() {
		debug("Cell is TP")
		if len(TeleportersPos) > 0 {
			cell.NextCells[Teleporter] = m[TeleportersPos[0]]
			m[TeleportersPos[0]].NextCells[Teleporter] = cell
		}

		TeleportersPos = append(TeleportersPos, pos)
	}

	if northCell, ok := m[Point{X: pos.X - 1, Y: pos.Y}]; ok && !northCell.IsBorder() {
		debug("Adding cell to north")
		cell.NextCells[North] = northCell
	}
	if eastCell, ok := m[Point{X: pos.X, Y: pos.Y + 1}]; ok && !eastCell.IsBorder() {
		cell.NextCells[East] = eastCell
		debug("Adding cell to east")
	}
	if southCell, ok := m[Point{X: pos.X + 1, Y: pos.Y}]; ok && !southCell.IsBorder() {
		cell.NextCells[North] = southCell
		debug("Adding cell to south")
	}
	if westCell, ok := m[Point{X: pos.X, Y: pos.Y - 1}]; ok && !westCell.IsBorder() {
		cell.NextCells[North] = westCell
		debug("Adding cell to west")
	}

	return m
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	fmt.Fprintln(os.Stderr, "Debug messages...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L, C int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L, &C)

	benderMap := NewMap()
	bender := NewBender()

	for i := 0; i < L; i++ {
		scanner.Scan()
		row := scanner.Text()
		for j := 0; j < C; j++ {
			pos := Point{X: i, Y: j}
			cell := NewCell(string(row[j]), pos)
			if cell.Value == Start {
				debug("Adding start", fmt.Sprintf("%#v, %#v", pos, cell))
				bender = bender.StartAt(cell, pos)
			}
			benderMap[pos] = cell
		}
	}

	debug(fmt.Sprintf("%#v", benderMap))
	debug(fmt.Sprintf("%#v", bender))
	bender.Move()

	// if !bender.Looping {
	// 	fmt.Println(bender.Moves.String()) // Write answer to stdout
	// } else {
	// 	fmt.Println("LOOP")
	// }
}
