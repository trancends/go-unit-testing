package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetings(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result := Greetings("John")
		assert.Equal(t, "Hello, John", result, "Result should be 'Hello, John'")
	})
	t.Run("Ben", func(t *testing.T) {
		result := Greetings("Ben")
		assert.Equal(t, "Hello, Ben", result, "Result should be 'Hello, Ben'")
	})
}

func TestGreetingsBen(t *testing.T) {
	result := Greetings("Ben")

	// if result != "Hello, Ben" {
	// 	t.Error("Wrong Result: " + result)
	// }
	assert.Equal(t, "Hello, Ben", result, "Result should be 'Hello, Ben'")
}

func TestHitung(t *testing.T) {
	testTable := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"10+20", 10, 20, 30},
		{"5+5", 5, 5, 10},
		{"-10+10", -10, 10, 0},
	}

	for _, testcase := range testTable {
		t.Run(testcase.name, func(t *testing.T) {
			result := testcase.x + testcase.y
			assert.Equal(t, testcase.expected, result, fmt.Sprintf("Result should be %d", result))
		})
	}
}
