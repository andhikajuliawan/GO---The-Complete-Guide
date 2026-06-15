package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	title := getUserInput("Note Title :")
	content := getUserInput("Note Content :")

	fmt.Println("Your note titled", title, "has the following content :")
	fmt.Println(content)
	fmt.Println("Saving the note succeeded!")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
