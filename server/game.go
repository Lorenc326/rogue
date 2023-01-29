package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rogue.game/core/session"
)

const winStatus = 228

var cache = map[string]*session.Session{}

func renderInitialSession(sess *session.Session) string {
	sess.Init()
	return sess.RenderASCII()
}

func handleActiveSession(c *gin.Context, sess *session.Session) bool {
	if sess.IsEnded {
		return true
	}

	var body session.Event
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
		cache[id] = &session.Session{}
		c.String(http.StatusOK, renderInitialSession(cache[id]))
		return
	}

	sess := cache[id]
	// errors are ignored, the only change to consider is done/active status code
	status := http.StatusOK
	if win := handleActiveSession(c, sess); win {
		status = winStatus
	}
	c.String(status, sess.RenderASCII())
}
