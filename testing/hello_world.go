package testing

import (
	"fmt"
)

func Greetings(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("cant be empty")
	}
	return "Hello, " + name, nil
}

func Hitung(x, y int) int {
	return x + y
}
