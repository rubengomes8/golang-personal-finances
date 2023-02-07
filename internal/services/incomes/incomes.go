package incomes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"

	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Incomes is the incomes service
type Incomes struct {
	Repository         repository.IncomeRepo
	CategoryRepository repository.IncomeCategoryRepo
	CardRepository     repository.CardRepo
}

// New creates a new Incomes service
func New(
	repo repository.IncomeRepo,
	categoryRepo repository.IncomeCategoryRepo,
	cardRepo repository.CardRepo,
) Incomes {
	return Incomes{
		Repository:         repo,
		CategoryRepository: categoryRepo,
		CardRepository:     cardRepo,
	}
}

// Add calls the repo to add an income
func (i Incomes) Add(ctx *gin.Context, income models.Income) (int, error) {

	card, err := i.CardRepository.GetCardByName(ctx, income.Card)
	if err != nil {
		log.Printf("could not get card by name: %v", err)
		return 0, ErrCardNotFound
	}

	category, err := i.CategoryRepository.GetIncomeCategoryByName(ctx, income.Category)
	if err != nil {
		log.Printf("could not get income category by name: %v", err)
		return 0, ErrCategoryNotFound

	}

	date, err := dateStringToTime(income.Date)
	if err != nil {
		log.Printf("error converting income date string to time - %v: %v", income.Date, err)
		return 0, ErrCouldNotParseTime
	}

	incomeRecord := dbModels.IncomeTable{
		Value:       income.Value,
		Date:        date,
		CategoryID:  category.ID,
		CardID:      card.ID,
		Description: income.Description,
	}

	id, err := i.Repository.InsertIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("could not insert income: %v", err)
		return 0, ErrCouldNotInsert
	}

	return int(id), nil
}

// Update calls the repo to update an income
func (i Incomes) Update(ctx *gin.Context, income models.Income) error {

	card, err := i.CardRepository.GetCardByName(ctx, income.Card)
	if err != nil {
		log.Printf("could not get card by name: %v", err)
		return ErrCardNotFound
	}

	category, err := i.CategoryRepository.GetIncomeCategoryByName(ctx, income.Category)
	if err != nil {
		log.Printf("could not get income category by name: %v", err)
		return ErrCategoryNotFound

	}

	date, err := dateStringToTime(income.Date)
	if err != nil {
		log.Printf("error converting income date string to time - %v: %v", income.Date, err)
		return ErrCouldNotParseTime
	}

	incomeRecord := dbModels.IncomeTable{
		Value:       income.Value,
		Date:        date,
		CategoryID:  category.ID,
		CardID:      card.ID,
		Description: income.Description,
	}

	_, err = i.Repository.UpdateIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("could not update income: %v", err)
		return ErrCouldNotUpdate
	}

	return nil
}

// Delete calls the repo to delete an income
func (i Incomes) Delete(ctx *gin.Context, id int64) error {

	err := i.Repository.DeleteIncome(ctx, id)
	if err != nil {
		log.Printf("could not delete income with this id - param id is %v - %v", id, err)
		return ErrCouldNotDelete
	}

	return nil
}

// GetByID calls the repo to get an income by id
func (i Incomes) GetByID(ctx *gin.Context, id int64) (dbModels.IncomeView, error) {

	incomeViewRecord, err := i.Repository.GetIncomeByID(ctx, int64(id))
	if err != nil {
		log.Printf("could not get income by id - param id is %v - %v", id, err)
		return dbModels.IncomeView{}, ErrNotFound
	}

	return incomeViewRecord, nil
}

// GetByCategory calls the repo to get incomes by category
func (i Incomes) GetByCategory(ctx *gin.Context, category string) ([]dbModels.IncomeView, error) {

	incomeViewRecords, err := i.Repository.GetIncomesByCategory(ctx, category)
	if err != nil {
		log.Printf("could not get incomes by category - category is %v - %v", category, err)
		return []dbModels.IncomeView{}, ErrCategoryNotFound
	}

	return incomeViewRecords, nil
}

// GetByCard calls the repo to get incomes by card
func (i Incomes) GetByCard(ctx *gin.Context, card string) ([]dbModels.IncomeView, error) {

	incomeViewRecords, err := i.Repository.GetIncomesByCard(ctx, card)
	if err != nil {
		log.Printf("could not get incomes by card - card is %v - %v", card, err)
		return []dbModels.IncomeView{}, ErrCategoryNotFound
	}

	return incomeViewRecords, nil
}

// GetByDates calls the repo to get incomes by dates
func (i Incomes) GetByDates(ctx *gin.Context, minDate, maxDate time.Time) ([]dbModels.IncomeView, error) {

	incomeViewRecords, err := i.Repository.GetIncomesByDates(ctx, minDate, maxDate)
	if err != nil {
		log.Printf("could not get incomes by dates - min_date is %v | max_date is %v - err: %v", minDate, maxDate, err)
		return []dbModels.IncomeView{}, ErrCategoryNotFound
	}

	return incomeViewRecords, nil
}

func dateStringToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
