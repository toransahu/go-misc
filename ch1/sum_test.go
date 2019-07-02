//
// tests.go
// Copyright (C) 2019 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package ch1

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5,4)
	if total != 9 {
		t.Errorf("Sum(5,4) == %v, want %v", total, 9)
	}
}