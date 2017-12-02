package lib

import (
	"go-mcts"
	"set"
)

// Index returns the array index for position (row, col)
func (s *GoState) Index(row int, col int) int {
	return row*s.Size + col
}

// RowCol calculates the row and column for position p
func (s *GoState) RowCol(p int) (int, int) {
	row := p / s.Size
	col := p % s.Size
	return row, col
}

// Set sets the board to color at position p
func (s *GoState) Set(p int, c int8) {
	s.Board[p] = c
	s.updateBoard(p, c)
}

// MakeMove takes a GoMove and sets the board at position m.Index to m.Player
// If the passed move is pass, nothing will happen
func (s *GoState) MakeMove(mm mcts.Move) {
	var m *GoMove = mm.(*GoMove)
	if m.Pass {
		return
	}
	s.Set(m.Index, m.Player)
	s.update(m)
	s.CurrentPlayer *= -1
}

// update refreshes the game state after playing a move
func (s *GoState) update(m *GoMove) {
	visited := make([]bool, len(s.Board))
	affected := set.New(set.ThreadSafe)
	s.Flood(FirstSweepM, m.Index, m.Player, visited, affected)
	allVisited := make([]bool, len(s.Board))
	for _, positionRaw := range affected.List() {
		position := positionRaw.(int)
		if allVisited[position] {
			continue
		}
		positionVisited := make([]bool, len(s.Board))
		counted := set.New(set.ThreadSafe)
		s.Count(position, m.Player, positionVisited, counted)
		for vi, v := range positionVisited {
			if v {
				s.liberties[vi] = counted.Size()
				allVisited[vi] = true
			}
		}
	}
}

// Get returns the color at position p
func (s *GoState) Get(p int) int8 {
	return s.Board[p]
}

// Empty returns whether the field at position p is empty
func (s *GoState) Empty(p int) bool {
	return s.Board[p] == E
}

// Neighbors returns the index of all neighbors for position p
func (s *GoState) Neighbors(p int) []int {
	neighbors := []int{}
	row, col := s.RowCol(p)
	if col > 0 {
		neighbors = append(neighbors, s.Index(row, col-1))
	}
	if col < s.Size-1 {
		neighbors = append(neighbors, s.Index(row, col+1))
	}
	if row > 0 {
		neighbors = append(neighbors, s.Index(row-1, col))
	}
	if row < s.Size-1 {
		neighbors = append(neighbors, s.Index(row+1, col))
	}
	return neighbors
}

// Liberties returns the amount of liberties for position p
func (s *GoState) Liberties(p int) int {
	return s.liberties[p]
}

// CanMove returns true if the current player can move to p
// Derived from https://github.com/wmontgomery4/go/blob/523c04faff12faf6ed0c5eb017ed7edd8134a60f/src/engine.py#L26
func (s *GoState) CanMove(p int) bool {
	f := s.Get(p)
	if f != E {
		return false
	}
	if p == s.Ko {
		return false
	}
	neighbors := s.Neighbors(p)
	for _, n := range neighbors {
		nC := s.Get(n)
		l := s.Liberties(n)
		if nC == E {
			return true
		} else if nC == -s.CurrentPlayer && l == 1 {
			return true
		} else if nC == s.CurrentPlayer && l > 1 {
			return true
		}
	}
	return false
}
