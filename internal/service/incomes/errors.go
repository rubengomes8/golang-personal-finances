package service

import "errors"

var (
	ErrInvalidIncome                = errors.New("income is not valid")
	ErrCardNotFoundByName           = errors.New("could not get card by name")
	ErrIncomeCategoryNotFoundByName = errors.New("could not get income category by name")
	ErrCouldNotParseDate            = errors.New("could not parse date")
	ErrCouldNotInsertIncome         = errors.New("could not insert income")
	ErrCouldNotUpdateIncome         = errors.New("could not update income")
	ErrCouldNotDeleteIncome         = errors.New("could not delete income")
	ErrCouldNotGetIncome            = errors.New("could not get income")
	ErrCouldNotGetIncomesByDates    = errors.New("could not get incomes by dates")
)
