package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Save() error {

	fileName := "todo.json"
	jsonContent, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonContent, 0644)

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("InvalidInput.")
	}
	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) DisplayNote() {
	fmt.Println(todo)
}
