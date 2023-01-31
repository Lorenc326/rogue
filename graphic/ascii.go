package graphic

import (
	"strings"

	"rogue.game/core/maps"
	"rogue.game/core/session"
	"rogue.game/core/symbol"
)

type vision = func(src maps.Floor, c maps.Coord, offset int) maps.Floor

type ascii struct {
	intend string
	offset int
	vision vision
}

func NewASCII(offset int, centered, wide bool) ascii {
	r := ascii{vision: fullVision}
	if wide {
		r.intend = " "
	}
	if centered {
		r.vision = centeredVision
	}
	return r
}

func (g ascii) Render(c session.DrawContext) string {
	m := c.Floor.Clone()
	m.Insert(c.Player.Coord, symbol.Player)
	m = g.vision(m, c.Player.Coord, g.offset)
	return mapToString(&m, g.intend)
}

func fullVision(src maps.Floor, _ maps.Coord, _ int) maps.Floor {
	return src
}

func centeredVision(src maps.Floor, c maps.Coord, offset int) maps.Floor {
	return src.SliceCentered(c, offset)
}

func mapToString(m *maps.Floor, intend string) string {
	rows := make([]string, len(*m))
	for i, row := range *m {
		rows[i] = strings.Join(row, intend)
	}
	return strings.Join(rows, "\n")
}
