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
		fmt.Errorf("Player %d tried to do an illegal move (%d,%d is nil)",
			player+1, x, y)
		return
	} else if tile.color != Empty {
		fmt.Errorf("Player %d tried to do an illegal move (%d,%d is %d)",
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
	for game.turn*len(game.players) <= game.board.w*game.board.h {
		game.Advance()
	}
	return game.board.GetLeader()
}

func (board *Board) GetColor(x, y int) Color {
	tile := board.At(x, y)
	if tile == nil {
		fmt.Errorf("Tried to get color at illegal position (%d,%d is nil)",
			x, y)
		return NotAColor
	}
	return tile.color
}
func (board *Board) GetShapes(x, y int) ([]Color, []int, []int) {
	tile := board.At(x, y)
	if tile == nil {
		fmt.Errorf("Tried to get shapes at illegal position (%d,%d is nil)",
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
