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
	return a, b
}

func (p *PlayerC) Placed(x, y int, c Engine.Color) {}
