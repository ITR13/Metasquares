package main

import (
	"MetasquaresAI/AIs"
	"MetasquaresAI/Animators"
	"MetasquaresAI/Engine"
	"fmt"
)

func main() {
	for w := 3; w < 20; w++ {
		for h := w; h == w; h++ {
			game := Engine.MakeGame(w, h,
				AI.GetBruteForcer(nil, &AI.HighBlocker{}),
				&AI.HighBlocker{},
			)
			aniBoard := Engine.MakeBoard(w, h)
			game.Animator = Animators.Text{true, false, aniBoard}
			for i := 1; i < w && i < h; i++ {
				aniBoard.RegisterSquaresWithDistance(i)
			}
			winner := game.Play()
			fmt.Printf("%d Won %dx%d!\n", winner, w, h)
		}
	}
}