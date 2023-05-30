package webserver

type HttpServer interface {
	Start(port string)
}
