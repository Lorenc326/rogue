package session

import (
	"errors"
	"time"

	"rogue.game/core/event"
	"rogue.game/core/maps"
	"rogue.game/core/player"
	"rogue.game/core/symbol"
)

type Session struct {
	step      int
	createdAt time.Time
	Floor     *maps.Floor
	IsEnded   bool
	player    *player.Player
	renderer  Renderer
}

func New(renderer Renderer) *Session {
	s := Session{}
	s.Floor = maps.Read("default")
	s.renderer = renderer
	s.createdAt = time.Now().UTC()
	s.player = &player.Player{Point: *s.Floor.Find(symbol.Player)}
	s.Floor.Replace(s.player.Point, symbol.Floor)
	return &s
}

func (s *Session) React(e event.Event) error {
	if s.IsEnded {
		return nil
	}
	switch e.Action {
	case event.Move:
		projected := s.player.Point
		projected.Move(e.Direction, 1)
		err := s.player.ValidateDestination(*s.Floor, projected)
		if err != nil {
			return err
		}
		s.player.Point = projected
	default:
		return errors.New("unsupported")
	}
	s.IsEnded = s.player.Victory(s.Floor)
	s.step++
	return nil
}
