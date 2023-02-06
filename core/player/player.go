package player

import (
	"errors"

	"rogue.game/core/dungeon"
	"rogue.game/core/geo"
	"rogue.game/core/maps"
)

type Player struct {
	geo.Point
}

func (u *Player) Victory(m maps.Floor) bool {
	location := m[u.Point.Y][u.Point.X]
	switch location {
	case dungeon.Finish:
		return true
	}
	return false
}

func (u *Player) ValidateDestination(m maps.Floor, c geo.Point) error {
	location := m[c.Y][c.X]
	switch location {
	case dungeon.Wall, dungeon.Lava:
		return errors.New("ahhhhhmmm nice")
	}
	return nil
}
