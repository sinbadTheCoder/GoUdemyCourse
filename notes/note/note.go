package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Println("\n------------------------")
	fmt.Println("| Title:", n.title)
	fmt.Println("| Content:", n.content)
	fmt.Println("| Time Created:", n.createdAt)
	fmt.Println("------------------------")
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("title cannot be empty")
	}
	if content == "" {
		return Note{}, errors.New("content cannot be empty")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
