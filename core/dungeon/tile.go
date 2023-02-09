package dungeon

import "github.com/Lorenc326/rogue/core/geo"

type Material int

const (
	Wall Material = iota
	Floor
	Door
	Tunnel
	Player
	Finish
	Lava
	Rat
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

func every(tiles [][]tile, subset geo.Rect, m Material) bool {
	for i := subset.TL.Y; i <= subset.BR.Y; i++ {
		for j := subset.TL.X; j <= subset.BR.X; j++ {
			if tiles[i][j].material != m {
				return false
			}
		}
	}
	return true
}
