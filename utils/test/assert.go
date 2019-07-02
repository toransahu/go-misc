package test

import "testing"

// Assert provides Testing utils structure
type Assert struct {
    actual string
    expected string
}

// Equal compares `want` & `got`
func (a *Assert) Equal(t *testing.T, actual string, expected string) {
    if actual != expected {
        t.Errorf("Want %v but got %v", expected, actual)
    }
}
