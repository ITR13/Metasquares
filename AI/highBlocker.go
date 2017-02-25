package AI

import "github.com/ITR13/metasquares/engine"

type HighBlocker struct {
	w, h, players int
}

type TileBlockInfo struct {
	count []int
	score []int
}

func (highBlocker *HighBlocker) Init(w, h, players int) {
	highBlocker.w, highBlocker.h, highBlocker.players = w, h, players
}

func (hb *HighBlocker) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	hx, hy, htbi := -1, -1, TileBlockInfo{make([]int, 4), make([]int, 4)}

	for x := 0; x < hb.w; x++ {
		for y := 0; y < hb.h; y++ {
			tile := board.GetColor(x, y)
			if tile == Engine.Empty {
				tbi := TileBlockInfo{make([]int, 4), make([]int, 4)}
				colors, placed, score := board.GetShapes(x, y)
				for i := 0; i < len(colors); i++ {
					if colors[i] != color &&
						colors[i] != Engine.Empty && colors[i] != Engine.Mixed {
						tbi.count[placed[i]]++
						tbi.score[placed[i]] += score[i]
					}
				}

				for i := 3; i >= 0; i-- {
					if tbi.count[i] > htbi.count[i] {
						hx, hy = x, y
						htbi = tbi
						break
					} else if tbi.count[i] < htbi.count[i] {
						break
					}
					if tbi.score[i] > htbi.score[i] {
						hx, hy = x, y
						htbi = tbi
						break
					} else if tbi.score[i] < htbi.score[i] {
						break
					}
				}
			}
		}
	}

	return hx, hy
}

func (highBlocker *HighBlocker) Placed(x, y int, c Engine.Color) {}
