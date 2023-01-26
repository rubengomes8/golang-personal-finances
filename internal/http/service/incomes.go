package service

import "github.com/rubengomes8/golang-personal-finances/internal/repository"

// Incomes handles the incomes http requests
type Incomes struct {
	ExpensesRepository            repository.ExpenseRepo
	ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
	CardRepository                repository.CardRepo
}

// NewIncomes creates a new Incomes service
func NewIncomes(
	cardRepo repository.CardRepo,
) (Incomes, error) {
	return Incomes{
		CardRepository: cardRepo,
	}, nil
}
