// Package greetings provides functions to generate greeting messages.
package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting message for the given name.
/*func Hello(name string) string {
	message := fmt.Sprintf("Hi ,%v.Welcome!",name)
	return message
}*/
//Now here we are handling an error if user did not passe the user name in code then how our package will show error to user.

//we are naming the function using func keyword then give a name to function then passed name parameter.
//we passed string and error
func Hello(name string) (string, error) {
	if name == "" {
		//returning to user null or error if user did not pass any context or input
		return "", errors.New("name is not passed at end user")
	}
	message := fmt.Sprintf("Hi,%v.Welcome!", name)
	return message, nil
}
