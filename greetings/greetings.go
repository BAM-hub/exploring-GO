package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	//create a random message using random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	//creates a slice of messages
	messages := []string{
		"Hi, %v. Welcome!",
		"Greate to see you, %v!",
		"Hail, %v!",
	}
    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
	return messages[rand.Intn(len(messages))]
}

// unused code
// message := fmt.Sprintf("Hello, %v. Welcome!", name) 
