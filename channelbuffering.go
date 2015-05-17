package main

import "fmt"

func main() {
	messages := make(chan string, 10)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)

	go func () {
		fmt.Println(<-messages)
		messages <- "done"
	}()

	var val string
	fmt.Scanln(&val)

	fmt.Println(<-messages)
}