package symbol

import "rogue.game/core/dungeon"

const (
	Player = "@"
	Floor  = "."
	Tunnel = "_"
	Wall   = "#"
	Void   = " "
	Rat    = "r"
	Door   = "="
	End    = "x"
	Lava   = "o"
	Key    = "k"
)

var MaterialToSymbol = map[dungeon.Material]string{
	dungeon.Wall:   Wall,
	dungeon.Floor:  Floor,
	dungeon.Door:   Door,
	dungeon.Tunnel: Tunnel,
	dungeon.Finish: End,
	dungeon.Player: Player,
	dungeon.Rat:    Rat,
}
