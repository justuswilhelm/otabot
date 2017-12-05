package mcts

// Scorer is the function an AI can use to just the benifit of an outcome from the eyes of a particular player.
type Scorer func(playerId int8, state GameState) float64
