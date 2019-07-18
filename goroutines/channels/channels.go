//
// channels.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

// Channels explains usage of channels with goroutines
func Channels() {
	message := make(chan string)

	go func(msg string) {
		fmt.Println("Msg yet to Send to channel")
		message <- msg
		fmt.Println("Msg Sent to channel")
	}("ping")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	msg := <-message
	fmt.Println(msg)
}

func main() {
	Channels()
}
