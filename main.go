package main

import (
	"errors"
	"fmt"
)

func main() {
	_, _, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title : ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	content, err := getUserInput("Note Content : ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	fmt.Println("Your note titled", title, "has the following content :")
	fmt.Println(content)
	fmt.Println("Saving the note succeeded!")
	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("field is required.")
	}

	return value, nil
}
