package main

import webserver "go-archtype-ddd/internal/infraestructure/web-server"

func main() {

	server := NewWebApplication()
	RunWebServer(server)
}

func RunWebServer(httpServer webserver.HttpServer) {
	httpServer.Start("3031")
}
