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

// returns the inverse direction
func (d Direction) Inverse() Direction {
	switch d {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	default:
		return None
	}
}

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
	FromDir Direction
	Pos     *Cell
	Route   []*Cell
	Steps   uint64
}

// creates a new bot with the given start cell
func NewBot(start *Cell) *Bot {
	return &Bot{
		FromDir: None,
		Pos:     start,
		Route:   nil,
		Steps:   0,
	}
}

// creates a shallow copy of the bot
func CopyBot(bot *Bot) Bot {
	return Bot{
		FromDir: bot.FromDir,
		Pos:     bot.Pos,
		Route:   bot.Route,
		Steps:   bot.Steps,
	}
}

// updates the bot's position and route
func (b *Bot) Move(d Direction) {
	switch d == b.FromDir.Inverse() {
	// move straight
	case true:
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
		b.Route = append(b.Route, b.Pos)
	case false:
		switch d {
		case Up:
			b.Pos = b.Pos.Top
			b.FromDir = d.Inverse()
		case Right:
			b.Pos = b.Pos.Right
			b.FromDir = d.Inverse()
		case Down:
			b.Pos = b.Pos.Bottom
			b.FromDir = d.Inverse()
		case Left:
			b.Pos = b.Pos.Left
			b.FromDir = d.Inverse()
		}
		b.Route = append(b.Route, b.Pos)
	}
	b.Steps++
}

// counts & returns all possible new paths from the current position
func (b *Bot) CountPaths() (uint8, []Direction) {
	paths := uint8(0)
	dirs := make([]Direction, 0)
	switch b.FromDir {
	case Up:
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			paths++
			dirs = append(dirs, Right)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			paths++
			dirs = append(dirs, Down)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			paths++
			dirs = append(dirs, Left)
		}
	case Right:
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			paths++
			dirs = append(dirs, Down)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			paths++
			dirs = append(dirs, Left)
		}
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			paths++
			dirs = append(dirs, Up)
		}
	case Down:
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			paths++
			dirs = append(dirs, Left)
		}
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			paths++
			dirs = append(dirs, Up)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			paths++
			dirs = append(dirs, Right)
		}
	case Left:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			paths++
			dirs = append(dirs, Up)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			paths++
			dirs = append(dirs, Right)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			paths++
			dirs = append(dirs, Down)
		}
	case None:
		if b.Pos.Top != nil && b.Pos.Top.Type != Wall {
			paths++
			dirs = append(dirs, Up)
		}
		if b.Pos.Right != nil && b.Pos.Right.Type != Wall {
			paths++
			dirs = append(dirs, Right)
		}
		if b.Pos.Bottom != nil && b.Pos.Bottom.Type != Wall {
			paths++
			dirs = append(dirs, Down)
		}
		if b.Pos.Left != nil && b.Pos.Left.Type != Wall {
			paths++
			dirs = append(dirs, Left)
		}
	}

	return paths, dirs
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
	Route []*Cell
}
