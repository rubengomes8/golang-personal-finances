package service

import (
	"context"
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/utils"

	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

type Incomes struct {
	repo         repository.IncomeRepo
	categoryRepo repository.IncomeCategoryRepo
	cardRepo     repository.CardRepo
}

// NewIncomes creates a new Incomes service
func NewIncomes(
	repo repository.IncomeRepo,
	categoryRepo repository.IncomeCategoryRepo,
	cardRepo repository.CardRepo,
) Incomes {
	return Incomes{
		repo:         repo,
		categoryRepo: categoryRepo,
		cardRepo:     cardRepo,
	}
}

// Create is the create income usecase
func (i Incomes) Create(ctx context.Context, income models.Income) (int, error) {

	err := validateNewIncome(income)
	if err != nil {
		return 0, ErrInvalidIncome
	}

	card, err := i.cardRepo.GetCardByName(ctx, income.Card)
	if err != nil {
		log.Printf("could not get card by name: %v", err)
		return 0, ErrCardNotFoundByName
	}

	category, err := i.categoryRepo.GetIncomeCategoryByName(ctx, income.Category)
	if err != nil {
		log.Printf("could not get income category by name: %v", err)
		return 0, ErrIncomeCategoryNotFoundByName
	}

	date, err := utils.DateStringToTime(income.Date)
	if err != nil {
		log.Printf("error converting income date string to time - %v: %v", income.Date, err)
		return 0, ErrCouldNotParseDate
	}

	incomeRecord := dbModels.IncomeTable{
		Value:       income.Value,
		Date:        date,
		CategoryID:  category.ID,
		CardID:      card.ID,
		Description: income.Description,
	}

	id, err := i.repo.InsertIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("could not insert income: %v", err)
		return 0, ErrCouldNotInsertIncome
	}

	return int(id), nil
}

// Update is the update income usecase
func (i Incomes) Update(ctx context.Context, income models.Income) error {

	err := validateNewIncome(income)
	if err != nil {
		return ErrInvalidIncome
	}

	card, err := i.cardRepo.GetCardByName(ctx, income.Card)
	if err != nil {
		log.Printf("could not get card by name: %v", err)
		return ErrCardNotFoundByName
	}

	category, err := i.categoryRepo.GetIncomeCategoryByName(ctx, income.Category)
	if err != nil {
		log.Printf("could not get income category by name: %v", err)
		return ErrIncomeCategoryNotFoundByName
	}

	date, err := utils.DateStringToTime(income.Date)
	if err != nil {
		log.Printf("error converting income date string to time - %v: %v", income.Date, err)
		return ErrCouldNotParseDate
	}

	incomeRecord := dbModels.IncomeTable{
		Value:       income.Value,
		Date:        date,
		CategoryID:  category.ID,
		CardID:      card.ID,
		Description: income.Description,
	}

	_, err = i.repo.UpdateIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("could not update income: %v", err)
		return ErrCouldNotInsertIncome
	}

	return nil
}

func (i Incomes) Delete(ctx context.Context, id int) error {

	err := i.repo.DeleteIncome(ctx, int64(id))
	if err != nil {
		log.Printf("could not delete income with this id - param id is %v - %v", id, err)
		return ErrCouldNotDeleteIncome
	}

	return nil
}

func (i Incomes) GetByID(ctx context.Context, id int) (models.Income, error) {

	incomeViewRecord, err := i.repo.GetIncomeByID(ctx, int64(id))
	if err != nil {
		log.Printf("could not get income by id - param id is %v - %v", id, err)
		return models.Income{}, ErrCouldNotGetIncome
	}

	return mapIncomeViewToIncome(incomeViewRecord), nil

}

func (i Incomes) GetAllByCard(ctx context.Context, card string) ([]models.Income, error) {

	incomeViewRecords, err := i.repo.GetIncomesByCard(ctx, card)
	if err != nil {
		log.Printf("could not get incomes by card - card is %v - %v", card, err)
		return []models.Income{}, ErrCardNotFoundByName
	}

	return mapIncomeViewsToIncomes(incomeViewRecords), nil
}

func (i Incomes) GetAllByCategory(ctx context.Context, category string) ([]models.Income, error) {

	incomeViewRecords, err := i.repo.GetIncomesByCategory(ctx, category)
	if err != nil {
		log.Printf("could not get incomes by category - category is %v - %v", category, err)
		return []models.Income{}, ErrIncomeCategoryNotFoundByName
	}

	return mapIncomeViewsToIncomes(incomeViewRecords), nil
}

func (i Incomes) GetAllByDates(ctx context.Context, paramMinDate, paramMaxDate string) ([]models.Income, error) {

	minDate, err := utils.DateStringToTime(paramMinDate)
	if err != nil {
		log.Printf("could not convert min date string to time - min date is %v - %v", minDate, err)
		return []models.Income{}, ErrCouldNotParseDate
	}

	maxDate, err := utils.DateStringToTime(paramMaxDate)
	if err != nil {
		log.Printf("could not convert max date string to time - max date is %v - %v", paramMaxDate, err)
		return []models.Income{}, ErrCouldNotGetIncomesByDates
	}

	incomeViewRecords, err := i.repo.GetIncomesByDates(ctx, minDate, maxDate)
	if err != nil {
		log.Printf("could not get incomes by dates - min_date is %v | max_date is %v - err: %v", paramMinDate, paramMaxDate, err)
		return []models.Income{}, ErrCouldNotGetIncomesByDates
	}

	return mapIncomeViewsToIncomes(incomeViewRecords), nil

}

func mapIncomeViewToIncome(incomeView dbModels.IncomeView) models.Income {
	return models.Income{
		ID:          int(incomeView.ID),
		Value:       incomeView.Value,
		Date:        utils.TimeToStringDate(incomeView.Date),
		Category:    incomeView.Category,
		Card:        incomeView.Card,
		Description: incomeView.Description,
	}
}

func mapIncomeViewsToIncomes(incomeViewRecords []dbModels.IncomeView) []models.Income {
	responseIncomes := []models.Income{}
	for _, incomeView := range incomeViewRecords {
		responseIncomes = append(responseIncomes, mapIncomeViewToIncome(incomeView))
	}
	return responseIncomes
}
