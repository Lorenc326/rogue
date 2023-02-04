package dungeon

import (
	"math/rand"

	"rogue.game/core/geo"
)

const (
	defaultRoomAttempts = 200
	defaultMinRoomSize  = 5
	defaultMaxRoomSize  = 15
)

type Dungeon struct {
	tiles  [][]tile
	rooms  []Room
	width  int
	height int

	numRegions   int
	minRoomSize  int
	maxRoomSize  int
	roomAttempts int

	seed int64
	rand rand.Rand
}

func Generate(seed int64, width, height int) *Dungeon {
	dungeon := new(seed, width, height)
	dungeon.createRooms()
	return dungeon
}

func new(seed int64, width, height int) *Dungeon {
	return &Dungeon{
		tiles:        fillTiles(width, height),
		rooms:        []Room{},
		width:        width,
		height:       height,
		seed:         seed,
		rand:         *rand.New(rand.NewSource(seed)),
		minRoomSize:  defaultMinRoomSize,
		maxRoomSize:  defaultMaxRoomSize,
		roomAttempts: defaultRoomAttempts,
	}
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
