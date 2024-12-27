package models

type Type uint8

const (
	Wall Type = iota
	Tile
	Start
	End
)

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Cell struct {
	Top    *Cell
	Right  *Cell
	Bottom *Cell
	Left   *Cell
	Type   Type
}

func NewCell(t Type) *Cell {
	return &Cell{
		Top:    nil,
		Right:  nil,
		Bottom: nil,
		Left:   nil,
		Type:   t,
	}
}

type Bot struct {
	Direction Direction
}

type Map struct {
	Start *Cell
	Rows  [][]*Cell
}

func NewMap() *Map {
	return &Map{
		Start: nil,
		Rows:  nil,
	}
}
