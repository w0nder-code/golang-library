package main

import (
	"rest/httpserv"
	"rest/library"
)

func main() {
	bookList := library.NewList()
	handlers := httpserv.NewHttpHandlers(bookList)
	server := httpserv.NewHTTPServer(handlers)

	server.StartServer()

}
