package main

import "fmt"

func main() {

	messages := make(chan string)

	fmt.Println("messages : ", messages)

	go func() { messages <- "ping" }()

	msg := <- messages
	fmt.Println(msg)
}