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
	s.player = &player.Player{Coord: *s.Floor.Find(symbol.Player)}
	s.Floor.Replace(s.player.Coord, symbol.Floor)
	return &s
}

func (s *Session) React(e event.Event) error {
	projected := *s.player
	switch e.Action {
	case event.Move:
		err := projected.Move(s.Floor, e.Direction, 1)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported")
	}
	s.player = &projected
	s.IsEnded = s.player.Victory(s.Floor)
	s.step++
	return nil
}
