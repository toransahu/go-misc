// main.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"github.com/toransahu/go-misc/ch1"
	"github.com/toransahu/go-misc/jsons"
)

func main() {
	fmt.Println(ch1.Sum(1, 2))
	// ch1.Dup()
	jsons.JSONReferenceType()
}
