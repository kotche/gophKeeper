package errs

import "fmt"

type ConflictLoginError struct {
	Login string
}

func (conflict ConflictLoginError) Error() string {
	return fmt.Sprintf("login %v already exists", conflict.Login)
}

type AuthenticationError struct{}

func (auth AuthenticationError) Error() string {
	return "invalid login/password"
}
