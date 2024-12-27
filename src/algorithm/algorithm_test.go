package algorithm

import (
	m "github.com/LeonDavidZipp/Pathfinder/src/models"
	p "github.com/LeonDavidZipp/Pathfinder/src/parsing"
	"github.com/stretchr/testify/assert"
	"testing"
)

var content1 []byte = []byte(
	`11111111111
1S1E1
10101
10001
11111`,
)

var content2 []byte = []byte(
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
}

func TestMove(t *testing.T) {
	moveMp, err := p.ParseMap([]byte("1111\n1011\n1S01\n1111"))
	assert.NoError(t, err)
	bot := m.NewBot(moveMp.Start)
	assert.Equal(t, moveMp.Start, bot.Pos)

	bot.Move(m.Right)
	assert.Equal(t, m.Left, bot.FromDir)
	assert.Equal(t, moveMp.Rows[2][2], bot.Pos)
	assert.Equal(t, 1, bot.Steps)

	bot.Move(m.Left)
	assert.Equal(t, m.Right, bot.FromDir)
	assert.Equal(t, moveMp.Rows[2][1], bot.Pos)
	assert.Equal(t, 2, bot.Steps)

	bot.Move(m.Up)
	assert.Equal(t, m.Down, bot.FromDir)
	assert.Equal(t, moveMp.Rows[1][1], bot.Pos)
	assert.Equal(t, 3, bot.Steps)

	bot.Move(m.Down)
	assert.Equal(t, m.Up, bot.FromDir)
	assert.Equal(t, moveMp.Rows[2][1], bot.Pos)
	assert.Equal(t, 4, bot.Steps)
}

func TestSolveWrapper(t *testing.T) {
	mp1, err := p.ParseMap(content1)
	assert.NoError(t, err)
	mp2, err := p.ParseMap(content2)
	assert.NoError(t, err)

	sol1, err1 := SolveWrapper(mp1)
	assert.Nil(t, err1)
	assert.NotNil(t, sol1)
	assert.Equal(t, 6, sol1.Steps)

	sol2, err2 := SolveWrapper(mp2)
	assert.Nil(t, err2)
	assert.NotNil(t, sol2)
	assert.Equal(t, 40, sol2.Steps)
}
