package session

import (
	"errors"
	"time"

	"rogue.game/core/maps"
	"rogue.game/core/player"
	"rogue.game/core/symbol"
)

type Event struct {
	Action    string `json:"action"`
	Direction string `json:"direction"`
}

type Session struct {
	step      int
	createdAt time.Time
	floormap  *maps.FloorMap
	IsEnded   bool
	player    *player.Player
	renderer  Renderer
}

func New(renderer Renderer) *Session {
	s := Session{}
	s.floormap = maps.Read("default")
	s.renderer = renderer
	s.createdAt = time.Now().UTC()
	s.player = &player.Player{Coord: *s.floormap.Find(symbol.Player)}
	s.floormap.Replace(s.player.Coord, symbol.Floor)
	return &s
}

func (s *Session) React(event Event) error {
	projected := *s.player
	switch event.Action {
	case "move":
		err := projected.Move(s.floormap, event.Direction)
		if err != nil {
			return err
		}
		s.player = &projected
		s.IsEnded = s.player.Victory(s.floormap)
	default:
		return errors.New("unsupported")
	}
	s.step++
	return nil
}
