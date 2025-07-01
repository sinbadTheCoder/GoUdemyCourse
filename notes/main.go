package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func outputData(data outputable) error {
	err := saveData(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	displayData(data)
	return nil
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func displayData(data displayer) {
	data.Display()
}

func main() {
	theNote, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	theTodo, err := getTodo()
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(theTodo)
	outputData(theNote)

	fmt.Println("Save succeeded")
}

func getTodo() (t todo.Todo, err error) {
	fmt.Println("Enter new todo below...")
	content, _ := getUserInput("Todo content:")

	t, err = todo.New(content)
	if err != nil {
		fmt.Println(err)
		return todo.Todo{}, err
	}

	return t, nil
}

func getNoteData() (n note.Note, err error) {

	fmt.Println("Enter new note below...")
	title, _ := getUserInput("Note title:")
	content, _ := getUserInput("Note content:")

	n, err = note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return note.Note{}, err
	}

	return n, nil
}

func getUserInput(prompt string) (string, error) {
	var text string
	fmt.Printf("%s ", prompt)

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	if text == "" {
		return "", errors.New("invalid input")
	}

	return text, nil
}
