package symbol

import "github.com/Lorenc326/rogue/core/dungeon"

const (
	Player = "@"
	Floor  = "."
	Tunnel = "_"
	Wall   = "#"
	Void   = " "
	Rat    = "r"
	Door   = "H"
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
