package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("working ....")
	time.Sleep(time.Second)
	fmt.Println("done ...")

	done <- true
	done <- false
}

func main() {
	done := make(chan bool, 2)
	go worker(done)
	fmt.Println("done : ", <-done)
	fmt.Println("done : ", <-done)
}