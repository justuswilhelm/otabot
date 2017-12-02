package lib

import (
	"testify/assert"
	"testing"
)

func makeState() *GoState {
	return NewGoState(3, 0.5)
}

func makeMove(s *GoState, p int) *GoMove {
	m := s.newMove(p, false)
	return m.(*GoMove)
}

func TestFromBoard(t *testing.T) {
	s := fromString([]string{
		"OX",
		"..",
	})
	assert.Equal(t, W, s.Get(0))
	assert.Equal(t, B, s.Get(1))
	assert.Equal(t, E, s.Get(2))
	assert.Equal(t, E, s.Get(3))
}

func TestRowCol(t *testing.T) {
	s := makeState()
	r, c := s.RowCol(0)
	assert.Equal(t, 0, r)
	assert.Equal(t, 0, c)

	r, c = s.RowCol(1)
	assert.Equal(t, 0, r)
	assert.Equal(t, 1, c)
}

func TestSetMove(t *testing.T) {
	p := 4
	s := makeState()
	assert.True(t, s.Empty(p))
	m := makeMove(s, p)
	assert.True(t, s.Empty(p))
	s.MakeMove(m)
	assert.False(t, s.Empty(p))
	assert.Equal(t, B, s.Get(p))
}

func TestIndex(t *testing.T) {
	s := makeState()
	assert.Equal(t, 0, s.Index(0, 0))
	assert.Equal(t, 1, s.Index(0, 1))
	assert.Equal(t, 3, s.Index(1, 0))
	assert.Equal(t, 6, s.Index(2, 0))
	assert.Equal(t, 7, s.Index(2, 1))
	assert.Equal(t, 8, s.Index(2, 2))
}

func TestNeighbors(t *testing.T) {
	s := makeState()
	assert.Equal(
		t,
		[]int{1, 3},
		s.Neighbors(0),
	)
	assert.Equal(
		t,
		[]int{3, 5, 1, 7},
		s.Neighbors(4),
	)
	assert.Equal(
		t,
		[]int{1, 5},
		s.Neighbors(2),
	)
	assert.Equal(
		t,
		[]int{0, 2, 4},
		s.Neighbors(1),
	)
}

func TestSimpleCapture(t *testing.T) {
	s := fromString([]string{
		"O..",
		"XX.",
		"...",
	})
	m := makeMove(s, 1)
	s.MakeMove(m)
	expected := []string{
		".X.",
		"XX.",
		"...",
	}
	assert.Equal(t, expected, s.String())
}

func TestClearBoard(t *testing.T) {
	s := fromString([]string{
		"OOO",
		"O.O",
		"OOO",
	})
	m := makeMove(s, 4)
	s.MakeMove(m)
	expected := []string{
		"...",
		".X.",
		"...",
	}
	assert.Equal(t, expected, s.String())
}

func TestWeirdPlacement(t *testing.T) {
	s := fromString([]string{
		".X.",
		"..O",
		"...",
	})
	m := makeMove(s, 8)
	s.MakeMove(m)
	expected := []string{
		".X.",
		"..O",
		"..X",
	}
	assert.Equal(t, expected, s.String())
}
