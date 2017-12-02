package main

import (
	"log"
	"time"

	"go-mcts"

	"github.com/justuswilhelm/otabot/lib"
)

const (
	ucbC        = 0.5
	iterations  = 100
	simulations = 100
)

func main() {
	board := lib.NewGoState(7, 5.5)
	i := 1000
	for len(board.AvailableMoves()) > 0 {
		log.Printf("Board: %s", board.Log())
		log.Printf("Score: %0.2f", board.Score())
		move := mcts.Uct(
			board,
			iterations,
			simulations,
			ucbC,
			board.CurrentPlayer,
			lib.Score,
		)
		board.MakeMove(move)
		i -= 1
		if i == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("Game Over")
}
