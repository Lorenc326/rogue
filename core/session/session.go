package session

import (
	"errors"
	"time"

	"rogue.game/core/dungeon"
	"rogue.game/core/event"
	"rogue.game/core/maps"
	"rogue.game/core/player"
)

type Session struct {
	step      int
	createdAt time.Time
	Floor     maps.Floor
	IsEnded   bool
	player    *player.Player
	renderer  Renderer
}

type SessionParametrs struct {
	Seed   int64
	Width  int
	Height int
}

func New(renderer Renderer, params SessionParametrs) *Session {
	s := Session{}
	d := dungeon.Generate(params.Seed, params.Width, params.Height)
	s.Floor = d.Tiles
	s.renderer = renderer
	s.createdAt = time.Now().UTC()
	s.player = &player.Player{Point: *s.Floor.Find(dungeon.Player)}
	s.Floor.Replace(s.player.Point, dungeon.Floor)
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
		err := s.player.ValidateDestination(s.Floor, projected)
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
