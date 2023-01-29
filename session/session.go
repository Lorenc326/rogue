package session

import (
	"rogue.game/maps"
)

type Session struct {
	gamemap *maps.GameMap
}

func (s *Session) Init() {
	s.gamemap = maps.Read("default")
}

func (s *Session) Render() *maps.UserVision {
	userMap := maps.UserVision{}
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			userMap[i][j] = (*s.gamemap)[i][j]
		}
	}
	return &userMap
}
