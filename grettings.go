// Package greetings provides functions to generate greeting messages.
package greetings

import "fmt"
// Hello returns a greeting message for the given name.
func Hello(name string) string {
	message := fmt.Sprintf("Hi ,%v.Welcome!",name)
	return message
}

