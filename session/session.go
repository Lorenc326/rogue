package session

import (
	"rogue.game/maps"
	"rogue.game/user"
)

type Session struct {
	step    int
	gamemap *maps.GameMap

	// user is extracted from map source
	user *user.User
}

func (s *Session) Init() {
	s.gamemap = maps.Read("default")
	s.user = &user.User{}
	s.user.Extract(s.gamemap)
}

func (s *Session) Render() *maps.UserVision {
	s.step++
	return s.user.RenderVision(s.gamemap)
}
