package server

import (
	"fmt"
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
		}

		output := cache[id].Render()
		c.String(http.StatusOK, output.String())
	})
	r.Run(fmt.Sprintf(":%d", port))
}
