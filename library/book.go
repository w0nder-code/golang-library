package rest

import "time"

type Book struct {
	Title    string
	Author   string
	Pages    int
	Readen   bool
	AddTime  time.Time
	ReadenAt *time.Time
}

func AddBook(title string, author string, pages int) Book {
	return Book{
		Title:    title,
		Author:   author,
		Pages:    pages,
		Readen:   false,
		AddTime:  time.Now(),
		ReadenAt: nil,
	}
}

func (b Book) ReadenFunc() {
	readenAt := time.Now()

	b.Readen = true
	b.ReadenAt = &readenAt
}
