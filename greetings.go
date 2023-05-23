package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf("Hey you dumb, %v! How you doing?", name)
	return message, nil
}
