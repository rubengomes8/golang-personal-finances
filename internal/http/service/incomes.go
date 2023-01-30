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

	var income models.Income
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

	var income models.Income
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
			ErrorMsg: "incomes with this id do not exist",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

// GetIncomeByID gets an income from the database that match the id provided
func (i *Incomes) GetIncomeByID(ctx *gin.Context) {

	paramID := ctx.Param("id")

	incomeID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting income id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	incomeViewRecord, err := i.Repository.GetIncomeByID(ctx, int64(incomeID))
	if err != nil {
		log.Printf("could not get income by id - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income with this id does not exist",
		})
		return
	}

	responseIncome := incomeViewToIncomeGetResponse(incomeViewRecord)

	ctx.JSON(http.StatusOK, responseIncome)
	ctx.Writer.Flush()
}

// DeleteIncome deletes an income from the database that match the id provided
func (i *Incomes) DeleteIncome(ctx *gin.Context) {

	paramID := ctx.Param("id")

	incomeID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting income id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	err = i.Repository.DeleteIncome(ctx, int64(incomeID))
	if err != nil {
		log.Printf("could not delete income with this id - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income with this id does not exist",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

func incomeViewToIncomeGetResponse(incomeView dbModels.IncomeView) models.Income {
	return models.Income{
		ID:          int(incomeView.ID),
		Value:       incomeView.Value,
		Date:        timeToStringDate(incomeView.Date),
		Category:    incomeView.Category,
		Card:        incomeView.Card,
		Description: incomeView.Description,
	}
}

// GetIncomesByCategory gets a list of incomes from the database that match the category provided
func (i *Incomes) GetIncomesByCategory(ctx *gin.Context) {

	paramCategory := ctx.Param("category")

	incomeViewRecords, err := i.Repository.GetIncomesByCategory(ctx, paramCategory)
	if err != nil {
		log.Printf("could not get incomes by category - category is %v - %v", paramCategory, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income category does not exist",
		})
		return
	}

	responseIncomes := incomeViewsToIncomesGetResponse(incomeViewRecords)

	ctx.JSON(http.StatusOK, responseIncomes)
	ctx.Writer.Flush()
}

// GetIncomesByCard gets a list of incomes from the database that match the card provided
func (i *Incomes) GetIncomesByCard(ctx *gin.Context) {

	paramCard := ctx.Param("card")

	incomeViewRecords, err := i.Repository.GetIncomesByCard(ctx, paramCard)
	if err != nil {
		log.Printf("could not get incomes by card - card is %v - %v", paramCard, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income card does not exist",
		})
		return
	}

	responseIncomes := incomeViewsToIncomesGetResponse(incomeViewRecords)

	ctx.JSON(http.StatusOK, responseIncomes)
	ctx.Writer.Flush()
}

// GetIncomesByDates gets a list of incomes from the database that match the dates' range provided
func (i *Incomes) GetIncomesByDates(ctx *gin.Context) {

	paramMinDate := ctx.Param("min_date")
	paramMaxDate := ctx.Param("max_date")

	minDate, err := dateStringToTime(paramMinDate)
	if err != nil {
		log.Printf("could not convert min date string to time - min date is %v - %v", paramMinDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse min date - must use YYYY-MM-DD date format",
		})
		return
	}

	maxDate, err := dateStringToTime(paramMaxDate)
	if err != nil {
		log.Printf("could not convert max date string to time - max date is %v - %v", paramMaxDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse max date - must use YYYY-MM-DD date format",
		})
		return
	}

	incomeViewRecords, err := i.Repository.GetIncomesByDates(ctx, minDate, maxDate)
	if err != nil {
		log.Printf("could not get incomes by dates - min_date is %v | max_date is %v - err: %v", paramMinDate, paramMaxDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not get incomes by dates",
		})
		return
	}

	responseIncomes := incomeViewsToIncomesGetResponse(incomeViewRecords)

	ctx.JSON(http.StatusOK, responseIncomes)
	ctx.Writer.Flush()
}

func incomeViewsToIncomesGetResponse(incomeViewRecords []dbModels.IncomeView) []models.Income {
	var responseIncomes []models.Income
	for _, inc := range incomeViewRecords {
		responseIncomes = append(responseIncomes, incomeViewToIncomeGetResponse(inc))
	}
	return responseIncomes
}
