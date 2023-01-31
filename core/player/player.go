package player

import (
	"errors"

	"rogue.game/core/geom"
	"rogue.game/core/maps"
	"rogue.game/core/symbol"
)

type Player struct {
	geom.Coord
}

func (u *Player) Victory(m *maps.Floor) bool {
	location := (*m)[u.Coord.I][u.Coord.J]
	switch location {
	case symbol.End:
		return true
	}
	return false
}

func (u *Player) ValidateDestination(m maps.Floor, c geom.Coord) error {
	location := m[c.I][c.J]
	switch location {
	case symbol.Wall, symbol.Lava:
		return errors.New("ahhhhhmmm nice")
	}
	return nil
}
