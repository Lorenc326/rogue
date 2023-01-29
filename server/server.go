package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rogue.game/session"
)

var cache = map[string]*session.Session{}

func Listen(port int) {
	r := gin.Default()
	r.POST("/game/:id", func(c *gin.Context) {
		id := c.Param("id")

		if _, ok := cache[id]; !ok {
			sess := &session.Session{}
			sess.Init()
			cache[id] = sess
			c.String(http.StatusOK, sess.Render().String())
			return
		}

		sess := cache[id]
		if sess.IsEnded {
			c.String(http.StatusOK, "you won!!!")
			return
		}

		var body session.Event
		if err := c.BindJSON(&body); err != nil {
			log.Println(err)
			c.String(http.StatusOK, sess.Render().String())
			return
		}
		if err := sess.React(body); err != nil {
			log.Println(err)
			c.String(http.StatusOK, sess.Render().String())
		}
		if sess.IsEnded {
			c.String(http.StatusOK, "you won!!!")
			return
		}
		c.String(http.StatusOK, sess.Render().String())
	})
	r.Run(fmt.Sprintf(":%d", port))
}
