package session

import (
	"errors"
	"time"

	"rogue.game/core/maps"
	"rogue.game/core/player"
)

type Event struct {
	Action    string `json:"action"`
	Direction string `json:"direction"`
}

type Session struct {
	step      int
	createdAt time.Time
	gamemap   *maps.GameMap
	IsEnded   bool
	// player is extracted from map source
	player *player.Player
}

func (s *Session) Init() {
	s.gamemap = maps.Read("default")
	s.createdAt = time.Now().UTC()
	s.player = &player.Player{}
	s.player.Extract(s.gamemap)
}

func (s *Session) React(event Event) error {
	projected := s.player.Clone()
	switch event.Action {
	case "move":
		err := projected.Move(s.gamemap, event.Direction)
		if err != nil {
			return err
		}
		s.player = projected
		s.IsEnded = s.player.Victory(s.gamemap)
	default:
		return errors.New("unsupported")
	}
	return nil
}

func (s *Session) RenderASCII() string {
	s.step++
	return s.player.RenderVision(s.gamemap).String()
}
