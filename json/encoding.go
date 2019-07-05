//
// encoding.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package json

import (
	"encoding/json"
)


func (m *Message) Encode() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
