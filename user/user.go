package user

import (
	"errors"

	"rogue.game/maps"
	"rogue.game/symbol"
)

const offset = 5

type User struct {
	I, J int // much clearer what is expected then with x, y
}

func (u *User) Extract(m *maps.GameMap) {
	for i, row := range *m {
		for j, col := range row {
			if symbol.User == col {
				u.I = i
				u.J = j
				(*m)[i][j] = symbol.Floor
				return
			}
		}
	}
}

// awful huck so map template HAS to have 5 spaces offsets to not break vision haha
func (u *User) RenderVision(m *maps.GameMap) *maps.UserVision {
	userMap := maps.UserVision{}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			userMap[i][j] = (*m)[u.I-offset+i][u.J-offset+j]
		}
	}
	u.insert(&userMap)
	return &userMap
}

func (u *User) insert(m *maps.UserVision) {
	(*m)[offset][offset] = symbol.User
}

// unnecessary?
func (u *User) Clone() *User {
	return &User{I: u.I, J: u.J}
}

func (u *User) Move(m *maps.GameMap, direction string) error {
	switch direction {
	case "top":
		u.I--
	case "right":
		u.J++
	case "bottom":
		u.I++
	case "left":
		u.J--
	default:
		return errors.New("not defined")
	}
	return u.validateDestination(m)
}

func (u *User) Victory(m *maps.GameMap) bool {
	location := (*m)[u.I][u.J]
	switch location {
	case symbol.End:
		return true
	}
	return false
}

func (u *User) validateDestination(m *maps.GameMap) error {
	location := (*m)[u.I][u.J]
	switch location {
	case symbol.Wall, symbol.Lava:
		return errors.New("ahhhhh")
	}
	return nil
}
