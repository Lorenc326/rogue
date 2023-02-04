package player

import (
	"errors"

	"rogue.game/core/geo"
	"rogue.game/core/maps"
	"rogue.game/core/symbol"
)

type Player struct {
	geo.Point
}

func (u *Player) Victory(m *maps.Floor) bool {
	location := (*m)[u.Point.Y][u.Point.X]
	switch location {
	case symbol.End:
		return true
	}
	return false
}

func (u *Player) ValidateDestination(m maps.Floor, c geo.Point) error {
	location := m[c.Y][c.X]
	switch location {
	case symbol.Wall, symbol.Lava:
		return errors.New("ahhhhhmmm nice")
	}
	return nil
}
