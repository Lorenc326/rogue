package dungeon

import "rogue.game/core/geo"

func (d *Dungeon) createMaze() {
	for x := 1; x < d.width-1; x++ {
		for y := 1; y < d.height-1; y++ {
			searchSubset := geo.NewRect(geo.Point{X: x - 1, Y: y - 1}, 3, 3)
			if every(d.tiles, searchSubset, Wall) {
				d.numRegions++
				d.continueMaze(x, y)
			}
		}
	}
}

func (d *Dungeon) continueMaze(x int, y int) {
	validTiles := []geo.Point{}

	// check each side for wall-only blocks
	if x-2 >= 0 {
		searchSubset := geo.NewRect(geo.Point{X: x - 2, Y: y - 1}, 2, 3)
		if every(d.tiles, searchSubset, Wall) {
			validTiles = append(validTiles, geo.Point{Y: y, X: x - 1})
		}
	}
	if x+2 < d.width {
		searchSubset := geo.NewRect(geo.Point{X: x + 1, Y: y - 1}, 2, 3)
		if every(d.tiles, searchSubset, Wall) {
			validTiles = append(validTiles, geo.Point{Y: y, X: x + 1})
		}
	}
	if y-2 >= 0 {
		searchSubset := geo.NewRect(geo.Point{X: x - 1, Y: y - 2}, 3, 2)
		if every(d.tiles, searchSubset, Wall) {
			validTiles = append(validTiles, geo.Point{Y: y - 1, X: x})
		}
	}
	if y+2 < d.height {
		searchSubset := geo.NewRect(geo.Point{X: x - 1, Y: y + 1}, 3, 2)
		if every(d.tiles, searchSubset, Wall) {
			validTiles = append(validTiles, geo.Point{Y: y + 1, X: x})
		}
	}

	if len(validTiles) < 1 {
		return
	}

	i := d.rand.Intn(len(validTiles))
	point := validTiles[i]

	d.tiles[point.Y][point.X].material = Tunnel
	d.tiles[point.Y][point.X].region = d.numRegions

	d.continueMaze(point.X, point.Y)
	d.continueMaze(x, y)
}

func (d *Dungeon) trimTunnels() {
	for x := 1; x < d.width-1; x++ {
		for y := 1; y < d.height-1; y++ {
			d.continueTrimTunnels(x, y)
		}
	}
}

func (d *Dungeon) continueTrimTunnels(x int, y int) {
	if d.tiles[y][x].material != Tunnel && d.tiles[y][x].material != Door {
		return
	}

	wallCount := 0
	nextPoint := geo.Point{}

	surroundingPoints := [4]geo.Point{
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
		{X: x, Y: y - 1},
		{X: x, Y: y + 1},
	}

	for i := range surroundingPoints {
		tile := d.tiles[surroundingPoints[i].Y][surroundingPoints[i].X]
		if tile.material == Wall {
			wallCount++
		} else if tile.material == Tunnel || tile.material == Door {
			nextPoint = geo.Point{X: surroundingPoints[i].X, Y: surroundingPoints[i].Y}
		}
	}

	if wallCount >= 3 {
		d.tiles[y][x].material = Wall
		d.tiles[y][x].region = 0
		if nextPoint.X != 0 || nextPoint.Y != 0 {
			d.continueTrimTunnels(nextPoint.X, nextPoint.Y)
		}
	}
}
