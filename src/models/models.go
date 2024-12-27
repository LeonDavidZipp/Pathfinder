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
	Up Direction = iota
	Right
	Down
	Left
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

// updates the bot's position and route
func (b *Bot) Move(d Direction) {
	switch d {
	case Up:
		b.Pos = b.Pos.Top
	case Right:
		b.Pos = b.Pos.Right
	case Down:
		b.Pos = b.Pos.Bottom
	case Left:
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
	case Up:
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, Right)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, Down)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, Left)
		}
	case Right:
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, Down)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, Left)
		}
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, Up)
		}
	case Down:
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, Left)
		}
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, Up)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, Right)
		}
	case Left:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, Up)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, Right)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, Down)
		}
	case None:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			dirs = append(dirs, Up)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			dirs = append(dirs, Right)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			dirs = append(dirs, Down)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			dirs = append(dirs, Left)
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
