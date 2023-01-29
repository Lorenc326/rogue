package session

import (
	"rogue.game/maps"
	"rogue.game/user"
)

type Session struct {
	step    int
	gamemap *maps.GameMap
	user    *user.User
}

func (s *Session) Init() {
	s.gamemap = maps.Read("default")
	s.user = &user.User{}
	s.user.Extract(s.gamemap)
}

func (s *Session) Render() *maps.UserVision {
	s.step++
	s.user.Insert(s.gamemap)
	userMap := maps.UserVision{}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			userMap[i][j] = (*s.gamemap)[i][j]
		}
	}
	return &userMap
}
