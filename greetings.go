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
