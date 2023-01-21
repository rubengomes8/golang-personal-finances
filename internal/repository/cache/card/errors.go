package card

import "fmt"

// CardNotFoundByIDError error when a card is not found by id on the cache
type CardNotFoundByIDError struct {
	id int64
}

// Error is the string representation of CardNotFoundByIDError
func (cnfei CardNotFoundByIDError) Error() string {
	return fmt.Sprintf("error: card with id: %d was not found by id in the repository", cnfei.id)
}

// CardNotFoundByNameError error when a card is not found by name on the cache
type CardNotFoundByNameError struct {
	name string
}

// Error is the string representation of CardNotFoundByNameError
func (cnfen CardNotFoundByNameError) Error() string {
	return fmt.Sprintf("error: card with id: %s was not found by name in the repository", cnfen.name)
}

// CardAlreadyExistsError error when a card already exists on the cache
type CardAlreadyExistsError struct {
	id int64
}

// Error is the string representation of CardAlreadyExistsError
func (caee CardAlreadyExistsError) Error() string {
	return fmt.Sprintf("error: card with id: %d already exists in the repository", caee.id)
}
