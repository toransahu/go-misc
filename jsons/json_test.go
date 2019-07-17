package jsons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodingDecoding(t *testing.T) {
	// GIVEN
	m := Message{1111, "Toran", "tOpSeCrEt"}
	// WHEN
	bytes, err := m.Encode()
	if err != nil {
	}
	// THEN
	// convert json string into byte array and compare
	// NOTE: here in json string - no spaces anywhere
	assert.Equal(t, []byte(`{"id":1111,"name":"Toran"}`), bytes, "They should be equal")

	// GIVEN
	anotherBytes := bytes
	// WHEN
	var anotherM Message
	anotherErr := Decode(anotherBytes, &anotherM) // pass by address
	if anotherErr != nil {
	}
	// THEN
	assert.Equal(t, Message{1111, "Toran", ""}, anotherM) // NOTE: secret is empty STRING

	// OR Dynamic Typing
	var varM interface{}
	Decode(anotherBytes, &varM) // WARN: handle error
	// NOTE: 1. dynamic map + struct creation
	// 		 2. unpacking of int as float64 hence %s/1111/1111.0
	assert.Equal(t, map[string]interface{}{"id": 1111.0, "name": "Toran"}, varM)
}
