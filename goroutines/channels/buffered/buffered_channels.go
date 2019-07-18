//
// buffered_channels.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func eg1() {
	fmt.Println("eg1 starts.....")
	message := make(chan string, 2) // create a buffered channel of buffer capacity for 2 values
	message <- "value 1"
	fmt.Println(len(message))
	fmt.Println(cap(message))
	fmt.Println(<-message)
	// error on next print
	fmt.Println(len(message)) // see now length has been reduced to 0
	fmt.Println(cap(message)) // anyhow capacity will be 2 always
	fmt.Println("eg1 ends.....")
}

func eg2() {
	fmt.Println("eg2 starts.....")
	message := make(chan string, 2) // create a buffered channel of buffer capacity for 2 values
	message <- "value 1"
	message <- "value 2"
	// message <- "value 3"
	fmt.Println(len(message))
	fmt.Println(cap(message))
	fmt.Println(<-message)
	fmt.Println(<-message)
	// error on next print, hence check first
	if len(message) > 0 {
		fmt.Println(<-message)
	}
	fmt.Println("eg2 ends.....")
}

func main() {
	eg1()
	eg2()
}
