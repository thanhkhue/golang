package main 

import "fmt"

func main() {

	messages := make(chan string) // create a new channel

	go func(){ // this is the way we create a new goroutine
		messages <- "ping"
	}() // send values into the channel

	msg := <- messages
	fmt.Println(msg)
}