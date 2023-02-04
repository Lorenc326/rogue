package maps

import (
	_ "embed"
	"errors"

	"rogue.game/core/geo"
)

// go:embed templates/default.txt
// var defaultMapStr string

// go:embed templates/easy.txt
// var easyMapStr string

type Floor [][]string

func Read(name string) *Floor {
	// rows := strings.Split(defaultMapStr, "\n")
	// hoh := generator.RenderDungeon(generator.GenerateDungeon(35, 30))
	// log.Println(hoh)
	// rows := strings.Split(hoh, "\n")
	// res := make([][]string, 0, len(rows))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))

	// prefix := strings.Repeat("#", 5)
	// for i, rowStr := range rows {
	// 	hm := prefix
	// 	if i == len(rows)-1 {
	// 		continue
	// 	}
	// 	res = append(res, strings.Split(hm+rowStr+hm, ""))
	// }
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))
	// res = append(res, strings.Split(strings.Repeat("#", 45), ""))

	// Floor := Floor(res)
	return &Floor{}
}

func (m *Floor) Slice(tl geo.Point, br geo.Point) Floor {
	if tl.Y > br.Y || tl.X > br.X {
		panic(errors.New("invalid Coords, can't slice map"))
	}
	height := br.Y - tl.Y
	width := br.X - tl.X
	sliced := make([][]string, height)
	for i := 0; i < height; i++ {
		sliced[i] = make([]string, width)
		for j := 0; j < width; j++ {
			sliced[i][j] = (*m)[tl.Y+i][tl.X+j]
		}
	}
	return Floor(sliced)
}

func (m *Floor) Insert(c geo.Point, symbol string) {
	(*m)[c.Y][c.X] = symbol
}

func (m *Floor) Find(symbol string) *geo.Point {
	for i, row := range *m {
		for j, col := range row {
			if col == symbol {
				return &geo.Point{Y: i, X: j}
			}
		}
	}
	return nil
}

func (m *Floor) Replace(c geo.Point, new string) {
	(*m)[c.Y][c.X] = new
}

// map template HAS to be with offset spaces in template (awful huck haha)
func (m *Floor) SliceCentered(c geo.Point, offset int) Floor {
	return m.Slice(
		geo.Point{Y: c.Y - offset, X: c.X - offset},
		geo.Point{Y: c.Y + offset + 1, X: c.X + offset + 1},
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
