//
// verbs.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package ch1

import "fmt"

// Values in diff formats within string.
func verbs() {
	fmt.Printf("%d\n", 10)   // decimal
	fmt.Printf("%x\n", 10)   // hexadecimal
	fmt.Printf("%o\n", 10)   // octal
	fmt.Printf("%b\n", 10)   // binary
	fmt.Printf("%f\n", 10.0) // float 6 decimal places
	fmt.Printf("%g\n", 10.0) // TODO: float zero/15/more decimal places
	fmt.Printf("%e\n", 10.0) // scientific notation - with e+ e-
	fmt.Printf("%t\n", true) // boolean
	fmt.Printf("%c\n", 10)   // TODO: rune (unicode code point)
	fmt.Printf("%s\n", "10") // string
	fmt.Printf("%q\n", "10") // quoted string

}
