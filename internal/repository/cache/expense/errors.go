package expense

import "fmt"

/* CATEGORY */

// CategoryNotFoundByIDError error when a category is not found by id on the cache
type CategoryNotFoundByIDError struct {
	id int64
}

// Error is the string representation of CategoryNotFoundByIDError
func (cnfie CategoryNotFoundByIDError) Error() string {
	return fmt.Sprintf("error: category with id: %d was not found by id in the repository", cnfie.id)
}

// CategoryNotFoundByNameError error when a category is not found by name on the cache
type CategoryNotFoundByNameError struct {
	name string
}

// Error is the string representation of CategoryNotFoundByNameError
func (cnfne CategoryNotFoundByNameError) Error() string {
	return fmt.Sprintf("error: category with id: %s was not found by name in the repository", cnfne.name)
}

// CategoryNotFoundByNameError error when a category already exists on the cache
type CategoryAlreadyExistsError struct {
	id int64
}

// Error is the string representation of CategoryAlreadyExistsError
func (caee CategoryAlreadyExistsError) Error() string {
	return fmt.Sprintf("error: category with id: %d already exists in the repository", caee.id)
}

/* SUBCATEGORY*/

// SubCategoryNotFoundByIDError error when a subcategory is not found by id on the cache
type SubCategoryNotFoundByIDError struct {
	id int64
}

// Error is the string representation of SubCategoryNotFoundByIDError
func (snfie SubCategoryNotFoundByIDError) Error() string {
	return fmt.Sprintf("error: subcategory with id: %d was not found by id in the repository", snfie.id)
}

// SubCategoryNotFoundByNameError error when a subcategory is not found by name on the cache
type SubCategoryNotFoundByNameError struct {
	name string
}

// Error is the string representation of SubCategoryNotFoundByNameError
func (snfne SubCategoryNotFoundByNameError) Error() string {
	return fmt.Sprintf("error: subcategory with id: %s was not found by name in the repository", snfne.name)
}

// SubCategoryAlreadyExistsError error when a subcategory already exists on the cache
type SubCategoryAlreadyExistsError struct {
	id int64
}

// Error is the string representation of SubCategoryAlreadyExistsError
func (saee SubCategoryAlreadyExistsError) Error() string {
	return fmt.Sprintf("error: subcategory with id: %d already exists in the repository", saee.id)
}

/* EXPENSE */

// NotFoundByIDError error when an expense is not found by id on the cache
type NotFoundByIDError struct {
	id int64
}

// Error is the string representation of NotFoundByIDError
func (nfie NotFoundByIDError) Error() string {
	return fmt.Sprintf("error: expense with id: %d was not found by id in the repository", nfie.id)
}

// NotFoundByNameError error when a expense is not found by name on the cache
type NotFoundByNameError struct {
	name string
}

// Error is the string representation of SubCategoryAlreadyExistsError
func (nfne NotFoundByNameError) Error() string {
	return fmt.Sprintf("error: expense with name: %s was not found by id in the repository", nfne.name)
}

// GettingCardByIDError error when a trying to get a card
type GettingCardByIDError struct {
	id int64
}

// Error is the string representation of GettingCardByIDError
func (gcie GettingCardByIDError) Error() string {
	return fmt.Sprintf("error: trying to get the card with id: %d", gcie.id)
}

// GettingCardByNameError error when a trying to get a card
type GettingCardByNameError struct {
	name string
}

// Error is the string representation of GettingCardByNameError
func (gcne GettingCardByNameError) Error() string {
	return fmt.Sprintf("error: trying to get the card with name: %s", gcne.name)
}

// GettingSubCategoryByIDError error when a trying to get a subcategory
type GettingSubCategoryByIDError struct {
	id int64
}

// Error is the string representation of GettingSubCategoryByIDError
func (gscie GettingSubCategoryByIDError) Error() string {
	return fmt.Sprintf("error: trying to get the subcategory with id: %d", gscie.id)
}

// GettingSubCategoryByNameError error when a trying to get a subcategory
type GettingSubCategoryByNameError struct {
	name string
}

// Error is the string representation of Name
func (gscne GettingSubCategoryByNameError) Error() string {
	return fmt.Sprintf("error: trying to get the subcategory name id: %s", gscne.name)
}

// GettingCategoryByIDError error when a trying to get a category
type GettingCategoryByIDError struct {
	id int64
}

// Error is the string representation of GettingCategoryByIDError
func (gcie GettingCategoryByIDError) Error() string {
	return fmt.Sprintf("error: trying to get the category with id: %d", gcie.id)
}

// GettingCategoryByNameError error when a trying to get a category
type GettingCategoryByNameError struct {
	name string
}

// Error is the string representation of GettingCategoryByNameError
func (gcne GettingCategoryByNameError) Error() string {
	return fmt.Sprintf("error: trying to get the category with name: %s", gcne.name)
}
