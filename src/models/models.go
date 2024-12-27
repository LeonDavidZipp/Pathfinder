package models

type Type uint8

const (
	Wall Type = iota
	Tile
	Start
	End
	Space
)

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
	None
)

type Cell struct {
	Top    *Cell
	Right  *Cell
	Bottom *Cell
	Left   *Cell
	Type   Type
}

// creates a new cell with the given type
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
	Dir   Direction
	Pos   *Cell
	Route []Direction
	Steps uint64
}

// creates a new bot with the given start cell
func NewBot(start *Cell) *Bot {
	return &Bot{
		Dir:   None,
		Pos:   start,
		Route: make([]Direction, 0),
		Steps: 0,
	}
}

// creates a deep copy of the bot
func CopyBot(src *Bot) Bot {
	dst := Bot{
		Dir:   src.Dir,
		Steps: src.Steps,
	}

	s := *(src.Pos)
	dst.Pos = &s
	dst.Route = append(dst.Route, src.Route...)

	return dst
}

// func CopyBot(src *Bot) Bot {
// 	dst := Bot{
// 		FromDir: src.FromDir,
// 		Pos:     src.Pos,
// 		Steps:   src.Steps,
// 		Route:   src.Route,
// 	}

// 	return dst
// }

// Northdates the bot's position and route
func (b *Bot) Move(d Direction) {
	switch d {
	case North:
		b.Pos = b.Pos.Top
	case East:
		b.Pos = b.Pos.Right
	case South:
		b.Pos = b.Pos.Bottom
	case West:
		b.Pos = b.Pos.Left
	}
	b.Dir = d
	b.Route = append(b.Route, d)
	b.Steps++
}

// counts & returns all possible new paths from the current position
func (b *Bot) CountPaths() []Direction {
	dirs := make([]Direction, 0)
	switch b.Dir {
	case North:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, North)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, East)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, West)
		}
	case East:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, North)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, East)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, South)
		}
	case South:
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, East)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, South)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, West)
		}
	case West:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, North)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, West)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, South)
		}
	case None:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, North)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, East)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, South)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, West)
		}
	}

	return dirs
}

type Map struct {
	Start *Cell
	Rows  [][]*Cell
}

// generates a new map from the given start cell and rows
func NewMap(start *Cell, rows [][]*Cell) *Map {
	return &Map{
		Start: start,
		Rows:  rows,
	}
}

type Solution struct {
	Steps uint64
	Route []Direction
}
