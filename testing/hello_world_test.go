package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HelloWordlSuite struct {
	suite.Suite
}

func TestGreetings(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result, err := Greetings("John")
		assert.Equal(t, "Hello, John", result, "Result should be 'Hello, John'")
		assert.Nil(t, err)
	})
	t.Run("Ben", func(t *testing.T) {
		result, _ := Greetings("Ben")
		assert.Equal(t, "Hello, Ben", result, "Result should be 'Hello, Ben'")
	})
}

func TestGreetingsFailed(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result, err := Greetings("John")
		assert.Equal(t, "Hell, John", result, "Result should be 'Hello, John'")
		assert.Nil(t, err)
	})
	t.Run("Ben", func(t *testing.T) {
		result, _ := Greetings("Ben")
		assert.Equal(t, "Hell, Ben", result, "Result should be 'Hello, Ben'")
	})
}

func TestHitung(t *testing.T) {
	t.Run("10+20", func(t *testing.T) {
		result := Hitung(10, 20)
		assert.Equal(t, 30, result, fmt.Sprintf("Result should be:%v", 30))
	})
}

func TestHitung2(t *testing.T) {
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

func (s *HelloWordlSuite) TestHelloWorld() {
	result, err := Greetings("John")
	assert.Equal(s.T(), "Hello, John", result, "Result should be 'Hello, John'")
	assert.Nil(s.T(), err)
}

func (s *HelloWordlSuite) TestHelloWorldFailed() {
	result, err := Greetings("John")
	assert.Equal(s.T(), "", result, "Result should be ''")
	assert.NotNil(s.T(), err)
}

func TestHelloWorldSuite(t *testing.T) {
	suite.Run(t, new(HelloWordlSuite))
}
