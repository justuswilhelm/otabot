package lib

import (
	"set"
	"testify/assert"
	"testing"
)

func countTrue(l []bool) int {
	count := 0
	for _, v := range l {
		if v {
			count += 1
		}
	}
	return count
}

func (s *GoState) makeVisited() []bool {
	return make([]bool, len(s.Board))
}

func TestFirstSweep(t *testing.T) {
	s := fromString([]string{
		".XO",
		".XX",
		"...",
	})
	p := 2
	visited := s.makeVisited()
	affected := set.New(set.ThreadSafe)
	s.FirstSweep(p, W, visited, affected)
	assert.Equal(t, E, s.Get(p))
	assert.Equal(t, 0, affected.Size())
	expected := []string{
		".X.",
		".XX",
		"...",
	}
	assert.Equal(t, expected, s.String())
}

func TestReduce(t *testing.T) {
	s := fromString([]string{
		".X.",
		".XO",
		"...",
	})
	p := 1
	assert.Equal(t, 2, s.Liberties(p))
	visited := s.makeVisited()
	s.Reduce(p, B, visited)
	assert.Equal(t, 1, s.Liberties(p))
	expected := []string{
		".X.",
		".XO",
		"...",
	}
	assert.Equal(t, expected, s.String())
}

func TestRemove(t *testing.T) {
	s := fromString([]string{
		".X.",
		".XO",
		"...",
	})
	p := 1
	visited := s.makeVisited()
	affected := set.New(set.ThreadSafe)
	s.Remove(p, B, visited, affected)
	assert.Equal(t, E, s.Get(p))
	assert.Equal(t, 1, affected.Size())
	expected := []string{
		"...",
		"..O",
		"...",
	}
	assert.Equal(t, expected, s.String())
}

func TestCount(t *testing.T) {
	s := fromString([]string{
		".X.",
		".XO",
		"...",
	})
	p := 1
	visited := s.makeVisited()
	affected := set.New(set.ThreadSafe)
	s.Count(p, B, visited, affected)
	assert.Equal(t, 4, affected.Size())
	assert.Equal(t, 2, countTrue(visited))
	expected := []string{
		".X.",
		".XO",
		"...",
	}
	assert.Equal(t, expected, s.String())
}
