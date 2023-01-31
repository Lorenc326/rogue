package geom

import (
	"errors"

	"rogue.game/core/event"
)

type Coord struct {
	I, J int // much clearer what is expected then with x, y
}

func (c *Coord) Move(direction string, step int) error {
	switch direction {
	case event.Up:
		c.I -= step
	case event.Right:
		c.J += step
	case event.Down:
		c.I += step
	case event.Left:
		c.J -= step
	default:
		return errors.New("not defined")
	}
	return nil
}
