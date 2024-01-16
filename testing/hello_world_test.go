package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetings(t *testing.T) {
	result := Greetings("John")

	// if result != "Hello, John" {
	// 	t.Error("Wrong Result: " + result)
	// }
	assert.Equal(t, "Hello, John", result, "Result should be 'Hello, John'")
}
