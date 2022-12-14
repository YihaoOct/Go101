package demo

import "fmt"

func ChannelBasics() {
	// communicating via a channel:
	// goroutine 1 send "hello" --------> messages (channel of string)
	// goroutine 2 send "world" -------->
	// goroutine 0 read 1       <--------
	// goroutine 0 read 2       <--------

	// create the messages channel
	messages := make(chan string)

	// start a thread trying to write to channel
	// think Task.Run(() => sendMessage(...))
	go sendMessage(messages, "hello")

	// start another thread trying to write to channel
	go sendMessage(messages, "world")

	// in the main goroutine, read twice from the channel and print the results
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func sendMessage(ch chan<- string, m string) {
	ch <- m
}
