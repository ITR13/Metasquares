package AI

import (
	"testing"

	"github.com/ITR13/metasquares/engine"
)

func TestInit(t *testing.T) {
	br := BruteForcer{}
	for w := 2; w < 5; w++ {
		for h := 2; h < 5; h++ {
			for players := 0; players < 4; players++ {
				br.Init(w, h, players)
				W, H := br.myBoard.GetSize()
				if w != W || h != H {
					t.Fatalf("%d,%d != %d,%d\n", w, h, W, H)
				}
				if int(br.players) != players {
					t.Fatalf("%d != %d\n", players, br.players)
				}
			}
		}
	}
}

func TestGetWinner(t *testing.T) {
	br := GetBruteForcer(false, 1, nil, nil)
	br.Init(3, 3, 2)
	br.Placed(0, 0, Engine.Red)
	br.Placed(0, 2, Engine.Red)
	br.Placed(2, 0, Engine.Red)
	br.Placed(0, 1, Engine.NotAColor)

	x, y, winner, _ := br.GetWinner(Engine.Red, 0)
	if x != 2 || y != 2 || winner != Engine.Red {
		t.Errorf("Wanted %d at %d,%d, got %d at %d,%d",
			Engine.Red, 2, 2, winner, x, y)
	} else {
		t.Log("Passed Easy Win")
	}

	x, y, winner, _ = br.GetWinner(Engine.Green, 0)
	if x != 2 || y != 2 || winner != Engine.Mixed {
		t.Errorf("Wanted %d at %d,%d, got %d at %d,%d",
			Engine.Mixed, 2, 2, winner, x, y)
	} else {
		t.Log("Passed Only Counter")
	}

	br.Init(4, 4, 2)
	br.Placed(1, 0, Engine.Red)
	br.Placed(1, 1, Engine.Red)
	br.Placed(2, 2, Engine.Red)
	br.Placed(2, 3, Engine.Red)
	br.Placed(0, 0, Engine.NotAColor)
	br.Placed(0, 1, Engine.NotAColor)
	br.Placed(0, 3, Engine.NotAColor)
	br.Placed(3, 0, Engine.NotAColor)
	br.Placed(3, 2, Engine.NotAColor)
	br.Placed(3, 3, Engine.NotAColor)

	x, y, winner, _ = br.GetWinner(Engine.Green, 0)
	if winner != Engine.Red {
		t.Errorf("Wanted %d at ?,?, got %d at %d,%d",
			Engine.Red, winner, x, y)
	} else {
		t.Log("Passed Predetermined")
	}

}

func TestGetWinnerWithAI(t *testing.T) {
	br := GetBruteForcer(true, 1, nil, &FirstAvailable{})
	br.Init(3, 3, 2)
	br.Placed(0, 0, Engine.Red)
	br.Placed(0, 2, Engine.Red)
	br.Placed(2, 0, Engine.Red)
	br.Placed(0, 1, Engine.NotAColor)

	x, y, winner, _ := br.GetWinner(Engine.Red, 0)
	if winner != Engine.Red {
		t.Errorf("Wanted %d at ?,?, got %d at %d,%d",
			Engine.Red, winner, x, y)
	} else {
		t.Log("Passed Easy Win")
	}

	x, y, winner, _ = br.GetWinner(Engine.Green, 1)
	if winner != Engine.Red {
		t.Errorf("Wanted %d at ?,?, got %d at %d,%d",
			Engine.Red, winner, x, y)
	} else {
		t.Log("Passed Only Counter")
	}

	br.Init(4, 4, 2)
	br.Placed(1, 0, Engine.Red)
	br.Placed(1, 1, Engine.Red)
	br.Placed(2, 2, Engine.Red)
	br.Placed(2, 3, Engine.Red)
	br.Placed(0, 0, Engine.NotAColor)
	br.Placed(0, 1, Engine.NotAColor)
	br.Placed(0, 3, Engine.NotAColor)
	br.Placed(3, 0, Engine.NotAColor)
	br.Placed(3, 2, Engine.NotAColor)
	br.Placed(3, 3, Engine.NotAColor)

	x, y, winner, _ = br.GetWinner(Engine.Green, 0)
	if winner != Engine.Red {
		t.Errorf("Wanted %d at ?,?, got %d at %d,%d",
			Engine.Red, winner, x, y)
	} else {
		t.Log("Passed Predetermined")
	}

}
