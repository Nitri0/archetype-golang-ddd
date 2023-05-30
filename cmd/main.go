package main

import (
	"archetype-golang-ddd/internal/application"
	stub_persistence "archetype-golang-ddd/internal/infraestructure/persistence/stub"
	webserver "archetype-golang-ddd/internal/infraestructure/web-server"
	gin_webserver "archetype-golang-ddd/internal/infraestructure/web-server/gin"
)

func main() {
	server := FactoryWebApplication()
	RunWebServer(server)
}

func RunWebServer(httpServer webserver.HttpServer) {
	httpServer.Start("3031")
}

func FactoryWebApplication() *gin_webserver.GinHttpServer {

	orderRepository := stub_persistence.NewStubOrderRepository()
	productRepository := stub_persistence.NewStubProductRepository()

	createOrderUseCase := application.NewCreateOrderCommandHandler(orderRepository, productRepository)

	server := gin_webserver.NewGinHttpServer(createOrderUseCase)
	return server

}
