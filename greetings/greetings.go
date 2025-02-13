package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("Name cannot be empty")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(ramdomMessage(), name)
	return message, nil
}
func ramdomMessage() string {
	//A slice of messages
	messages := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return messages[rand.Intn(len(messages))]

}
