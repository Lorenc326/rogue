package maps

import (
	_ "embed"
	"errors"

	"rogue.game/core/dungeon"
	"rogue.game/core/geo"
)

type Floor [][]dungeon.Material

func (m *Floor) Slice(tl geo.Point, br geo.Point) Floor {
	if tl.Y > br.Y || tl.X > br.X {
		panic(errors.New("invalid Coords, can't slice map"))
	}
	height := br.Y - tl.Y
	width := br.X - tl.X
	sliced := make([][]dungeon.Material, height)
	for i := 0; i < height; i++ {
		sliced[i] = make([]dungeon.Material, width)
		for j := 0; j < width; j++ {
			sliced[i][j] = (*m)[tl.Y+i][tl.X+j]
		}
	}
	return Floor(sliced)
}

func (m *Floor) Insert(c geo.Point, symbol dungeon.Material) {
	(*m)[c.Y][c.X] = symbol
}

func (m *Floor) Find(symbol dungeon.Material) *geo.Point {
	for i, row := range *m {
		for j, col := range row {
			if col == symbol {
				return &geo.Point{Y: i, X: j}
			}
		}
	}
	return nil
}

func (m *Floor) Replace(c geo.Point, new dungeon.Material) {
	(*m)[c.Y][c.X] = new
}

func (m *Floor) SliceCentered(c geo.Point, offset int) Floor {
	return m.Slice(
		geo.Point{Y: c.Y - offset, X: c.X - offset},
		geo.Point{Y: c.Y + offset + 1, X: c.X + offset + 1},
	)
}

func (m *Floor) Clone() Floor {
	height := len(*m)
	sliced := make([][]dungeon.Material, height)
	for i := 0; i < height; i++ {
		width := len((*m)[i])
		sliced[i] = make([]dungeon.Material, width)
		copy(sliced[i], (*m)[i])
	}
	return Floor(sliced)
}
