package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	phrases := []string{
		"E ai %v! Como vai o estudo de golang?",
		"Eh muita confusao na StackSpot viu, %v.. Nao tem quem aguente!",
		"Fala %v! Quando tu vai jantar?",
	}

	return phrases[rand.Intn(len(phrases))]
}

func HelloNew(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = msg
	}

	return messages, nil
}
