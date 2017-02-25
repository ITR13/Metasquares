package AI

import (
	"MetasquaresAI/Engine"
)

type FirstAvailable struct {
	w, h, players int
}

func (fa *FirstAvailable) Init(w, h, players int) {
	fa.w, fa.h, fa.players = w, h, players
}

func (fa *FirstAvailable) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	for x := 0; x < fa.w; x++ {
		for y := 0; y < fa.h; y++ {
			if board.GetColor(x, y) == Engine.Empty {
				return x, y
			}
		}
	}
	return -1, -1
}

func (fa *FirstAvailable) Placed(x, y int, c Engine.Color) {}
