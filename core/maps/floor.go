package maps

import (
	_ "embed"
	"errors"
	"strings"
)

//go:embed templates/default.txt
var defaultMapStr string

// go:embed templates/easy.txt
// var easyMapStr string

type Coord struct {
	I, J int // much clearer what is expected then with x, y
}

type FloorMap [][]string

func Read(name string) *FloorMap {
	rows := strings.Split(defaultMapStr, "\n")
	res := make([][]string, 0, len(rows))
	for _, rowStr := range rows {
		res = append(res, strings.Split(rowStr, ""))
	}

	floormap := FloorMap(res)
	return &floormap
}

func (m *FloorMap) Slice(tl Coord, br Coord) FloorMap {
	if tl.I > br.I || tl.J > br.J {
		panic(errors.New("invalid coords, can't slice map"))
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
	return FloorMap(sliced)
}

func (m *FloorMap) Insert(c Coord, symbol string) {
	(*m)[c.I][c.J] = symbol
}

func (m *FloorMap) Find(symbol string) *Coord {
	for i, row := range *m {
		for j, col := range row {
			if col == symbol {
				return &Coord{I: i, J: j}
			}
		}
	}
	return nil
}

func (m *FloorMap) Replace(c Coord, new string) {
	(*m)[c.I][c.J] = new
}
