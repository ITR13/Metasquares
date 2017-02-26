package Engine

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ox, oy := 0, 0
	v := &Vector{0, 0}
	for x := -5; x <= 5; x++ {
		for y := -5; y <= 5; y++ {
			v.Add(&Vector{x, y})
			ox, oy = ox+x, oy+y
			if ox != v.X || oy != v.Y {
				t.Fatalf("Failed adding %d,%d to %d,%d (got %d,%d)\n",
					x, y, ox-x, oy-y, v.X, v.Y)
			}
		}
	}

}

func TestRotate90(t *testing.T) {
	for x := -5; x <= 5; x++ {
		for y := -5; y <= 5; y++ {
			a, b := &Vector{x, y}, &Vector{x, y}
			b.Rotate90()
			c := &Vector{b.X, b.Y}
			b.Rotate90()
			if a.X != -b.X || a.Y != -b.Y {
				t.Fatalf("180: %d,%d != %d,%d\n", -a.X, -a.Y, b.X, b.Y)
			}
			b.Rotate90()
			if c.X != -b.X || c.Y != -b.Y {
				t.Fatalf("270: %d,%d != %d,%d\n", -c.X, -c.Y, b.X, b.Y)
			}
			b.Rotate90()
			if a.X != b.X || a.Y != b.Y {
				t.Fatalf("360: %d,%d != %d,%d\n", a.X, a.Y, b.X, b.Y)
			}
		}
	}

	one := &Vector{1, 0}
	one.Rotate90()
	if one.X != 0 || one.Y != 1 {
		t.Fatal("Rotating the wrong direction (not that it doesn't matter)\n")
	}

}

func TestIterateBorder(t *testing.T) {
	for x := -5; x <= 5; x++ {
		for y := -5; y <= 5; y++ {
			v := &Vector{x, y}
			if v.IterateBorder() != (x > 0 && y >= 0) {
				t.Fatalf("Wrong return when iterating %d,%d\n", x, y)
			}
		}
	}

	for d := 1; d < 5; d++ {
		v := &Vector{d, 0}
		for y := 0; y < d; y++ {
			if v.X != d || v.Y != y {
				t.Errorf("Failed when iterating to %d,%d: %v\n", d, y, v)
				v.X, v.Y = d, y
			}
			if !v.IterateBorder() {
				t.Errorf("Early iteration error (%d,%d): %v\n", d, y, v)
			}
			if testing.Verbose() {
				t.Logf("EV: %d,%d: %v\n", d, y, v)
			}
		}
		for x := d; x >= 0; x-- {
			if v.X != x || v.Y != d {
				t.Errorf("Failed when iterating to %d,%d: %v\n", x, d, v)
				v.X, v.Y = x, d
			}
			if v.IterateBorder() == (x == 0) {
				t.Errorf("Late iteration error (%d,%d): %v\n", d, x, v)
			}
			if testing.Verbose() {
				t.Logf("LV: %d,%d: %v\n", x, d, v)
			}
		}
	}
	if !t.Failed() {
		t.Log("Succeded move-tests")
	}

	board := MakeBoard(5, 5)
	origin := Vector{2, 2}
	colored := 0
	for i := 1; i < 3; i++ {
		move := &Vector{i, 0}
		for j := 0; j < 2*i; j++ {
			if !move.IterateBorder() {
				t.Errorf("Early finish on %d (%d left)", i, 2*i-j)
				break
			}
			for k := 0; k < 4; k++ {
				origin2 := origin
				origin2.Add(move)
				tile := board.AtVector(&origin2)
				move.Rotate90()
				if tile == nil {
					t.Fatalf("Failed to get tile %v + %v (%v)\n",
						origin, move, origin2)
				} else {
					tile.SetColor(Red)
					colored++
					if CountColored(board) != colored {
						t.Fatalf("tile at %v was already colored (itr: %d,%d)\n",
							origin2, i, k)
					}
				}
			}
		}
	}

	if board.AtVector(&origin).color != Empty {
		t.Fatal("Something colored the middle tile\n")
	}
	board.AtVector(&origin).SetColor(Red)
	colored = CountColored(board)
	if colored != 25 {
		t.Fatalf("Failed to color all tiles (colored %d/25)\n", colored)
	}
	t.Log("Passed coloring test")
}

func CountColored(board *Board) int {
	c := 0
	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			if board.tiles[x][y].color != Empty {
				c++
			}
		}
	}
	return c
}

func TestGetTiles(t *testing.T) {
	diag := &Vector{1, 1}
	board := MakeBoard(3, 3)
	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			origin := &Vector{x, y}
			diamond := board.GetTiles(origin, diag)
			if (diamond == nil) == (x == 1 && y == 0) {
				t.Errorf("Error when getting diamond-shape (%d,%d)\n", x, y)
			}
			if origin.X != x || origin.Y != y {
				t.Fatalf("Origin not back at correct position after finishing"+
					" (%d,%d != %d,%d)\n", x, y, origin.X, origin.Y)
			}
			if diag.X != 1 || diag.Y != 1 {
				t.Fatalf("Diagonal not back at correct position after "+
					"finishing (%d,%d != %d,%d)\n", x, y, diag.X, diag.Y)
			}
		}
	}
}

func TestAt(t *testing.T) {
	board := MakeBoard(5, 5)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			v := &Vector{x, y}
			if board.AtVector(v) != board.tiles[x][y] {
				t.Fatalf("Wrong node gotten at %d,%d\n", x, y)
			}
		}
	}
}

func TestSquare(t *testing.T) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			v := &Vector{x, y}
			if v.Square() != x*x+y*y {
				t.Fatalf("%d*%d != %v^2", x, y, v)
			}
		}
	}
}
