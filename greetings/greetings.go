package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {

	message := fmt.Sprintf(randomFormat(), name)
	if name == "" {
		return "", errors.New("Empty name provided as parameter")
	}
	return message, nil
}

func randomFormat() string {
	formats := []string{"Hi there, %v. Welcome!", "Yo, %v, how're you doing?", "Welcome home, %v"}
	return formats[rand.Intn(len(formats))]

}
