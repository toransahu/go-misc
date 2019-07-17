//
// reference_type.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package jsons

import (
	"encoding/json"
	"fmt"
)

// JSONReferenceType explains unmarshalling/deserialization of only limited objects
func JSONReferenceType() {
	i := map[string]interface{}{
		"id":   1111,
		"name": "Toran",
		"interests": []interface{}{
			"technology",
			"trekking",
			"biking",
		},
		"edu": map[string]interface{}{
			"degree":  "graduation",
			"college": "NITrr",
			"passout": 2015,
			"courses": []string{"CS", "NW", "OS"},
		},
	}

	type Education struct {
		degree  string
		college string
		passout int
		courses []string
	}

	type Foo struct {
		Edu *Education `json:"edu"` // export (Capital leter varname) & add json tag in small to match above map
	}

	bytes, err := json.Marshal(i)

	if err != nil {
		fmt.Printf("Some error while marshalling: %v", err)
	}
	fmt.Println(bytes)

	var f Foo
	json.Unmarshal(bytes, &f)
	fmt.Printf("%+v\n", f) // see pointer is not nil, there is soemthing TODO: find the value
}
