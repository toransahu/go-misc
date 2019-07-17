//
// interface.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

// UseInterface explains creating a arbitrary struct without using any well defined struct - similar to python dict
// or say creating map of string vs objects
func UseInterface() {
	i1 := map[string]interface{}{
		"id":   1111,
		"name": "Toran",
		"interests": []interface{}{
			"technology",
			"trekking",
			"biking",
		},
		"extra": map[string]interface{}{
			"edu": "graduation",
		},
	}
	fmt.Println(i1)

	i2 := []interface{}{
		map[string]interface{}{
			"id":   1111,
			"name": "Toran",
		},
		map[string]interface{}{
			"id":   1112,
			"name": "Abhishek",
		},
	}
	fmt.Println(i2)
}

// UseInterfaceTypeAssertion explains usage of empty inteface{} & type assertions.
func UseInterfaceTypeAssertion() {
	var i interface{}
	i = "Some String"
	fmt.Println(i, i.(string)) // type assertion

	i = 11
	fmt.Println(i, i.(int)) // type assertion

	i = 11.11
	fmt.Println(i, i.(float64)) // type assertion
	fmt.Println(i, i.(int))     // type assertion - panic: interface conversion: interface {} is float64, not int
}

// InterfaceTypeAssertionInSwitch explains usage if interface.(type) within switch case (only possible with switch)
func InterfaceTypeAssertionInSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is an integer\n", v)
	case float64:
		fmt.Printf("%v is a float64\n", v)
	case string:
		fmt.Printf("%v is a string\n", v)
	default:
		fmt.Println("Couldn't detect")
	}

}

func main() {
	UseInterface()

	var i interface{}
	i = 1
	InterfaceTypeAssertionInSwitch(i)
	i = 1.1
	InterfaceTypeAssertionInSwitch(i)
	i = "Something"
	InterfaceTypeAssertionInSwitch(i)

	// UseInterfaceTypeAssertion()
}
