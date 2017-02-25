package Engine

type Vector struct {
	x, y int
}

func (vector *Vector) Add(other *Vector) {
	vector.x += other.x
	vector.y += other.y
}

func (vector *Vector) Rotate90() {
	vector.x, vector.y = -vector.y, vector.x
}

func (vector *Vector) IterateBorder() bool {
	if vector.y < 0 || vector.x <= 0 {
		return false
	}
	if vector.y < vector.x {
		vector.y++
	} else {
		vector.x--
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
	return board.At(pos.x, pos.y)
}

func (vector *Vector) Square() int {
	//TODO Use root rounded up instead
	return vector.x*vector.x + vector.y*vector.y
}
