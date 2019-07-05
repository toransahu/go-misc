//
// decoding.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package json

import "encoding/json"

func Decode(bytes []byte, v interface{}) error {
	err := json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}
