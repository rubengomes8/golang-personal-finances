package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Incomes handles the incomes http requests
type Incomes struct {
	Repository         repository.IncomeRepo
	CategoryRepository repository.IncomeCategoryRepo
	CardRepository     repository.CardRepo
}

// NewIncomes creates a new Incomes service
func NewIncomes(
	repo repository.IncomeRepo,
	categoryRepo repository.IncomeCategoryRepo,
	cardRepo repository.CardRepo,
) (Incomes, error) {
	return Incomes{
		Repository:         repo,
		CategoryRepository: categoryRepo,
		CardRepository:     cardRepo,
	}, nil
}

// CreateIncome creates an income on the database
func (i *Incomes) CreateIncome(ctx *gin.Context) {

	var income models.IncomeCreateRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&income)
	if err != nil {
		log.Printf("could not decode create income body: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not decode income",
		})
		return
	}

	card, err := i.CardRepository.GetCardByName(ctx, income.Card)
	if err != nil {
		log.Printf("could not get card by name: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "card does not exist",
		})
		return
	}

	category, err := i.CategoryRepository.GetIncomeCategoryByName(ctx, income.Category)
	if err != nil {
		log.Printf("could not get income category by name: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income category does not exist",
		})
		return
	}

	date, err := dateStringToTime(income.Date)
	if err != nil {
		log.Printf("error converting income date string to time - %v: %v", income.Date, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse date - must use YYYY-MM-DD date format",
		})
		return
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
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: "could not create income",
		})
		return
	}

	ctx.JSON(http.StatusCreated, &models.IncomeCreateResponse{ID: int(id)})
	ctx.Writer.Flush()
}

// UpdateIncome updates an income on the database
func (i *Incomes) UpdateIncome(ctx *gin.Context) {

	var income models.IncomeCreateRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&income)
	if err != nil {
		log.Printf("could not decode update income body: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not decode income",
		})
		return
	}

	card, err := i.CardRepository.GetCardByName(ctx, income.Card)
	if err != nil {
		log.Printf("could not get card by name: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "card does not exist",
		})
		return
	}

	category, err := i.CategoryRepository.GetIncomeCategoryByName(ctx, income.Category)
	if err != nil {
		log.Printf("could not get income category by name: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income category does not exist",
		})
		return
	}

	date, err := dateStringToTime(income.Date)
	if err != nil {
		log.Printf("error converting date string to time - %v: %v", income.Date, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse date - must use YYYY-MM-DD date format",
		})
		return
	}

	paramID := ctx.Param("id")

	incomeID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting income id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	incomeRecord := dbModels.IncomeTable{
		ID:          int64(incomeID),
		Value:       income.Value,
		Date:        date,
		CategoryID:  category.ID,
		CardID:      card.ID,
		Description: income.Description,
	}

	_, err = i.Repository.UpdateIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("could not update income with param id = %v: %v", paramID, err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: "incomes with this id does not exist",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}
