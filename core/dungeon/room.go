package dungeon

import (
	"rogue.game/core/geo"
)

type Room struct {
	geo.Rect
	edges []geo.Point
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

func (d *Dungeon) collectEdges() {
	for i := range d.rooms {
		rs := d.rooms

		for x := rs[i].TL.X; x <= rs[i].BR.X; x++ {
			topTile := d.tiles[rs[i].TL.Y-2][x]
			if topTile.material == Tunnel || topTile.material == Floor {
				rs[i].edges = append(rs[i].edges, geo.Point{X: x, Y: rs[i].TL.Y - 1})
			}
			bottomTile := d.tiles[rs[i].BR.Y+2][x]
			if bottomTile.material == Tunnel || bottomTile.material == Floor {
				rs[i].edges = append(rs[i].edges, geo.Point{X: x, Y: rs[i].BR.Y + 1})
			}
		}

		for y := rs[i].TL.Y; y <= rs[i].BR.Y; y++ {
			leftTile := d.tiles[y][rs[i].TL.X-2]
			if leftTile.material == Tunnel || leftTile.material == Floor {
				rs[i].edges = append(rs[i].edges, geo.Point{X: rs[i].TL.X - 1, Y: y})
			}
			rightTile := d.tiles[y][rs[i].BR.X+2]
			if rightTile.material == Tunnel || rightTile.material == Floor {
				rs[i].edges = append(rs[i].edges, geo.Point{X: rs[i].BR.X + 1, Y: y})
			}
		}
	}
}

func (d *Dungeon) connectRooms() {
	for i := range d.rooms {
		r := &d.rooms[i]
		edge := r.edges[d.rand.Intn(len(r.edges))]
		roomRegion := d.tiles[r.TL.Y][r.TL.X].region
		edgeBox := geo.NewRect(geo.Point{X: edge.X - 1, Y: edge.Y - 1}, 3, 3).Ring()

		for _, point := range edgeBox {
			tile := d.tiles[point.Y][point.X]
			if (tile.material == Floor || tile.material == Tunnel) && tile.region != roomRegion {

				d.tiles[edge.Y][edge.X].material = Door
				for x := r.TL.X; x <= r.BR.X; x++ {
					for y := r.TL.Y; y <= r.BR.Y; y++ {
						d.tiles[y][x].region = tile.region
					}
				}
				break
			}
		}
	}

	connectedRegions := map[int]bool{}
RoomsLoop:
	for i := range d.rand.Perm(len(d.rooms)) {
		for j := range d.rand.Perm(len(d.rooms[i].edges)) {
			room := d.rooms[i]
			edge := room.edges[j]
			x := edge.X
			y := edge.Y

			surroundingPoints := [4]geo.Point{
				{X: x - 1, Y: y},
				{X: x + 1, Y: y},
				{X: x, Y: y - 1},
				{X: x, Y: y + 1},
			}

			curRegion := -1
			for k := range surroundingPoints {
				tile := d.tiles[surroundingPoints[k].Y][surroundingPoints[k].X]
				if curRegion == -1 && tile.region != 0 {
					curRegion = tile.region
				} else if tile.region != curRegion &&
					tile.region != 0 &&
					!connectedRegions[tile.region] {

					d.tiles[y][x].material = Door
					connectedRegions[tile.region] = true
					connectedRegions[curRegion] = true

					continue RoomsLoop
				}
			}

		}
	}
}
