package rest

import (
	"errors"
	"time"
)

type Error struct {
	Message string
	Time    time.Time
}

func NewError(message string) Error {
	return Error{
		Message: message,
		Time:    time.Now(),
	}
}

var ErrBookNotFound = errors.New("Book not found")
var ErrBookAlreadyExist = errors.New("Book already exists in library")
