package assert

import "testing"

// Equal compares `want` & `got`
// Limitation here string Vs string
func Equal(t *testing.T, actual string, expected string) {
    if actual != expected {
        t.Errorf("Want %v but got %v", expected, actual)
    }
}
