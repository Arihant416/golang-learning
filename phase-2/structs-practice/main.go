package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/note"
	"practice/todo"
	"strings"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	DisplayNote()
}

// type outputtable interface{
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteDetails()
	text := getUserInput("Todo text: ")

	todo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.DisplayNote()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Failed to Save the Data")
		return err
	}
	fmt.Println("Note Saved Successfully")
	return nil
}

func getNoteDetails() (string, string) {
	title := getUserInput("Enter Title: ")
	content := getUserInput("Enter Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
