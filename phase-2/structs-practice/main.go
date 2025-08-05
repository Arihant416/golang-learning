package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/note"
	"practice/todo"
	"strings"
)

func main() {
	title, content := getNoteDetails()
	text := getUserInput("Todo text: ")

	todo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.DisplayNote()
	err = todo.Save()

	if err != nil {
		fmt.Println("Failed to Save the Todo")
		return
	}
	fmt.Println("Todo Saved Successfully")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNote()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Failed to Save the Note")
		return
	}
	fmt.Println("Note Saved Successfully")
}

func getNoteDetails() (string, string) {
	title := getUserInput("Enter Title:")
	content := getUserInput("Enter Content")

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
