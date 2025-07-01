package main

import (
	"fmt"
	"reflect"
)

type mystr string

func (message *mystr) log() {
	fmt.Println(*message)
}

func main() {
	var message mystr = "This is a log message"
	message.log()

	printAnything("Hello alien!")
	printAnything(162)
	printAnything(100.2)
	printAnything(message)

	convertAnyToInt(42)
	convertAnyToInt(3.14)
	convertAnyToInt("fourty-two")
	convertAnyToInt("42")
}

func printAnything(value any) {

	switch value.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	case float64:
		fmt.Println(value)
	}
}

func convertAnyToInt(value any) {
	intValue, ok := value.(int)
	if ok {
		fmt.Println("successfully converted value to int: ", intValue)
		fmt.Printf("value is of type [%s]\n", reflect.ValueOf(value).Kind())
		fmt.Printf("adding one to int value gives %d\n", intValue+1)
		fmt.Printf("intValue is of type [%s]\n", reflect.ValueOf(value).Kind())
	} else {
		fmt.Println("failed to convert value to int: ", value)
		fmt.Printf("value is of type [%s]\n", reflect.ValueOf(value).Kind())
	}
}
