package player

import (
	"errors"

	"rogue.game/core/action"
	"rogue.game/core/maps"
	"rogue.game/core/symbol"
)

type Player struct {
	Coord maps.Coord
}

func (u *Player) Move(m *maps.FloorMap, direction string) error {
	switch direction {
	case action.Up:
		u.Coord.I--
	case action.Right:
		u.Coord.J++
	case action.Down:
		u.Coord.I++
	case action.Left:
		u.Coord.J--
	default:
		return errors.New("not defined")
	}
	return u.validateDestination(m)
}

func (u *Player) Victory(m *maps.FloorMap) bool {
	location := (*m)[u.Coord.I][u.Coord.J]
	switch location {
	case symbol.End:
		return true
	}
	return false
}

func (u *Player) validateDestination(m *maps.FloorMap) error {
	location := (*m)[u.Coord.I][u.Coord.J]
	switch location {
	case symbol.Wall, symbol.Lava:
		return errors.New("ahhhhh")
	}
	return nil
}
