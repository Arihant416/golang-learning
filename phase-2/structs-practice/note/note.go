package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createAt"`
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonContent, 0644)

}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("InvalidInput.")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (userNote Note) DisplayNote() {
	fmt.Println(userNote.Title)
	fmt.Println(userNote.Content)
	fmt.Println(userNote.CreatedAt.Date())
}
