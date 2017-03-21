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
	
package main

import (
	"fmt"

	"github.com/ITR13/metasquares/AI"
	"github.com/ITR13/metasquares/animator"
	"github.com/ITR13/metasquares/engine"
)

func main() {
	winners := [9]int{}
	for w := 6; w < 7; w++ {
		for h := w; h == w; h++ {
			game := Engine.MakeGame(w, h,
				&AI.MixedTaker{},
				AI.GetBruteForcer(false, 2, nil, &AI.MixedTaker{}),
			)
			aniBoard := Engine.MakeBoard(w, h)
			game.Animator = Animators.Text{true, true, aniBoard}
			//			game.Animator = nil
			for i := 1; i < w && i < h; i++ {
				aniBoard.RegisterSquaresWithDistance(i)
			}
			winner := game.Play()
			fmt.Printf("%d Won %dx%d!\n", winner, w, h)
			winners[winner-1]++
		}
	}
	fmt.Println(winners)
}
