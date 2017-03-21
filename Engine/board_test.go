/*
    This file is part of InvertoTanks.

    Foobar is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    InvertoTanks is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with InvertoTanks.  If not, see <http://www.gnu.org/licenses/>.
*/
	
package Engine

import (
	"testing"
)

func TestSetColor(t *testing.T) {
	l := 6
	if testing.Short() {
		l = 3
	}

	tiles := make([]*Tile, l)
	shapes := make([]*Shape, l)
	for i := 0; i < l; i++ {
		tiles[i] = &Tile{Empty, 0, 0, make([]*Shape, l-i)}
	}

	for i := 0; i < l; i++ {
		shapes[i] = &Shape{make([]*Tile, i+1), 0, Empty, 0}

		for j := 0; j < i+1; j++ {
			tiles[j].shapes[i-j] = shapes[i]
			shapes[i].nodes[j] = tiles[j]
		}
	}

	for i := 0; i < l; i++ {
		tiles[i].SetColor(Red)
		for j := 0; j < l; j++ {
			if j > i {
				if tiles[j].color != Empty {
					t.Fatalf("Tile #%d is c%d, should be %d (iteration %d)\n",
						j, tiles[i].color, Empty, i)
				}
				if shapes[j].filled != i+1 {
					t.Fatalf("Shape #%d has f%d, should be %d (iteration %d)\n",
						j, shapes[i].filled, i+1, i)
				}
			} else {
				if tiles[j].color != Red {
					t.Fatalf("Tile #%d is c%d, should be %d (iteration %d)\n",
						j, tiles[i].color, Red, i)
				}
				if shapes[j].filled != j+1 {
					t.Fatalf("Shape #%d has f%d, should be %d (iteration %d)\n",
						j, shapes[i].filled, j+1, i)
				}
			}
			if shapes[j].color != Red {
				t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
					j, shapes[i].color, Red, i)
			}

		}
	}
	t.Log("Passed setting all shapes to Red")

	for i := 0; i < l; i++ {
		tiles[i].SetColor(Green)
		for j := 0; j < l; j++ {
			if j > i {
				if tiles[j].color != Red {
					t.Fatalf("Tile #%d is c%d, should be %d (iteration %d)\n",
						j, tiles[i].color, Red, i)
				}
				if shapes[j].color != Mixed {
					t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
						j, shapes[i].color, Mixed, i)
				}
			} else {
				if tiles[j].color != Green {
					t.Fatalf("Tile #%d is c%d, should be %d (iteration %d)\n",
						j, tiles[i].color, Green, i)
				}
				if shapes[j].color != Green {
					t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
						j, shapes[i].color, Green, i)
				}
			}
			if shapes[j].filled != j+1 {
				t.Fatalf("Shape #%d has f%d, should be %d (iteration %d)\n",
					j, shapes[i].filled, j+1, i)
			}
		}
	}
	t.Log("Passed setting all shapes to Green")

	for i := 0; i < l; i++ {
		tiles[i].SetColor(Empty)
		for j := 0; j < l; j++ {
			if j > i {
				if tiles[j].color != Green {
					t.Fatalf("Tile #%d is c%d, should be %d (iteration %d)\n",
						j, tiles[i].color, Green, i)
				}
				if shapes[j].color != Green {
					t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
						j, shapes[i].color, Green, i)
				}
				if shapes[j].filled != j-i {
					t.Fatalf("Shape #%d has f%d, should be %d (iteration %d)\n",
						j, shapes[i].filled, j-i, i)
				}
			} else {
				if tiles[j].color != Empty {
					t.Fatalf("Tile #%d is c%d, should be %d (iteration %d)\n",
						j, tiles[i].color, Empty, i)
				}
				if shapes[j].color != Empty {
					t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
						j, shapes[i].color, Empty, i)
				}
				if shapes[j].filled != 0 {
					t.Fatalf("Shape #%d has f%d, should be %d (iteration %d)\n",
						j, shapes[i].filled, 0, i)
				}
			}
		}
	}
	t.Log("Passed setting all shapes to Empty")

	if l > 2 {
		tiles[0].SetColor(Green)
		for i := 1; i < l; i++ {
			tiles[i].SetColor(Red)
		}
		for i := 1; i < l; i++ {
			if shapes[i].color != Mixed {
				t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
					i, shapes[i].color, Mixed)
			}
		}
		t.Log("Passed setting first Green, rest Red")

		tiles[0].SetColor(Blue)
		for i := 1; i < l; i++ {
			if shapes[i].color != Mixed {
				t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
					i, shapes[i].color, Mixed)
			}
		}
		t.Log("Passed setting first Blue, rest Red")

		tiles[0].SetColor(Empty)
		for i := 1; i < l; i++ {
			if shapes[i].color != Red {
				t.Fatalf("Shape #%d is c%d, should be %d (iteration %d)\n",
					i, shapes[i].color, Red)
			}
		}
		t.Log("Passed setting first Empty, rest Red")

	} else {
		t.Log("Skipped testing for mixed colors to Empty, as l is %d", l)
	}

}

func TestRegisterSquaresWithDistance(t *testing.T) {
	w, h := 4, 4
	if testing.Verbose() {
		w, h = 8, 8
	}

	board := MakeBoard(w, h)
	for i := 1; i < w && i < h; i++ {
		board.RegisterSquaresWithDistance(i)
		for j := 0; j < len(board.shapes); j++ {
			for k := 0; k < j; k++ {
				if board.shapes[j].ContainedIn(board.shapes[k]) {
					t.Fatalf("Shape #%d is contained in Shape #%d "+
						"(iteration %d)\n%v\n%v\n", j, k, i,
						board.shapes[j].nodes, board.shapes[k].nodes)
				}
			}
		}
		if testing.Verbose() {
			t.Logf("Board of size %d,%d has %d squares with SD%d\n",
				w, h, len(board.shapes), i)
		}
		board.ClearShapes()
	}

	board = MakeBoard(3, 3)
	board.RegisterSquaresWithDistance(2)
	if len(board.shapes) != 1 {
		t.Errorf("Board fails at registering courners in 3x3 (found %d)\n",
			len(board.shapes))
	}
	board.ClearShapes()
	board.RegisterSquaresWithDistance(1)
	if len(board.shapes) != 5 {
		t.Errorf("Board fails at registering 1x1s in 3x3 (%d != %d)\n",
			5, len(board.shapes))
	}

	if testing.Verbose() {
		for d := 0; d < 8; d++ {
			board = MakeBoard(d, d)
			for i := 1; i < d; i++ {
				board.RegisterSquaresWithDistance(i)
			}
			t.Logf("%d^2:\t%d squares\n", d, len(board.shapes))
		}
	}

}

/*func TestCopy(t *testing.T) {
	board := MakeBoard(3, 3)
	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			board.tiles[x][y].SetColor(Red)
		}
	}

	board2 := board.Copy()
	if board.w != board2.w || board.h != board2.h {
		t.Fatalf("Copy not same size: %d,%d != %d,%d\n",
			board.w, board.h, board2.w, board2.h)
	}

	for x := 0; x < board2.w; x++ {
		for y := 0; y < board2.h; y++ {
			if board2.tiles[x][y].color != Red {
				t.Fatal("Copy failed to copy color")
			}
		}
	}

	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			board.tiles[x][y].SetColor(Green)
		}
	}

	for x := 0; x < board2.w; x++ {
		for y := 0; y < board2.h; y++ {
			if board2.AtVector(&Vector{x, y}).color != Red {
				t.Fatal("Copy disconnect tiles")
			}
		}
	}

	for x := 0; x < board2.w; x++ {
		for y := 0; y < board2.h; y++ {
			board2.tiles[x][y].SetColor(Blue)
		}
	}

	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			if board.tiles[x][y].color != Green {
				t.Fatal("Copy disconnect tiles")
			}
		}
	}
}
*/
