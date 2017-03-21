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
	
package AI

import (
	"github.com/ITR13/metasquares/engine"
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
