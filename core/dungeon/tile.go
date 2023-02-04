package dungeon

type Material int

const (
	Wall Material = iota
	Floor
	Door
	Tunnel
	Player
	Finish
)

type tile struct {
	region   int
	material Material
}

func fillTiles(width, height int) [][]tile {
	tiles := make([][]tile, height)
	for i := range tiles {
		tiles[i] = make([]tile, width)
	}
	return tiles
}
