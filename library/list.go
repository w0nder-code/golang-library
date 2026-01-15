package rest

import (
	"sync"

	"github.com/k0kubun/pp"
)

var mtx = sync.RWMutex{}

type List struct {
	bookList map[string]Book
}

func NewList() *List {
	return &List{
		bookList: make(map[string]Book),
	}
}

func (l *List) AddNewBook(book Book) error {
	mtx.Lock()
	defer mtx.Unlock()
	if _, ok := l.bookList[book.Title]; ok {
		return ErrBookAlreadyExist
	}

	l.bookList[book.Title] = book
	return nil
}

func (l List) ListBooks() (map[string]Book, error) {
	mtx.Lock()
	defer mtx.Unlock()
	tempMap := make(map[string]Book)

	if l.bookList == nil {
		return nil, ErrBookNotFound
	}

	for k, v := range l.bookList {
		tempMap[k] = v
	}

	pp.Println(tempMap)
	return tempMap, nil
}

func (l List) ListBook(title string) (Book, error) {
	mtx.Lock()
	defer mtx.Unlock()

	if _, ok := l.bookList[title]; !ok {
		return Book{}, ErrBookNotFound
	}

	pp.Println(l.bookList[title])
	tempMap := make(map[string]Book)

	tempMap[title] = l.bookList[title]

	return tempMap[title], nil
}

func (l *List) DeleteBook(title string) error {
	mtx.Lock()
	defer mtx.Unlock()

	if _, ok := l.bookList[title]; !ok {
		return ErrBookNotFound
	}

	delete(l.bookList, title)
	return nil

}

func (l *List) ReadBook(title string) error {
	mtx.Lock()
	defer mtx.Unlock()

	book, ok := l.bookList[title]
	if !ok {
		return ErrBookNotFound
	}

	book.ReadenFunc()
	l.bookList[title] = book

	return nil
}
