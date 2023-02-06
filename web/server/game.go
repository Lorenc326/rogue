package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rogue.game/core/event"
	"rogue.game/core/session"
	"rogue.game/graphic"
)

const winStatus = 228

var cache = map[string]*session.Session{}

func handleActiveSession(c *gin.Context, sess *session.Session) bool {
	if sess.IsEnded {
		return true
	}

	var body event.Event
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return false
	}
	if err := sess.React(body); err != nil {
		log.Println(err)
		return false
	}
	return sess.IsEnded
}

func handleGame(c *gin.Context) {
	id := c.Param("id")

	if _, ok := cache[id]; !ok {
		cache[id] = session.New(
			graphic.NewASCII(5, true, false),
			session.SessionParametrs{Seed: 100, Width: 100, Height: 75},
		)
		c.String(http.StatusOK, cache[id].Draw())
		return
	}

	s := cache[id]
	// errors are ignored, the only change to consider is done/active status code
	status := http.StatusOK
	if win := handleActiveSession(c, s); win {
		status = winStatus
	}
	c.String(status, s.Draw())
}
