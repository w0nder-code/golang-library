package httpserv

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() {
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(s.httpHandlers.HandleAddBook)
	router.Path("/book/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetBook)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleReadBook)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteBook)
	router.Path("/books").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllBooks)

	http.ListenAndServe(":9091", router)

}
