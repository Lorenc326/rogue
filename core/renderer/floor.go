package renderer

import (
	"strings"

	"rogue.game/core/maps"
	"rogue.game/core/player"
	"rogue.game/core/symbol"
)

func ASCII(m *maps.FloorMap) string {
	rows := make([]string, len(*m))
	for i, row := range *m {
		rows[i] = strings.Join(row, "")
	}
	return strings.Join(rows, "\n")
}

func ASCIIWide(m *maps.FloorMap) string {
	rows := make([]string, len(*m))
	for i, row := range *m {
		rows[i] = strings.Join(row, " ")
	}
	return " " + strings.Join(rows, "\n")
}

// map template HAS to be with offset spaces in template (awful huck haha)
func PlayerCenteredMap(m *maps.FloorMap, p *player.Player, offset int) *maps.FloorMap {
	newMap := m.Slice(
		maps.Coord{I: p.Coord.I - offset, J: p.Coord.J - offset},
		maps.Coord{I: p.Coord.I + offset + 1, J: p.Coord.J + offset + 1},
	)
	// user is not present on map before render for flexibility
	newMap.Insert(maps.Coord{I: offset, J: offset}, symbol.Player)
	return &newMap
}
