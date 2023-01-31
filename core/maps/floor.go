package maps

import (
	_ "embed"
	"errors"
	"strings"

	"rogue.game/core/geom"
)

//go:embed templates/default.txt
var defaultMapStr string

// go:embed templates/easy.txt
// var easyMapStr string

type Floor [][]string

func Read(name string) *Floor {
	rows := strings.Split(defaultMapStr, "\n")
	res := make([][]string, 0, len(rows))
	for _, rowStr := range rows {
		res = append(res, strings.Split(rowStr, ""))
	}

	Floor := Floor(res)
	return &Floor
}

func (m *Floor) Slice(tl geom.Coord, br geom.Coord) Floor {
	if tl.I > br.I || tl.J > br.J {
		panic(errors.New("invalid Coords, can't slice map"))
	}
	height := br.I - tl.I
	width := br.J - tl.J
	sliced := make([][]string, height)
	for i := 0; i < height; i++ {
		sliced[i] = make([]string, width)
		for j := 0; j < width; j++ {
			sliced[i][j] = (*m)[tl.I+i][tl.J+j]
		}
	}
	return Floor(sliced)
}

func (m *Floor) Insert(c geom.Coord, symbol string) {
	(*m)[c.I][c.J] = symbol
}

func (m *Floor) Find(symbol string) *geom.Coord {
	for i, row := range *m {
		for j, col := range row {
			if col == symbol {
				return &geom.Coord{I: i, J: j}
			}
		}
	}
	return nil
}

func (m *Floor) Replace(c geom.Coord, new string) {
	(*m)[c.I][c.J] = new
}

// map template HAS to be with offset spaces in template (awful huck haha)
func (m *Floor) SliceCentered(c geom.Coord, offset int) Floor {
	return m.Slice(
		geom.Coord{I: c.I - offset, J: c.J - offset},
		geom.Coord{I: c.I + offset + 1, J: c.J + offset + 1},
	)
}

func (m *Floor) Clone() Floor {
	height := len(*m)
	sliced := make([][]string, height)
	for i := 0; i < height; i++ {
		width := len((*m)[i])
		sliced[i] = make([]string, width)
		copy(sliced[i], (*m)[i])
	}
	return Floor(sliced)
}
