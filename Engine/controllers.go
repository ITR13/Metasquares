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
	
package Engine

import (
	"fmt"
	"math/rand"
)

type Game struct {
	board    *Board
	players  []Controller
	turn     int
	Animator Animator
}

type Animator interface {
	Placed(x, y int, color Color)
}

type BoardAbstract interface {
	GetColor(x, y int) Color
	GetShapes(x, y int) ([]Color, []int, []int)
	GetSize() (int, int)
}

type Controller interface {
	Init(w, h, players int)
	Place(color Color, board BoardAbstract) (int, int)
	Placed(x, y int, c Color)
}

func MakeGame(w, h int, players ...Controller) *Game {
	board := MakeBoard(w, h)
	for i := 1; i < w && i < h; i++ {
		board.RegisterSquaresWithDistance(i)
	}
	for i := 0; i < len(players); i++ {
		players[i].Init(w, h, len(players))
	}
	return &Game{board, players, 0, nil}
}

func (game *Game) AdvanceSingle(player uint8) {
	rand.Seed(int64(game.turn))
	x, y := game.players[player].Place(Color(player+1), game.board)
	tile := game.board.At(x, y)
	if tile == nil {
		fmt.Printf("Player %d tried to do an illegal move (%d,%d is nil)\n",
			player+1, x, y)
		return
	} else if tile.color != Empty {
		fmt.Printf("Player %d tried to do an illegal move (%d,%d is %d)\n",
			player+1, x, y, tile.color)
		return
	}
	tile.SetColor(Color(player + 1))
	if game.Animator != nil {
		game.Animator.Placed(x, y, tile.color)
	}
	for i := 0; i < len(game.players); i++ {
		game.players[i].Placed(x, y, Color(player+1))
	}
}

func (game *Game) Advance() {
	game.turn++
	for i := 0; i < len(game.players); i++ {
		game.AdvanceSingle(uint8(i))
	}
}

func (game *Game) Play() Color {
	piecesLeft := game.board.w * game.board.h
	skips := 0
	for game.turn*len(game.players) < piecesLeft {
		if skips > 10 {
			fmt.Println("Skipped more than 10 times, skipped")
			break
		}

		for game.turn*len(game.players) < piecesLeft {
			game.Advance()
		}
		for x := 0; x < game.board.w; x++ {
			for y := 0; y < game.board.h; y++ {
				if game.board.GetColor(x, y) == Empty {
					piecesLeft++
				}
			}
		}
		skips++
	}
	return game.board.GetLeader()
}

func (board *Board) GetColor(x, y int) Color {
	tile := board.At(x, y)
	if tile == nil {
		fmt.Errorf("Tried to get color at illegal position (%d,%d is nil)\n",
			x, y)
		return NotAColor
	}
	return tile.color
}
func (board *Board) GetShapes(x, y int) ([]Color, []int, []int) {
	tile := board.At(x, y)
	if tile == nil {
		fmt.Errorf("Tried to get shapes at illegal position (%d,%d is nil)\n",
			x, y)
		return make([]Color, 0), make([]int, 0), make([]int, 0)
	}
	colors := make([]Color, len(tile.shapes))
	filled := make([]int, len(tile.shapes))
	values := make([]int, len(tile.shapes))

	for i := 0; i < len(tile.shapes); i++ {
		colors[i] = tile.shapes[i].color
		filled[i] = tile.shapes[i].filled
		values[i] = tile.shapes[i].worth
	}
	return colors, filled, values
}
