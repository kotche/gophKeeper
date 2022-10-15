package errs

import "fmt"

// ConflictLoginError error of an existing login in the database
type ConflictLoginError struct {
	Username string
}

func (conflict ConflictLoginError) Error() string {
	return fmt.Sprintf("login %v already exists", conflict.Username)
}

// AuthenticationError user authentication error
type AuthenticationError struct{}

func (auth AuthenticationError) Error() string {
	return "invalid login/password"
}
