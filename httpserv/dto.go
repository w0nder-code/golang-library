package httpserv

import (
	"encoding/json"
	"time"
)

type BookDTO struct {
	Title string `json:"title"`
	Athor string `json:"author"`
	Pages int    `json:"pages"`
}

type ErrorDTO struct {
	Message string `json:"message"`
	Time    time.Time
}

func (e ErrorDTO) ErrorToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}
