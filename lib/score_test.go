package lib

import (
	"testify/assert"
	"testing"
)

func scoreFromBoard(board []string) map[int8]int {
	s := fromString(board)
	return s.ScoreComponents()
}

func TestMakeEmptyRegionEmpty(t *testing.T) {
	s := fromString([]string{
		"..",
		"..",
	})
	reg := s.MakeEmptyRegion(0)
	assert.Equal(t, 4, reg.Points.Size())
	assert.Equal(t, 0, reg.NeighboringColors.Size())
}

func TestMakeEmptyRegionOneColor(t *testing.T) {
	s := fromString([]string{
		"..",
		"XX",
	})
	reg := s.MakeEmptyRegion(0)
	assert.Equal(t, 2, reg.Points.Size())
	assert.Equal(t, 1, reg.NeighboringColors.Size())
	reg = s.MakeEmptyRegion(1)
	assert.Equal(t, 2, reg.Points.Size())
	assert.Equal(t, 1, reg.NeighboringColors.Size())
	reg = s.MakeEmptyRegion(2)
	assert.Equal(t, 3, reg.Points.Size())
	assert.Equal(t, 1, reg.NeighboringColors.Size())
}

func TestMakeEmptyRegionTwoColors(t *testing.T) {
	s := fromString([]string{
		"..",
		"XO",
	})
	reg := s.MakeEmptyRegion(0)
	assert.Equal(t, 2, reg.Points.Size())
	assert.Equal(t, 2, reg.NeighboringColors.Size())
}

func TestScoreEmpty(t *testing.T) {
	score := scoreFromBoard([]string{
		"...",
		"...",
		"...",
	})
	assert.Equal(t, 0, score[B])
	assert.Equal(t, 0, score[W])
}

func TestScoreDiamond(t *testing.T) {
	score := scoreFromBoard([]string{
		".X.",
		"X.X",
		".X.",
	})
	assert.Equal(t, 9, score[B])
	assert.Equal(t, 0, score[W])

	score = scoreFromBoard([]string{
		".O.",
		"O.O",
		".O.",
	})
	assert.Equal(t, 0, score[B])
	assert.Equal(t, 9, score[W])
}

func TestScoreGroup(t *testing.T) {
	score := scoreFromBoard([]string{
		".XO.",
		".XO.",
		".XO.",
		".XO.",
	})
	assert.Equal(t, 8, score[B])
	assert.Equal(t, 8, score[W])
}

func TestScoreNested(t *testing.T) {
	score := scoreFromBoard([]string{
		".XO.",
		".XOO",
		".XXX",
		"....",
	})
	assert.Equal(t, 12, score[B])
	assert.Equal(t, 4, score[W])
}
