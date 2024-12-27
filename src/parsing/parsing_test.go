package parsing

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	m "github.com/LeonDavidZipp/Pathfinder/src/models"
)

func createTempFile() (string, error) {
	tmpfile, err := os.CreateTemp("", "test_maze*.txt")
	if err != nil {
		return "", err
	}

		content := "11111111111\n1S1E1\n10101\n10001\n11111"
	if _, err := tmpfile.WriteString(content); err != nil {
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}

func TestParsing(t *testing.T) {
	path, err := createTempFile()
	assert.NoError(t, err)

	content, err := ReadFile(path)
	assert.NoError(t, err)

	mp, err := ParseMap(content)
	assert.NoError(t, err)

	rows := bytes.Split(content, []byte("\n"))
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
}
