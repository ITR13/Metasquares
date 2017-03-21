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
	"fmt"
	"io"

	"github.com/ITR13/metasquares/engine"
)

type PlayerC struct {
	Reader io.Reader
}

func (p *PlayerC) Init(w, h, players int) {}

func (p *PlayerC) Place(color Engine.Color,
	board Engine.BoardAbstract) (int, int) {
	var a, b int
	fmt.Fscanf(p.Reader, "%d %d\n", &a, &b)
	return a - 1, b - 1
}

func (p *PlayerC) Placed(x, y int, c Engine.Color) {}
