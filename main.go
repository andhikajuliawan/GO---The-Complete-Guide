package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	Display()
	saver
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("1")

	title, content := getNoteData()
	todoText := getUserInput("Todo Text :")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	outputData(userTodo)

	add := add(4, 6)
	fmt.Println(add)
}

func add[T int | string | float64](a, b T) T {
	return a + b
}

func printSomething(value interface{}) {

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("string:", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title :")
	content := getUserInput("Note Content :")

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
