package user

import (
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
	u.insert(m)
	userMap := maps.UserVision{}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			userMap[i][j] = (*m)[u.I-offset+i][u.J-offset+j]
		}
	}
	return &userMap
}

func (u *User) insert(m *maps.GameMap) {
	(*m)[u.I][u.J] = symbol.User
}
