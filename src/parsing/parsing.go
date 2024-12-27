package parsing

import (
	"bytes"
	"fmt"
	"io"
	"os"

	m "github.com/LeonDavidZipp/Pathfinder/src/models"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// assumes map is valid
func ParseMap(content []byte) (*m.Map, error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("empty map")
	}

	mp := m.NewMap()

	// byte rows
	r := bytes.Split(content, []byte("\n"))
	// cell rows
	c := make([][]*m.Cell, len(r))
	for i := range c {
		c[i] = make([]*m.Cell, len(r[i]))
	}

	// read in first overall element of the map
	var cur *m.Cell
	switch r[0][0] {
	case '0':
		cur = m.NewCell(m.Tile)
	case '1':
		cur = m.NewCell(m.Wall)
	case 'S':
		cur = m.NewCell(m.Start)
		mp.Start = cur
	case 'E':
		cur = m.NewCell(m.End)
	}

	c[0][0] = cur

	// read in remaining first row of the map
	for i := 1; i < len(r[0]); i++ {
		switch r[0][i] {
		case '0':
			cur.Right = m.NewCell(m.Tile)
		case '1':
			cur.Right = m.NewCell(m.Wall)
		case 'S':
			cur.Right = m.NewCell(m.Start)
			mp.Start = cur.Right
		case 'E':
			cur.Right = m.NewCell(m.End)
		}

		cur.Right.Left = cur
		cur = cur.Right
		c[0][i] = cur
	}

	// now read in rest; differentiate between first column and remaining columns
	for i := 1; i < len(r); i++ {
		al := len(r[i-1])

		switch r[i][0] {
		case '0':
			cur = m.NewCell(m.Tile)
		case '1':
			cur = m.NewCell(m.Wall)
		case 'S':
			cur = m.NewCell(m.Start)
			mp.Start = cur
		case 'E':
			cur = m.NewCell(m.End)
		}

		cur.Top = c[i-1][0]
		c[i-1][0].Bottom = cur
		c[i][0] = cur

		for j := 1; j < len(r[i]); j++ {
			switch r[i][j] {
			case '0':
				cur.Right = m.NewCell(m.Tile)
			case '1':
				cur.Right = m.NewCell(m.Wall)
			case 'S':
				cur.Right = m.NewCell(m.Start)
				mp.Start = cur.Right
			case 'E':
				cur.Right = m.NewCell(m.End)
			}

			cur.Right.Left = cur
			if j < al {
				cur.Right.Top = c[i-1][j]
				c[i-1][j].Bottom = cur.Right
			}
			c[i][j] = cur.Right
			cur = cur.Right
		}
	}

	mp.Rows = c

	return mp, nil
}