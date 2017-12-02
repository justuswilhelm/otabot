// board fill algorithms
package lib

import (
	"set"
)

// Flood floods a position with markers
func (s *GoState) Flood(m int, p int, c int8, visited []bool, affected set.Interface) {
	for _, n := range s.Neighbors(p) {
		switch m {
		case FirstSweepM:
			s.FirstSweep(n, c, visited, affected)
		case ReduceM:
			s.Reduce(n, c, visited)
		case RemoveM:
			s.Remove(n, c, visited, affected)
		case CountM:
			s.Count(n, c, visited, affected)
		}
	}
}

func (s *GoState) FirstSweep(p int, c int8, visited []bool, affected set.Interface) {
	if visited[p] || s.Board[p] != -c {
		return
	}
	if s.liberties[p] > 1 {
		s.Reduce(p, -c, visited)
	} else {
		s.Remove(p, -c, visited, affected)
	}
}

// Reduce decreases the liberties for a group by 1
func (s *GoState) Reduce(p int, c int8, visited []bool) {
	if visited[p] || s.Board[p] != c {
		return
	}
	visited[p] = true
	s.liberties[p] -= 1
	s.Flood(ReduceM, p, c, visited, nil)
}

// Remove removes a group  with color c starting at p
func (s *GoState) Remove(p int, c int8, visited []bool, affected set.Interface) {
	if visited[p] || s.Get(p) == E {
		return
	}
	if s.Board[p] == -c {
		affected.Add(p)
		return
	}
	visited[p] = true
	s.Board[p] = E
	s.liberties[p] = 0
	s.Flood(RemoveM, p, c, visited, affected)
}

// Count liberties across group with color c starting at p
func (s *GoState) Count(p int, c int8, visited []bool, affected set.Interface) {
	if visited[p] || s.Board[p] == -c {
		return
	}
	if s.Board[p] == E {
		affected.Add(p)
		return
	}
	visited[p] = true
	s.Flood(CountM, p, c, visited, affected)
}

// updateBoard updates the board state after placing a stone
func (s *GoState) updateBoard(p int, c int8) {
	size := len(s.Board)
	visited := make([]bool, size)
	affected := set.New(set.ThreadSafe)
	affected.Add(p)
	s.Flood(FirstSweepM, p, c, visited, affected)
	allVisited := make([]bool, size)
	for _, pRaw := range affected.List() {
		p := pRaw.(int)
		if allVisited[p] {
			continue
		}
		visited = make([]bool, size)
		counted := set.New(set.ThreadSafe)
		s.Count(p, c, visited, counted)
		s.liberties[p] = counted.Size()
		for i, v := range allVisited {
			allVisited[i] = v || visited[i]
		}
	}
}
