//
// maps.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package etc

import "fmt"

func main() {
	asciiTable := make(map[string]int)
	asciiTable["A"] = 65
	asciiTable["a"] = 97
	fmt.Println(asciiTable)

	map1 := make(map[string]string)
	map1["A"] = "65"
	fmt.Println(map1)

	map2 := map[string]string{"A": "65", "a": "97"}
	fmt.Println(map2)

	map3 := map[string][]string{
		"in": []string{"India", "Bharat"},
	}
	fmt.Println(map3)
}
