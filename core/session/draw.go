package session

import (
	"github.com/Lorenc326/rogue/core/maps"
	"github.com/Lorenc326/rogue/core/player"
)

type Renderer interface {
	Render(c DrawContext) string
}

type DrawContext struct {
	Floor   maps.Floor
	IsEnded bool
	Player  *player.Player
}

func (s *Session) Draw() string {
	return s.renderer.Render(DrawContext{
		Floor:   s.Floor,
		IsEnded: s.IsEnded,
		Player:  s.player,
	})
}
