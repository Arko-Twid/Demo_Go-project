package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//Check if the parameter name is empty
	if len(name) == 0 {
		return name, errors.New("Cannot accept an empty name")
	}
	//Creating a message when the parameter is not empty
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}
func randomFormat() string {
	// Creating a slice of welcome messages
	formats := []string{
		"Hey man, %v! Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}
