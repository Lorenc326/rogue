package geo

import (
	"errors"

	"github.com/Lorenc326/rogue/core/event"
)

type Point struct {
	Y, X int
}

type Distance = Point

func (p *Point) Move(direction string, step int) error {
	switch direction {
	case event.Up:
		p.Y -= step
	case event.Right:
		p.X += step
	case event.Down:
		p.Y += step
	case event.Left:
		p.X -= step
	default:
		return errors.New("not defined")
	}
	return nil
}

func (p Point) Equal(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}

func (p Point) Add(d Distance) Point {
	return Point{X: p.X + d.X, Y: p.Y + d.Y}
}

func (p Point) Sub(d Distance) Point {
	return Point{X: p.X - d.X, Y: p.Y - d.Y}
}
