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

type Color uint8

const (
	Empty   Color = iota
	Red     Color = iota
	Green   Color = iota
	Blue    Color = iota
	Cyan    Color = iota
	Magenta Color = iota
	Yellow  Color = iota
	Black   Color = iota
	White   Color = iota

	Mixed     Color = iota
	NotAColor Color = iota
)

type Board struct {
	w, h   int
	tiles  [][]*Tile
	shapes []*Shape
}

type Shape struct {
	nodes  []*Tile
	worth  int
	color  Color
	filled int
}

type Tile struct {
	color  Color
	x, y   int
	shapes []*Shape
}

func (tile *Tile) SetColor(c Color) {
	if tile.color != c {
		if tile.color == Empty {
			tile.color = c
			for i := 0; i < len(tile.shapes); i++ {
				shape := tile.shapes[i]
				shape.filled++
				if shape.color == Empty {
					shape.color = c
				} else if shape.color != c {
					shape.color = Mixed
				}
			}
		} else {
			tile.color = c
			for i := 0; i < len(tile.shapes); i++ {
				shape := tile.shapes[i]
				if c == Empty {
					shape.filled--
				}
				shape.color = shape.nodes[0].color
				for j := 1; j < len(shape.nodes); j++ {
					color := shape.nodes[j].color
					if color == Empty {
						continue
					} else if shape.color == Empty {
						shape.color = color
					} else if shape.color != color {
						shape.color = Mixed
						break
					}
				}
			}
		}
	}
}

func MakeBoard(w, h int) *Board {
	tiles := make([][]*Tile, w)
	for x := 0; x < w; x++ {
		tiles[x] = make([]*Tile, h)
		for y := 0; y < h; y++ {
			tiles[x][y] = &Tile{Empty, x, y, make([]*Shape, 0)}
		}
	}
	return &Board{w, h, tiles, make([]*Shape, 0)}
}

func (board *Board) ClearShapes() {
	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			board.tiles[x][y].shapes = make([]*Shape, 0)
		}
	}
	board.shapes = make([]*Shape, 0)
}

func (board *Board) RegisterSquaresWithDistance(n int) {
	if n < 1 {
		panic("Tried to register squares with distance<1")
	} else if n >= board.w || n >= board.h {
		panic("Tried to register squares that don't fit")
	}

	for x := 0; x < board.w; x++ {
		for y := 0; y < board.h; y++ {
			origin, move := &Vector{x, y}, &Vector{n, 0}
			for move.IterateBorder() {
				tiles := board.GetTiles(origin, move)
				if tiles != nil {
					board.MakeShape(*tiles, move.Square())
				}
			}
		}
	}

}

func (board *Board) MakeShape(tiles []*Tile, worth int) {
	shape := &Shape{tiles, worth, Empty, 0}
	board.shapes = append(board.shapes, shape)
	for i := 0; i < len(tiles); i++ {
		tiles[i].shapes = append(tiles[i].shapes, shape)
	}
}

func (shape *Shape) ContainedIn(other *Shape) bool {
	for i := 0; i < len(shape.nodes); i++ {
		unique := true
		for j := 0; j < len(other.nodes); j++ {
			if shape.nodes[i] == other.nodes[j] {
				unique = false
				break
			}
		}
		if unique {
			return false
		}
	}
	return true
}

func (board *Board) At(x, y int) *Tile {
	if x == -1 && y == -1 {
		return &Tile{Empty, -1, -1, make([]*Shape, 0)}
	}

	if x < 0 || y < 0 || x >= board.w || y >= board.h {
		return nil
	}
	return board.tiles[x][y]
}

func (board *Board) GetSize() (int, int) {
	return board.w, board.h
}

func (board *Board) GetLeader() Color {
	scores := board.GetScores()
	winner, score := Empty, 0
	for i := Empty + 1; i < Mixed; i++ {
		if scores[i] > score {
			winner, score = Color(i), scores[i]
		} else if scores[i] == score {
			winner = Mixed
		}
	}
	return winner
}
func (board *Board) GetScores() []int {
	scores := make([]int, NotAColor+1)
	for i := 0; i < len(board.shapes); i++ {
		shape := board.shapes[i]
		if shape.filled == len(shape.nodes) {
			scores[shape.color] += shape.worth
		}
	}
	return scores
}
