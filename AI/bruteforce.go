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
	
package AI

import (
	"fmt"

	"github.com/ITR13/metasquares/engine"
)

type BruteForcer struct {
	log       int
	PlayerAIs []Engine.Controller
	myBoard   *Engine.Board
	players   Engine.Color
	path      *[2]interface{}
}

func GetBruteForcer(save bool, log int,
	players ...Engine.Controller) *BruteForcer {

	if save {
		return &BruteForcer{log, players, nil, 0, &[2]interface{}{nil, nil}}
	} else {
		return &BruteForcer{log, players, nil, 0, nil}
	}
}

func (bruteForcer *BruteForcer) Init(w, h, players int) {
	bruteForcer.myBoard = Engine.MakeBoard(w, h)
	for i := 1; i < w && i < h; i++ {
		bruteForcer.myBoard.RegisterSquaresWithDistance(i)
	}
	bruteForcer.players = Engine.Color(players)

	if bruteForcer.path != nil {
		bruteForcer.path = &[2]interface{}{nil, nil}
	}

	for i := 0; i < len(bruteForcer.PlayerAIs); i++ {
		if bruteForcer.PlayerAIs[i] != nil {
			bruteForcer.PlayerAIs[i].Init(w, h, players)
		}
	}
}

func (bruteForcer *BruteForcer) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	if bruteForcer.path != nil && bruteForcer.path[0] != nil {
		v := bruteForcer.path[0].(*Engine.Vector)
		return v.X, v.Y
	}

	x, y, _, path := bruteForcer.GetWinner(color, 0)
	if bruteForcer.path != nil {
		bruteForcer.path = path
	}
	return x, y
}

func (bruteForcer *BruteForcer) Placed(x, y int, c Engine.Color) {
	bruteForcer.myBoard.At(x, y).SetColor(c)

	if bruteForcer.path != nil && bruteForcer.path[0] != nil {
		v := bruteForcer.path[0].(*Engine.Vector)
		if v.X != x || v.Y != y {
			bruteForcer.path = &[2]interface{}{nil, nil}
			if bruteForcer.log > 0 {
				fmt.Printf("Expected %d,%d, but got %d,%d\n", v.X, v.Y, x, y)
			}
		} else {
			if bruteForcer.path[1] != nil {
				bruteForcer.path = bruteForcer.path[1].(*[2]interface{})
			} else {
				bruteForcer.path = nil
			}
		}
	}

	for i := 0; i < len(bruteForcer.PlayerAIs); i++ {
		if bruteForcer.PlayerAIs[i] != nil {
			bruteForcer.PlayerAIs[i].Placed(x, y, c)
		}
	}
}

func (bruteForcer *BruteForcer) GetWinner(c Engine.Color,
	layer int) (int, int, Engine.Color, *[2]interface{}) {

	if c > bruteForcer.players {
		c = 1
	}
	board := bruteForcer.myBoard
	w, h := board.GetSize()

	if len(bruteForcer.PlayerAIs) >= int(c) &&
		bruteForcer.PlayerAIs[c-1] != nil && layer >= bruteForcer.log {
		x, y := bruteForcer.PlayerAIs[c-1].Place(c, board)
		if x != -1 || y != -1 {
			bruteForcer.Placed(x, y, c)
			defer bruteForcer.Placed(x, y, Engine.Empty)
		}
		_, _, moveWinner, path := bruteForcer.GetWinner(c+1, layer+1)
		if path != nil {
			return x, y, moveWinner, &[2]interface{}{&Engine.Vector{x, y}, path}
		} else {
			return x, y, moveWinner, nil
		}
	}

	winner := Engine.Empty
	wx, wy := -1, -1
	var wpath *[2]interface{}

	winpaths, drawpaths, losspaths := 0, 0, 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if board.GetColor(x, y) == Engine.Empty {
				bruteForcer.Placed(x, y, c)
				if layer+1 < bruteForcer.log {
					fmt.Printf("--- %d,%d ---\n", x, y)
				}
				_, _, moveWinner, path := bruteForcer.GetWinner(c+1, layer+1)
				bruteForcer.Placed(x, y, Engine.Empty)

				switch moveWinner {
				case c:
					winpaths++
				case Engine.Mixed:
					drawpaths++
				default:
					losspaths++
				}

				if winner == c {
					continue
				}

				if moveWinner == c {
					if layer >= bruteForcer.log {
						if path != nil {
							return x, y, moveWinner,
								&[2]interface{}{&Engine.Vector{x, y}, path}
						} else {
							return x, y, moveWinner, nil
						}
					}
					wx, wy = x, y
					winner = moveWinner
					wpath = &[2]interface{}{&Engine.Vector{x, y}, path}
				} else if winner == Engine.Empty {
					wx, wy = x, y
					winner = moveWinner
					wpath = &[2]interface{}{&Engine.Vector{x, y}, path}
				} else if winner != Engine.Mixed {
					wx, wy = x, y
					winner = moveWinner
					wpath = &[2]interface{}{&Engine.Vector{x, y}, path}
				}
			}
		}
	}

	if winner == Engine.Empty {
		return -1, -1, board.GetLeader(),
			&[2]interface{}{nil, nil}
	}

	if layer < bruteForcer.log {
		fmt.Printf("Wins:\t%d\nDraws:\t%d\nLoss:\t%d\nTotal:\t%d\n",
			winpaths, drawpaths, losspaths, winpaths+drawpaths+losspaths)
		if layer != 0 {
			fmt.Printf("^^^ %d,%d ^^^\n", wx, wy)
		}
	}

	if len(bruteForcer.PlayerAIs) >= int(c) &&
		bruteForcer.PlayerAIs[c-1] != nil && layer < bruteForcer.log {
		x, y := bruteForcer.PlayerAIs[c-1].Place(c, board)
		fmt.Printf("^^^ %d,%d ^^^\n", x, y)
		if x != -1 || y != -1 {
			bruteForcer.Placed(x, y, c)
			defer bruteForcer.Placed(x, y, Engine.Empty)
		}
		_, _, moveWinner, path := bruteForcer.GetWinner(c+1, layer+1)
		if path != nil {
			return x, y, moveWinner, &[2]interface{}{&Engine.Vector{x, y}, path}
		} else {
			return x, y, moveWinner, nil
		}
	}

	return wx, wy, winner, wpath
}
