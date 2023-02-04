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

func (d *Dungeon) createRooms() {
	var rooms []Room

	for i := 0; i < d.roomAttempts; i++ {
		width := d.rand.Intn(d.maxRoomSize-d.minRoomSize) + d.minRoomSize
		height := d.rand.Intn(d.maxRoomSize-d.minRoomSize) + d.minRoomSize

		maxX := d.width - width - 2
		maxY := d.height - height - 2

		tl := geo.Point{
			X: d.rand.Intn(maxX-3) + 3,
			Y: d.rand.Intn(maxY-3) + 3,
		}

		room := Room{Rect: geo.NewRect(tl, width, height)}
		walledRoom := room.WithWall()

		shouldAppend := true
		for _, existingRoom := range rooms {
			if walledRoom.Intersects(existingRoom.Rect) {
				shouldAppend = false
				break
			}
		}
		if shouldAppend {
			rooms = append(rooms, room)
		}
	}

	for _, r := range rooms {
		d.numRegions++
		for i := r.TL.X; i <= r.BR.X; i++ {
			for j := r.TL.Y; j <= r.BR.Y; j++ {
				d.tiles[j][i].material = Floor
				d.tiles[j][i].region = d.numRegions
			}
		}
	}

	d.rooms = rooms
}
