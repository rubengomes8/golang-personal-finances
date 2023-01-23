package cache

import "fmt"

// UserNotFoundByNameError error when a user is not found by name on the cache
type UserNotFoundByNameError struct {
	username string
}

// Error is the string representation of UserNotFoundByNameError
func (unfe UserNotFoundByNameError) Error() string {
	return fmt.Sprintf("error: user with username: %s was not found by username in the repository", unfe.username)
}

// UserAlreadyExistsError error when a user already exists on the cache
type UserAlreadyExistsError struct {
	username string
}

// Error is the string representation of UserAlreadyExistsError
func (uae UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("error: user with username: %s already exists in the repository", uae.username)
}
