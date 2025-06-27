package main

import (
	"fmt"
)

type mystr string

func (message *mystr) log() {
	fmt.Println(*message)
}

func main() {
	var message mystr = "This is a log message"
	message.log()
}

