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
	Tiles [][]Material

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
	dungeon.collectEdges()
	dungeon.connectRooms()
	dungeon.trimTunnels()
	dungeon.placeObjects()
	dungeon.exportTiles()
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

func (d *Dungeon) placeObjects() {
	if len(d.rooms) < 2 {
		panic("what a lame tiny dung you have")
	}

	var rA, rB Room
	maxAttempts := 10
	// retry it rooms are adjacent
	for i := 0; i < maxAttempts; i++ {
		shuffle := d.rand.Perm(len(d.rooms))
		rA, rB = d.rooms[shuffle[0]], d.rooms[shuffle[1]]
		if !rA.WithWall().Intersects(rB.WithWall().Rect) {
			break
		}
	}
	d.randAppendObject(rA, Player)
	d.randAppendObject(rB, Finish)
}

func (d *Dungeon) randAppendObject(r Room, object Material) {
	x := d.rand.Intn(r.Width()) + r.TL.X
	y := d.rand.Intn(r.Height()) + r.TL.Y
	d.tiles[y][x] = tile{material: object}
}

func (d *Dungeon) exportTiles() {
	tiles := make([][]Material, len(d.tiles))
	for i := range tiles {
		tiles[i] = make([]Material, len(d.tiles[i]))
		for j := range tiles[i] {
			tiles[i][j] = d.tiles[i][j].material
		}
	}
	d.Tiles = tiles
}
