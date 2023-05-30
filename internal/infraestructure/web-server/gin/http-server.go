package gin_webserver

import (
	"github.com/gin-gonic/gin"
)

type GinHttpServer struct {
	router *gin.Engine
}

func (g *GinHttpServer) Start(port string) {
	g.router.Run(":" + port)
}
