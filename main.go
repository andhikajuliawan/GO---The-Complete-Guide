package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	var newNote note.Note
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title : ")
	content := getUserInput("Note Content : ")

	fmt.Println("Your note titled", title, "has the following content :")
	fmt.Println(content)
	fmt.Println("Saving the note succeeded!")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}
