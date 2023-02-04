package dungeon

import (
	"math/rand"
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
	dungeon.createMaze()
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
