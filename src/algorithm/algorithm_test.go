package algorithm

import (
	"os"
	"testing"

	m "github.com/LeonDavidZipp/Pathfinder/src/models"
	p "github.com/LeonDavidZipp/Pathfinder/src/parsing"
	"github.com/stretchr/testify/assert"
)

var content1 []byte = []byte(
	`1111
1SE1
1111`,
)

var content2 []byte = []byte(
	`111
1S1
101
1E1
111`,
)

var content3 []byte = []byte(
	`111
1E1
101
101
1S1
111`,
)

var content4 []byte = []byte(
	`11111111111
1S000000001
11111111101
        1E1
        111`,
)

var content5 []byte = []byte(
	`11111111111
1S1E1
10101
10001
11111`,
)

var content6 []byte = []byte(
	`11111
1S1E1 11111
10101 10001
10101 10101
101011101011111
101000001000001
101111111111101
100000000000001
111111111111111`,
)

var mp1 *m.Map
var mp2 *m.Map
var mp3 *m.Map
var mp4 *m.Map
var mp5 *m.Map
var mp6 *m.Map

func TestMain(m *testing.M) {
	var err error
	mp1, err = p.ParseMap(content1)
	if err != nil {
		panic(err)
	}

	mp2, err = p.ParseMap(content2)
	if err != nil {
		panic(err)
	}

	mp3, err = p.ParseMap(content3)
	if err != nil {
		panic(err)
	}

	mp4, err = p.ParseMap(content4)
	if err != nil {
		panic(err)
	}

	mp5, err = p.ParseMap(content5)
	if err != nil {
		panic(err)
	}

	mp6, err = p.ParseMap(content6)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestMove(t *testing.T) {
	moveMp, err := p.ParseMap([]byte("1111\n1011\n1S01\n1111"))
	assert.NoError(t, err)
	bot := m.NewBot(moveMp.Start)
	assert.Equal(t, moveMp.Start, bot.Pos)

	bot.Move(m.Right)
	assert.Equal(t, m.Left, bot.FromDir)
	assert.Equal(t, moveMp.Rows[2][2], bot.Pos)
	assert.Equal(t, uint64(1), bot.Steps)

	bot.Move(m.Left)
	assert.Equal(t, m.Right, bot.FromDir)
	assert.Equal(t, moveMp.Rows[2][1], bot.Pos)
	assert.Equal(t, uint64(2), bot.Steps)

	bot.Move(m.Up)
	assert.Equal(t, m.Down, bot.FromDir)
	assert.Equal(t, moveMp.Rows[1][1], bot.Pos)
	assert.Equal(t, uint64(3), bot.Steps)

	bot.Move(m.Down)
	assert.Equal(t, m.Up, bot.FromDir)
	assert.Equal(t, moveMp.Rows[2][1], bot.Pos)
	assert.Equal(t, uint64(4), bot.Steps)
}

func TestSolveWrapper(t *testing.T) {
	sol1, err := SolveWrapper(mp1)
	assert.Nil(t, err)
	assert.NotNil(t, sol1)
	assert.Equal(t, 1, sol1.Steps)

	sol2, err := SolveWrapper(mp2)
	assert.Nil(t, err)
	assert.NotNil(t, sol2)
	assert.Equal(t, 2, sol2.Steps)

	sol3, err := SolveWrapper(mp3)
	assert.Nil(t, err)
	assert.NotNil(t, sol3)
	assert.Equal(t, 3, sol3.Steps)

	sol4, err := SolveWrapper(mp4)
	assert.Nil(t, err)
	assert.NotNil(t, sol4)
	assert.Equal(t, 10, sol4.Steps)

	sol5, err := SolveWrapper(mp5)
	assert.Nil(t, err)
	assert.NotNil(t, sol5)
	assert.Equal(t, 6, sol5.Steps)
}
