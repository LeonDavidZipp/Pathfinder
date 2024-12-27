package algorithm

import (
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

func TestSolveWrapper(t *testing.T) {
	mp1, err := p.ParseMap(content1)
	assert.NoError(t, err)
	mp2, err := p.ParseMap(content2)
	assert.NoError(t, err)

	sol1, err1 := SolveWrapper(mp1)
	assert.Nil(t, err1)
	assert.NotNil(t, sol1)

	sol2, err2 := SolveWrapper(mp2)
	assert.Nil(t, err2)
	assert.NotNil(t, sol2)
}
