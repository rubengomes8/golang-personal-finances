package service

import (
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
)

const (
	maxIncomeDescLen = 50
)

func validateNewIncome(income models.Income) error {

	if income.Value <= 0 {
		return ErrInvalidIncome
	}

	if len(income.Description) > maxIncomeDescLen {
		return ErrInvalidIncome
	}

	return nil
}
