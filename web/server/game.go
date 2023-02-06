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

type ErrResponse struct {
	Error string `json:"error"`
}

type StartBody struct {
	Seed   int `json:"seed"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func handleStart(c *gin.Context) {
	id := c.Param("id")

	var body session.SessionParametrs
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ErrResponse{err.Error()})
		return
	}
	body.SetDefaults()

	cache[id] = session.New(graphic.NewASCII(5, true, false), body)
	c.String(http.StatusOK, cache[id].Draw())
}

func handleGame(c *gin.Context) {
	id := c.Param("id")

	s, ok := cache[id]
	if !ok {
		c.JSON(http.StatusBadRequest, ErrResponse{"don't play with me, there is no such game"})
		return
	}

	// errors are ignored, the only change to consider is done/active status code
	status := http.StatusOK
	if win := handleActiveSession(c, s); win {
		status = winStatus
	}
	c.String(status, s.Draw())
}

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
