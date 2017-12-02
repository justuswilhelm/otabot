// Package lib functions for go-mcts interface compatibility
package lib

import (
	"go-mcts"
)

// Clone a GoState
func (s *GoState) Clone() mcts.GameState {
	board := make([]int8, len(s.Board))
	copy(board, s.Board)
	liberties := make([]int, len(s.Board))
	copy(liberties, s.liberties)
	return &GoState{
		CurrentPlayer: s.CurrentPlayer,
		Board:         board,
		liberties:     liberties,
		Komi:          s.Komi,
		Size:          s.Size,
		Ko:            s.Ko,
	}
}

// AvailableMoves TODO
func (s *GoState) AvailableMoves() []mcts.Move {
	var moves []mcts.Move
	for i := range s.Board {
		if s.CanMove(i) {
			moves = append(moves, s.newMove(i, false))
		}
	}
	moves = append(moves, s.newMove(0, true))
	return moves
}

// RandomizeUnknowns has no effect since Go has no random hidden information.
func (s *GoState) RandomizeUnknowns() {}

// Score returns the score for the current player
func Score(player int8, state mcts.GameState) float64 {
	s := state.(*GoState)
	score := s.Score()
	if player == W {
		score *= -1
	}
	return score
}
