/*
    This file is part of Metasquares.

    Metasquares is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Metasquares is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with Metasquares.  If not, see <http://www.gnu.org/licenses/>.
*/
	
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
				c, x+1, y+1, board.GetScores())
		} else {
			fmt.Printf(" *\tPlaced %d at %d,%d\n", c, x+1, y+1)
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
		fmt.Printf(" *\tPlaced %d at %d,%d\n", c, x+1, y+1)
	}
}
