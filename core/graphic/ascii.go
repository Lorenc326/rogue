package graphic

import (
	"strings"

	"rogue.game/core/maps"
	"rogue.game/core/session"
)

type ascii struct {
	centered bool
	intend   string
	offset   int
}

func NewASCII(offset int, centered, wide bool) *ascii {
	r := &ascii{offset: offset, centered: centered}
	if wide {
		r.intend = " "
	}
	return r
}

func (g *ascii) Render(c session.DrawContext) string {
	m := c.Floormap
	if g.centered {
		m = PlayerCenteredMap(m, c.Player, g.offset)
	}
	return mapToString(m, g.intend)
}

func mapToString(m *maps.FloorMap, intend string) string {
	rows := make([]string, len(*m))
	for i, row := range *m {
		rows[i] = strings.Join(row, intend)
	}
	return intend + strings.Join(rows, "\n")
}
