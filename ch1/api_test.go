package ch1

import (
    "github.com/toransahu/go-misc/utils/test"
    "testing"
)

// MockApollo mocks Apollo 
type MockApollo struct {
    
}

// About mimics Apollo->About for mocked env
func (mockApollo *MockApollo) About() string {
    return "v0.1"
}

// TestPAPI_GetApolloVersion tests PAPI->GetApolloVersion
func TestPAPI_GetApolloVersion(t *testing.T) {
    mockedApollo := new(MockApollo)
    papi := PAPI{
        apollo:mockedApollo,
    }
    assert := new(test.Assert)  // or Assert{}
    assert.Equal(t, papi.GetApolloVersion(), "v0.1")
}
