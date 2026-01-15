package httpserv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rest/library"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	bookList *library.List
}

func NewHttpHandlers(bookList *library.List) *HTTPHandlers {
	return &HTTPHandlers{
		bookList: bookList,
	}
}

/*
pattern: /books
method: POST
info: JSON in HTTP body

succeed:
  - status code: 201 Created
  - response body: JSON represent created book

failed:
  - status code: 400, 409, 500
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleAddBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO BookDTO

	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ErrorToString(), http.StatusBadRequest)
		return
	}

	if err := bookDTO.ValidateToAdd(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ErrorToString(), http.StatusBadRequest)
		return
	}

	libraryNewBook := library.AddBook(bookDTO.Title, bookDTO.Athor, bookDTO.Pages)
	if err := h.bookList.AddNewBook(libraryNewBook); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, library.ErrBookAlreadyExist) {
			http.Error(w, errDTO.ErrorToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(libraryNewBook, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
		return
	}
}

/*
pattern: /books/{title}
method: GET
info: pattern

succeed:
	- status code: 200 OK
	- response body: JSON represent found book
failed:
	- status code: 400, 404, 500
	- response body: JSON with error + time

*/

func (h *HTTPHandlers) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	foundBook, err := h.bookList.ListBook(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, errDTO.ErrorToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
	}

	b, err := json.MarshalIndent(foundBook, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

/*
	pattern: /books/{title}
method: GET
info: path

succeed:
	- status code: 200 OK
	- response body: JSON represent found books
failed:
	- status code: 400, 404, 500
	- response body: JSON with error + time

*/

func (h *HTTPHandlers) HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := h.bookList.ListBooks()
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ErrorToString(), http.StatusBadRequest)
	}

	b, err := json.MarshalIndent(books, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

/*
pattern: /books/{title}
method: DELETE
info: pattern

succeed:
  - status code: 200 OK
  - response body: -ь

failed:
  - status code: 400, 404, 500
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	err := h.bookList.DeleteBook(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, errDTO.ErrorToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusOK)
}

/*
pattern: /books/{title}
method: PATCH
info: pattern

succeed:
  - status code: 200 OK
  - response body: JSON represent patched book

failed:
  - status code: 400, 404, 500
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	err := h.bookList.ReadBook(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, errDTO.ErrorToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusOK)
}
