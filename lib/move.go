package lib

import (
	"go-mcts"
)

func (s *GoState) newMove(i int, p bool) mcts.Move {
	return &GoMove{
		Player: s.CurrentPlayer,
		Index:  i,
		Pass:   p,
	}
}

// Probability TODO
func (m *GoMove) Probability() float64 {
	return 1.0
}
