package algorithm

import (
	"context"
	"os"
	"testing"
	"time"

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

var content7 []byte = []byte(
	`11111111
111000E1
1S001111
11100001
11111111`,
)

var content8 []byte = []byte(
	`111111111111111111
1S0000000010000001
101110101010101101
101110111010111101
101000101010111111
101011101110000001
101000000001111111
101111110100000001
101E00000001010101
111111111111111111`,
)

var (
	mp1 *m.Map
	mp2 *m.Map
	mp3 *m.Map
	mp4 *m.Map
	mp5 *m.Map
	mp6 *m.Map
	mp7 *m.Map
	mp8 *m.Map
	ctx context.Context
)

func TestMain(m *testing.M) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

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

	mp7, err = p.ParseMap(content7)
	if err != nil {
		panic(err)
	}

	mp8, err = p.ParseMap(content8)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCountPaths(t *testing.T) {
	mp, err := p.ParseMap([]byte("11111\n11011\n10S01\n11011\n11111"))
	assert.NoError(t, err)
	bot := m.NewBot(mp.Start)

	bot.Dir = m.None
	paths := bot.CountPaths()
	assert.Len(t, paths, 4)

	bot.Dir = m.North
	paths = bot.CountPaths()
	assert.Len(t, paths, 3)

	bot.Dir = m.East
	paths = bot.CountPaths()
	assert.Len(t, paths, 3)

	bot.Dir = m.South
	paths = bot.CountPaths()
	assert.Len(t, paths, 3)

	bot.Dir = m.West
	paths = bot.CountPaths()
	assert.Len(t, paths, 3)

	bot.Move(m.East)
	assert.Equal(t, m.East, bot.Dir)
	paths = bot.CountPaths()
	assert.Len(t, paths, 0)

	bot.Dir = m.West
	paths = bot.CountPaths()
	assert.Len(t, paths, 1)

	bot.Move(m.West)
	paths = bot.CountPaths()
	assert.Len(t, paths, 3)
	assert.Equal(t, m.West, bot.Dir)
}

func TestMove(t *testing.T) {
	moveMp, err := p.ParseMap([]byte("1111\n1011\n1S01\n1111"))
	assert.NoError(t, err)
	bot := m.NewBot(moveMp.Start)
	assert.Equal(t, moveMp.Start, bot.Pos)

	bot.Move(m.East)
	assert.Equal(t, m.East, bot.Dir)
	assert.Equal(t, moveMp.Rows[2][2], bot.Pos)
	assert.Equal(t, uint64(1), bot.Steps)

	bot.Move(m.West)
	assert.Equal(t, m.West, bot.Dir)
	assert.Equal(t, moveMp.Rows[2][1], bot.Pos)
	assert.Equal(t, uint64(2), bot.Steps)

	bot.Move(m.North)
	assert.Equal(t, m.North, bot.Dir)
	assert.Equal(t, moveMp.Rows[1][1], bot.Pos)
	assert.Equal(t, uint64(3), bot.Steps)

	bot.Move(m.South)
	assert.Equal(t, m.South, bot.Dir)
	assert.Equal(t, moveMp.Rows[2][1], bot.Pos)
	assert.Equal(t, uint64(4), bot.Steps)
}

func TestSolveWrapper(t *testing.T) {
	sol1, err := SolveWrapper(ctx, mp1)
	assert.Nil(t, err)
	assert.NotNil(t, sol1)
	assert.Equal(t, uint64(1), sol1.Steps)

	sol2, err := SolveWrapper(ctx, mp2)
	assert.Nil(t, err)
	assert.NotNil(t, sol2)
	assert.Equal(t, uint64(2), sol2.Steps)

	sol3, err := SolveWrapper(ctx, mp3)
	assert.Nil(t, err)
	assert.NotNil(t, sol3)
	assert.Equal(t, uint64(3), sol3.Steps)

	sol4, err := SolveWrapper(ctx, mp4)
	assert.Nil(t, err)
	assert.NotNil(t, sol4)
	assert.Equal(t, uint64(10), sol4.Steps)

	sol5, err := SolveWrapper(ctx, mp5)
	assert.Nil(t, err)
	assert.NotNil(t, sol5)
	assert.Equal(t, uint64(6), sol5.Steps)

	sol6, err := SolveWrapper(ctx, mp6)
	assert.Nil(t, err)
	assert.NotNil(t, sol6)
	assert.Equal(t, uint64(40), sol6.Steps)

	sol7, err := SolveWrapper(ctx, mp7)
	assert.Nil(t, err)
	assert.NotNil(t, sol7)
	assert.Equal(t, uint64(6), sol7.Steps)

	sol8, err := SolveWrapper(ctx, mp8)
	assert.Nil(t, err)
	assert.NotNil(t, sol8)
	assert.Equal(t, uint64(23), sol8.Steps)
}
