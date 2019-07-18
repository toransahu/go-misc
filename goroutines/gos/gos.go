//
// gos.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// RunGos explains goroutines using `go` statement
func RunGos() {
	fmt.Println("`Go` routine examples:")
	f("direct call")

	go f("goroutine call")

	go func(msg string) {
		fmt.Println(msg)
	}("goroutine call from anonymous func")

	fmt.Scanln()
	fmt.Println("Done!")

}

func main() {
	RunGos()
}
