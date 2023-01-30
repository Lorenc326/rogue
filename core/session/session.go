package session

import (
	"errors"
	"time"

	"rogue.game/core/maps"
	"rogue.game/core/player"
	"rogue.game/core/renderer"
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
	// player is extracted from map source
	player *player.Player
}

func (s *Session) Init() {
	s.floormap = maps.Read("default")
	s.createdAt = time.Now().UTC()
	s.player = &player.Player{Coord: *s.floormap.Find(symbol.Player)}
	s.floormap.Replace(s.player.Coord, symbol.Floor)
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

func (s *Session) RenderASCII() string {
	return renderer.ASCII(renderer.PlayerCenteredMap(s.floormap, s.player, 5))
}

func (s *Session) RenderASCIIWide() string {
	return renderer.ASCIIWide(renderer.PlayerCenteredMap(s.floormap, s.player, 10))
}
