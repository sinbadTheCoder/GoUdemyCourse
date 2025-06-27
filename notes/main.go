package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	theNote, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	theNote.Display()
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
