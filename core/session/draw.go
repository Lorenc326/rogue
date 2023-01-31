package session

import (
	"rogue.game/core/maps"
	"rogue.game/core/player"
)

type Renderer interface {
	Render(c DrawContext) string
}

type DrawContext struct {
	Floormap *maps.FloorMap
	IsEnded  bool
	Player   *player.Player
}

func (s *Session) Draw() string {
	return s.renderer.Render(DrawContext{
		Floormap: s.floormap,
		IsEnded:  s.IsEnded,
		Player:   s.player,
	})
}
