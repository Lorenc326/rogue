package dungeon

import "rogue.game/core/geo"

type Room struct {
	geo.Rect
}

// offset to calc room size including walls
var wallOffset = geo.Distance{X: 1, Y: 1}

// Add wall ofset to the room rect
func (r Room) WithWall() Room {
	r.Rect = geo.Rect{TL: r.TL.Sub(wallOffset), BR: r.BR.Add(wallOffset)}
	return r
}
