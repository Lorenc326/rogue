package session

import (
	"errors"

	"rogue.game/maps"
	"rogue.game/user"
)

type Event struct {
	Action    string `json:"action"`
	Direction string `json:"direction"`
}

type Session struct {
	step    int
	gamemap *maps.GameMap
	IsEnded bool
	// user is extracted from map source
	user *user.User
}

func (s *Session) Init() {
	s.gamemap = maps.Read("default")
	s.user = &user.User{}
	s.user.Extract(s.gamemap)
}

func (s *Session) React(event Event) error {
	projected := s.user.Clone()
	switch event.Action {
	case "move":
		err := projected.Move(s.gamemap, event.Direction)
		if err != nil {
			return err
		}
		s.IsEnded = s.user.Victory(s.gamemap)
	default:
		return errors.New("unsupported")
	}
	return nil
}

func (s *Session) Render() *maps.UserVision {
	s.step++
	return s.user.RenderVision(s.gamemap)
}
