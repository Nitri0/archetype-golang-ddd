package gin_webserver

import (
	"archetype-golang-ddd/internal/application"
	gin_handlers "archetype-golang-ddd/internal/infraestructure/web-server/gin/handlers"

	"github.com/gin-gonic/gin"
)

func NewGinHttpServer(
	createOrderUseCase *application.CreateOrderCommandHandler,
) *GinHttpServer {

	router := gin.Default()

	createOrderHttpHandler := gin_handlers.NewCreateOrderHttpHandler(createOrderUseCase)

	router.POST("/new-order", createOrderHttpHandler.Handle)

	return &GinHttpServer{
		router: router,
	}
}
