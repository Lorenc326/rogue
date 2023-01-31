package player

import (
	"errors"

	"rogue.game/core/event"
	"rogue.game/core/maps"
	"rogue.game/core/symbol"
)

type Player struct {
	Coord maps.Coord
}

func (u *Player) Move(m *maps.Floor, direction string, step int) error {
	switch direction {
	case event.Up:
		u.Coord.I -= step
	case event.Right:
		u.Coord.J += step
	case event.Down:
		u.Coord.I += step
	case event.Left:
		u.Coord.J -= step
	default:
		return errors.New("not defined")
	}
	return u.validateDestination(m)
}

func (u *Player) Victory(m *maps.Floor) bool {
	location := (*m)[u.Coord.I][u.Coord.J]
	switch location {
	case symbol.End:
		return true
	}
	return false
}

func (u *Player) validateDestination(m *maps.Floor) error {
	location := (*m)[u.Coord.I][u.Coord.J]
	switch location {
	case symbol.Wall, symbol.Lava:
		return errors.New("ahhhhh")
	}
	return nil
}
