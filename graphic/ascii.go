package graphic

import (
	_ "embed"
	"strings"

	"rogue.game/core/dungeon"
	"rogue.game/core/geo"
	"rogue.game/core/maps"
	"rogue.game/core/session"
	"rogue.game/graphic/symbol"
)

//go:embed assets/victory.txt
var vistoryStr string

type ascii struct {
	intend   string
	centered bool
	offset   geo.Distance
}

func NewASCII(offset int, centered, wide bool) ascii {
	r := ascii{centered: centered, offset: geo.Distance{X: offset, Y: offset}}
	if wide {
		r.intend = " "
	}
	return r
}

func (g ascii) Render(c session.DrawContext) string {
	if c.IsEnded {
		return vistoryStr
	}

	m := c.Floor.Clone()
	m.Insert(c.Player.Point, dungeon.Player)

	var visionRect geo.Rect
	if g.centered {
		size := g.offset.X*2 + 1
		visionRect = geo.NewRect(c.Player.Point.Sub(g.offset), size, size)
	} else {
		visionRect = geo.NewRect(geo.Point{X: 0, Y: 0}, len(c.Floor[0]), len(c.Floor))
	}

	return rectToString(m, visionRect, g.intend)
}

func rectToString(m maps.Floor, r geo.Rect, intend string) string {
	builder := strings.Builder{}
	for y := r.TL.Y; y <= r.BR.Y; y++ {
		for x := r.TL.X; x <= r.BR.X; x++ {
			var mat dungeon.Material
			// if user is at the edge of the map and he has centered view
			// its expected to receive negative coords so we can prefill empty space
			if x < 0 || y < 0 {
				mat = dungeon.Wall
			} else {
				mat = m[y][x]
			}
			builder.WriteString(symbol.MaterialToSymbol[mat] + intend)
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
