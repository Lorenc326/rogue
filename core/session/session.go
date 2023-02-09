package session

import (
	"errors"
	"time"

	"github.com/Lorenc326/rogue/core/dungeon"
	"github.com/Lorenc326/rogue/core/event"
	"github.com/Lorenc326/rogue/core/maps"
	"github.com/Lorenc326/rogue/core/player"
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
	Seed   int64 `json:"seed"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

func (p *SessionParametrs) SetDefaults() {
	if p.Seed == 0 {
		p.Seed = time.Now().UnixMilli()
	}
	if p.Width == 0 {
		p.Width = 100
	}
	if p.Height == 0 {
		p.Height = 75
	}
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
