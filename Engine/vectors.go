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

type Vector struct {
	X, Y int
}

func (vector *Vector) Add(other *Vector) {
	vector.X += other.X
	vector.Y += other.Y
}

func (vector *Vector) Rotate90() {
	vector.X, vector.Y = -vector.Y, vector.X
}

func (vector *Vector) IterateBorder() bool {
	if vector.Y < 0 || vector.X <= 0 {
		return false
	}
	if vector.Y < vector.X {
		vector.Y++
	} else {
		vector.X--
	}
	return true
}

func (board *Board) GetTiles(origin, move *Vector) *[]*Tile {
	//TODO Consider checking for nil before making tiles, or using premade
	o2, m2 := *origin, *move
	tiles := make([]*Tile, 4)
	for i := 0; i < 4; i++ {
		(&o2).Add(&m2)
		tiles[i] = board.AtVector(&o2)
		if tiles[i] == nil {
			return nil
		}
		m2.Rotate90()
	}
	return &tiles
}

func (board *Board) AtVector(pos *Vector) *Tile {
	return board.At(pos.X, pos.Y)
}

func (vector *Vector) Square() int {
	//TODO Use root rounded up instead
	return vector.X*vector.X + vector.Y*vector.Y
}
