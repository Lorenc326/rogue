package player

import (
	"errors"

	"github.com/Lorenc326/rogue/core/dungeon"
	"github.com/Lorenc326/rogue/core/geo"
	"github.com/Lorenc326/rogue/core/maps"
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
