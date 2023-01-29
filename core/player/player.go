package player

import (
	"errors"

	"rogue.game/core/maps"
	"rogue.game/core/symbol"
)

const offset = 5

type Player struct {
	I, J int // much clearer what is expected then with x, y
}

func (u *Player) Extract(m *maps.GameMap) {
	for i, row := range *m {
		for j, col := range row {
			if symbol.Player == col {
				u.I = i
				u.J = j
				(*m)[i][j] = symbol.Floor
				return
			}
		}
	}
}

// awful huck so map template HAS to have 5 spaces offsets to not break vision haha
func (u *Player) RenderVision(m *maps.GameMap) *maps.UserVision {
	userMap := maps.UserVision{}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			userMap[i][j] = (*m)[u.I-offset+i][u.J-offset+j]
		}
	}
	u.insert(&userMap)
	return &userMap
}

func (u *Player) insert(m *maps.UserVision) {
	(*m)[offset][offset] = symbol.Player
}

// unnecessary?
func (u *Player) Clone() *Player {
	return &Player{I: u.I, J: u.J}
}

func (u *Player) Move(m *maps.GameMap, direction string) error {
	switch direction {
	case "up":
		u.I--
	case "right":
		u.J++
	case "down":
		u.I++
	case "left":
		u.J--
	default:
		return errors.New("not defined")
	}
	return u.validateDestination(m)
}

func (u *Player) Victory(m *maps.GameMap) bool {
	location := (*m)[u.I][u.J]
	switch location {
	case symbol.End:
		return true
	}
	return false
}

func (u *Player) validateDestination(m *maps.GameMap) error {
	location := (*m)[u.I][u.J]
	switch location {
	case symbol.Wall, symbol.Lava:
		return errors.New("ahhhhh")
	}
	return nil
}
