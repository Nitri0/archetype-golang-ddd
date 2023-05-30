package main

import (
	"go-archtype-ddd/internal/application"
	stub_persistence "go-archtype-ddd/internal/infraestructure/persistence/stub"
	gin_webserver "go-archtype-ddd/internal/infraestructure/web-server/gin"
)

func NewWebApplication() *gin_webserver.GinHttpServer {

	orderRepository := stub_persistence.NewStubOrderRepository()
	productRepository := stub_persistence.NewStubProductRepository()

	createOrderUseCase := application.NewCreateOrderCommandHandler(orderRepository, productRepository)

	server := gin_webserver.NewGinHttpServer(createOrderUseCase)
	return server

}
