package main

import (
	"fmt"
	"os"

	"github.com/ITR13/metasquares/AI"
	"github.com/ITR13/metasquares/animator"
	"github.com/ITR13/metasquares/engine"
)

func main() {
	winners := [3]int{0, 0, 0}
	for w := 8; w < 12; w++ {
		for h := w; h == w; h++ {
			game := Engine.MakeGame(w, h,
				&AI.MixedTaker{},
				&AI.PlayerC{os.Stdin},
			)
			aniBoard := Engine.MakeBoard(w, h)
			game.Animator = Animators.Text{true, false, aniBoard}
			for i := 1; i < w && i < h; i++ {
				aniBoard.RegisterSquaresWithDistance(i)
			}
			winner := game.Play()
			fmt.Printf("%d Won %dx%d!\n", winner, w, h)
			if winner == 9 {
				winner = 3
			}
			winners[winner-1]++
		}
	}
	fmt.Printf("\nWins:\n P1:\t%d\n P2:\t%d\n Draw:\t%d\n",
		winners[0], winners[1], winners[2])
}
