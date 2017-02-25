package Animators

import (
	"github.com/ITR13/metasquares/engine"

	"fmt"
)

type Text struct {
	PrintScore, PrintBoard bool
	Board                  *Engine.Board
}

func (text Text) Placed(x, y int, c Engine.Color) {
	board := text.Board
	if board != nil {
		board.At(x, y).SetColor(c)
		if text.PrintScore {
			fmt.Printf(" *\tPlaced %d at %d,%d: %v\n",
				c, x, y, board.GetScores())
		} else {
			fmt.Printf(" *\tPlaced %d at %d,%d\n", c, x, y)
		}

		if text.PrintBoard {
			w, h := board.GetSize()
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					fmt.Print("---")
				}
				fmt.Println()
				for x := 0; x < w; x++ {
					tile := board.GetColor(x, y)
					fmt.Printf("|%d", tile)
				}
				fmt.Print("|\n")
			}
			for x := 0; x < w; x++ {
				fmt.Print("---")
			}
			fmt.Println()
		}
	} else {
		fmt.Printf(" *\tPlaced %d at %d,%d\n", c, x, y)
	}
}
