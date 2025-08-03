package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input.")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (userNote Note) DisplayNote() {
	fmt.Println(userNote.title)
	fmt.Println(userNote.content)
	fmt.Println(userNote.createdAt.Date())
}
