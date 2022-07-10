package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	if name == "" {
		return "", errors.New("Empty name provided as parameter")
	}
	return message, nil
}
