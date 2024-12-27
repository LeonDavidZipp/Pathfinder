package parsing

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	m "github.com/LeonDavidZipp/Pathfinder/src/models"
)

func createTempFile(content string) (string, error) {
	tmpfile, err := os.CreateTemp("", "test_maze*.txt")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.WriteString(content); err != nil {
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}

var mp *m.Map
var content []byte

func TestMain(m *testing.M) {
	path, err := createTempFile("11111111111\n1S1E1\n10101\n10001\n11111")
	if err != nil {
		panic(err)
	}

	content, err = ReadFile(path)
	if err != nil {
		panic(err)
	}

	mp, err = ParseMap(content)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestParseMap(t *testing.T) {
	rows := bytes.Split(content, []byte("\n"))

	// check if the map was parsed correctly
	for i := range mp.Rows {
		assert.Equal(t, len(mp.Rows[i]), len(rows[i]))
		for j := range mp.Rows[i] {
			switch rows[i][j] {
			case '0':
				assert.Equal(t, mp.Rows[i][j].Type, m.Tile)
			case '1':
				assert.Equal(t, mp.Rows[i][j].Type, m.Wall)
			case 'S':
				assert.Equal(t, mp.Rows[i][j].Type, m.Start)
			case 'E':
				assert.Equal(t, mp.Rows[i][j].Type, m.End)
			}
		}
	}

	// check if pointers where set correctly for first row
	assert.Nil(t, mp.Rows[0][0].Top)
	assert.Nil(t, mp.Rows[0][0].Left)
	for i := 1; i < len(mp.Rows[0])-1; i++ {
		assert.Equal(t, mp.Rows[0][i].Left, mp.Rows[0][i-1])
		assert.Nil(t, mp.Rows[0][i].Top)
	}
	assert.Nil(t, mp.Rows[0][len(mp.Rows[0])-1].Top)
	assert.Nil(t, mp.Rows[0][len(mp.Rows[0])-1].Right)

	// check if pointers where set correctly for middle rows
	for i := 1; i < len(mp.Rows); i++ {
		al := len(mp.Rows[i-1])
		assert.Equal(t, mp.Rows[i][0].Top, mp.Rows[i-1][0])
		assert.Nil(t, mp.Rows[i][0].Left)
		for j := 1; j < len(mp.Rows[i])-1; j++ {
			assert.Equal(t, mp.Rows[i][j].Left, mp.Rows[i][j-1])
			if j < al {
				assert.Equal(t, mp.Rows[i][j].Top, mp.Rows[i-1][j])
			} else {
				assert.Nil(t, mp.Rows[i][j].Top)
			}
		}
		assert.Nil(t, mp.Rows[i][len(mp.Rows[i])-1].Right)
		assert.Equal(t, mp.Rows[i][len(mp.Rows[i])-1].Left, mp.Rows[i][len(mp.Rows[i])-2])
		if len(mp.Rows[i])-1 < al {
			assert.Equal(t, mp.Rows[i][len(mp.Rows[i])-1].Top, mp.Rows[i-1][len(mp.Rows[i])-1])
		} else {
			assert.Nil(t, mp.Rows[i][len(mp.Rows[i])-1].Top)
		}
	}

	// check if pointers where set correctly for last row
	assert.Nil(t, mp.Rows[len(mp.Rows)-1][0].Bottom)
	assert.Nil(t, mp.Rows[len(mp.Rows)-1][0].Left)
	assert.Equal(t, mp.Rows[len(mp.Rows)-1][0].Top, mp.Rows[len(mp.Rows)-2][0])

	al := len(mp.Rows[len(mp.Rows)-2])
	for i := 1; i < len(mp.Rows[len(mp.Rows)-1])-1; i++ {
		assert.Equal(t, mp.Rows[len(mp.Rows)-1][i].Left, mp.Rows[len(mp.Rows)-1][i-1])
		assert.Nil(t, mp.Rows[len(mp.Rows)-1][i].Bottom)
		if i < al {
			assert.Equal(t, mp.Rows[len(mp.Rows)-1][i].Top, mp.Rows[len(mp.Rows)-2][i])
		} else {
			assert.Nil(t, mp.Rows[len(mp.Rows)-1][i].Top)
		}
	}

	// check if start cell was set correctly
	assert.Equal(t, mp.Start, mp.Rows[1][1])
}
