package service

import "github.com/rubengomes8/golang-personal-finances/internal/repository"

type IncomeConfiguration func(service *Incomes) error

func NewIncomesWithConfiguration(cfgs ...IncomeConfiguration) (*Incomes, error) {
	service := &Incomes{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithIncomesRepository(incomeRepo repository.IncomeRepo) IncomeConfiguration {
	return func(service *Incomes) error {
		service.repo = incomeRepo
		return nil
	}
}

func WithCategoryRepository(categoryRepo repository.IncomeCategoryRepo) IncomeConfiguration {
	return func(service *Incomes) error {
		service.categoryRepo = categoryRepo
		return nil
	}
}

func WithCardRepository(cardrepo repository.CardRepo) IncomeConfiguration {
	return func(service *Incomes) error {
		service.cardRepo = cardrepo
		return nil
	}
}
