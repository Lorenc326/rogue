package graphic

import (
	_ "embed"
	"strings"

	"rogue.game/core/geo"
	"rogue.game/core/maps"
	"rogue.game/core/session"
	"rogue.game/core/symbol"
)

//go:embed assets/victory.txt
var vistoryStr string

type vision = func(src maps.Floor, c geo.Point, offset int) maps.Floor

type ascii struct {
	intend string
	offset int
	vision vision
}

func NewASCII(offset int, centered, wide bool) ascii {
	r := ascii{vision: fullVision, offset: offset}
	if wide {
		r.intend = " "
	}
	if centered {
		r.vision = centeredVision
	}
	return r
}

func (g ascii) Render(c session.DrawContext) string {
	if c.IsEnded {
		return vistoryStr
	}
	m := c.Floor.Clone()
	m.Insert(c.Player.Point, symbol.Player)
	m = g.vision(m, c.Player.Point, g.offset)
	return mapToString(&m, g.intend)
}

func fullVision(src maps.Floor, _ geo.Point, _ int) maps.Floor {
	return src
}

func centeredVision(src maps.Floor, c geo.Point, offset int) maps.Floor {
	return src.SliceCentered(c, offset)
}

func mapToString(m *maps.Floor, intend string) string {
	rows := make([]string, len(*m))
	for i, row := range *m {
		rows[i] = strings.Join(row, intend)
	}
	return strings.Join(rows, "\n")
}
