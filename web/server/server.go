package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Listen(port int) {
	r := gin.Default()

	// I know it's not very RESTish but I don't care
	r.POST("/game/:id", handleGame)
	r.POST("/game/:id/start", handleStart)

	r.Run(fmt.Sprintf(":%d", port))
}
