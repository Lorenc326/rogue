package graphic

import (
	"rogue.game/core/maps"
	"rogue.game/core/player"
	"rogue.game/core/symbol"
)

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
