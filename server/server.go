package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Listen(port int) {
	r := gin.Default()
	r.POST("/game/:id", func(c *gin.Context) {
		id := c.Param("id")
		log.Println(id)
		c.String(http.StatusOK, "map")
	})
	r.Run(fmt.Sprintf(":%d", port))
}
