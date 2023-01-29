package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Listen(port int) {
	r := gin.Default()
	r.POST("/game/:id", handleGame)
	r.Run(fmt.Sprintf(":%d", port))
}
