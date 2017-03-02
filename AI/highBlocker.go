package AI

import (
	"github.com/ITR13/metasquares/engine"
)

type TileTakeInfo struct {
	x, y  int
	count []int
	score []int
}

type HighBlocker struct{}

func (highBlocker *HighBlocker) Init(w, h, players int) {}

func (hb *HighBlocker) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	take := Taker(board, func(c Engine.Color) bool {
		return c != color && c != Engine.Empty && c != Engine.Mixed
	})
	if take.x == -1 && take.y == -1 {
		take = Taker(board, func(c Engine.Color) bool {
			return c == color || c == Engine.Empty
		})
	}

	return take.x, take.y
}

func (highBlocker *HighBlocker) Placed(x, y int, c Engine.Color) {}

type HighTaker struct{}

func (highTaker *HighTaker) Init(w, h, players int) {}

func (ht *HighTaker) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {

	take := Taker(board, func(c Engine.Color) bool {
		return c == color || c == Engine.Empty
	})
	if take.x == -1 && take.y == -1 {
		take = Taker(board, func(c Engine.Color) bool {
			return c != color && c != Engine.Empty && c != Engine.Mixed
		})
	}

	return take.x, take.y
}

func (highTaker *HighTaker) Placed(x, y int, c Engine.Color) {}

func Taker(board Engine.BoardAbstract,
	take func(Engine.Color) bool) TileTakeInfo {
	htti := TileTakeInfo{-1, -1, make([]int, 4), make([]int, 4)}

	w, h := board.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			tile := board.GetColor(x, y)
			if tile == Engine.Empty {
				tbi := TileTakeInfo{x, y, make([]int, 4), make([]int, 4)}
				colors, placed, score := board.GetShapes(x, y)
				for i := 0; i < len(colors); i++ {
					if take(colors[i]) {
						tbi.count[placed[i]]++
						tbi.score[placed[i]] += score[i]
					}
				}

				for i := 3; i >= 0; i-- {
					if tbi.count[i] > htti.count[i] {
						htti = tbi
						break
					} else if tbi.count[i] < htti.count[i] {
						break
					}
					if tbi.score[i] > htti.score[i] {
						htti = tbi
						break
					} else if tbi.score[i] < htti.score[i] {
						break
					}
				}
			}
		}
	}

	return htti
}

type MixedTaker struct{}

func (mixedTaker *MixedTaker) Init(w, h, players int) {}

func (mt *MixedTaker) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {

	take := Taker(board, func(c Engine.Color) bool {
		return c == color || c == Engine.Empty
	})
	block := Taker(board, func(c Engine.Color) bool {
		return c != color && c != Engine.Empty && c != Engine.Mixed
	})
	mixed := Taker(board, func(c Engine.Color) bool {
		return c != Engine.Empty && c != Engine.Mixed
	})

	if (block.x != -1 || block.y != -1) &&
		((block.x == take.x && block.y == take.y) ||
			(block.x == mixed.x && block.y == mixed.y)) {
		return block.x, block.y
	} else if (take.x != -1 || take.y != -1) &&
		take.x == mixed.x && take.y == mixed.y {
		return take.x, take.y
	}

	highest := block
	for i := 3; i >= 0; i-- {
		if block.count[i] > mixed.count[i] {
			break
		} else if block.count[i] < mixed.count[i] {
			highest = mixed
			break
		}
	}

	for i := 3; i >= 0; i-- {
		if highest.count[i] > take.count[i] {
			break
		} else if highest.count[i] < take.count[i] {
			highest = take
			break
		}
	}

	if highest.x == -1 && highest.y == -1 {
		w, h := board.GetSize()
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if board.GetColor(x, y) == Engine.Empty {
					return x, y
				}
			}
		}
	}

	return highest.x, highest.y
}

func (mixedTaker *MixedTaker) Placed(x, y int, c Engine.Color) {}
