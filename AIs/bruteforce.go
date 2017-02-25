package AI

import (
	"MetasquaresAI/Engine"
)

type BruteForcer struct {
	PlayerAIs []Engine.Controller
	myBoard   *Engine.Board
	players   Engine.Color
}

func GetBruteForcer(players ...Engine.Controller) *BruteForcer {
	return &BruteForcer{players, nil, 0}
}

func (bruteForcer *BruteForcer) Init(w, h, players int) {
	bruteForcer.myBoard = Engine.MakeBoard(w, h)
	for i := 1; i < w && i < h; i++ {
		bruteForcer.myBoard.RegisterSquaresWithDistance(i)
	}
	bruteForcer.players = Engine.Color(players)

	for i := 0; i < len(bruteForcer.PlayerAIs); i++ {
		if bruteForcer.PlayerAIs[i] != nil {
			bruteForcer.PlayerAIs[i].Init(w, h, players)
		}
	}
}

func (bruteForcer *BruteForcer) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	x, y, _ := bruteForcer.GetWinner(color)
	return x, y
}

func (bruteForcer *BruteForcer) Placed(x, y int, c Engine.Color) {
	bruteForcer.myBoard.At(x, y).SetColor(c)
	for i := 0; i < len(bruteForcer.PlayerAIs); i++ {
		if bruteForcer.PlayerAIs[i] != nil {
			bruteForcer.PlayerAIs[i].Placed(x, y, c)
		}
	}
}

func (bruteForcer *BruteForcer) GetWinner(c Engine.Color) (x, y int,
	color Engine.Color) {

	if c > bruteForcer.players {
		c = 1
	}
	board := bruteForcer.myBoard
	w, h := board.GetSize()

	if len(bruteForcer.PlayerAIs) >= int(c) &&
		bruteForcer.PlayerAIs[c-1] != nil {
		x, y := bruteForcer.PlayerAIs[c-1].Place(c, board)
		if x != -1 || y != -1 {
			bruteForcer.Placed(x, y, c)
			defer bruteForcer.Placed(x, y, Engine.Empty)
		}
		_, _, moveWinner := bruteForcer.GetWinner(c + 1)
		return x, y, moveWinner
	}

	winner := Engine.Empty
	wx, wy := -1, -1
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if board.GetColor(x, y) == Engine.Empty {
				bruteForcer.Placed(x, y, c)
				_, _, moveWinner := bruteForcer.GetWinner(c + 1)
				bruteForcer.Placed(x, y, Engine.Empty)

				if moveWinner == c {
					return x, y, moveWinner
				}
				if winner == Engine.Empty {
					wx, wy = x, y
					winner = moveWinner
				} else if winner != Engine.Mixed {
					wx, wy = x, y
					winner = moveWinner
				}
			}
		}
	}

	if winner == Engine.Empty {
		return -1, -1, board.GetLeader()
	}

	return wx, wy, winner
}

type MonoBruteForcer struct {
	bruteForcer BruteForcer
	path        []int
	placed      int
}

func (mbf *MonoBruteForcer) Init(w, h, players int) {
	mbf.bruteForcer.Init(w, h, players)
	mbf.path = make([]int, w*h)
	mbf.placed = 0
}

func (mbf *MonoBruteForcer) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	if mbf.path[mbf.placed]!=-1{
		
	}
	x, y, _ := mbf.bruteForcer.GetWinner(color)
	return x, y
}

func (mbf *MonoBruteForcer) Placed(x, y int, c Engine.Color) {
	if color == Engine.Empty {
		if board.GetColor()
	}

	mbf.bruteForcer.Placed(x, y, c)
}
