package user

import (
	"rogue.game/maps"
	"rogue.game/symbol"
)

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

func (u *User) Insert(m *maps.GameMap) {
	(*m)[u.I][u.J] = symbol.User
}
