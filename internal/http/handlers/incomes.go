package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/services/incomes"

	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Incomes handles the incomes http requests
type Incomes struct {
	service incomes.Incomes
}

// NewIncomes creates a new Incomes service
func NewIncomes(
	service incomes.Incomes,
) (Incomes, error) {
	return Incomes{
		service: service,
	}, nil
}

// HandleCreateIncome handles an income create request
func (i *Incomes) HandleCreateIncome(ctx *gin.Context) {

	var income models.Income
	err := json.NewDecoder(ctx.Request.Body).Decode(&income)
	if err != nil {
		log.Printf("could not decode create income body: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not decode income",
		})
		return
	}

	id, err := i.service.Add(ctx, income)
	if err != nil {
		log.Printf("could not add income: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not add income",
		})
		return
	}

	ctx.JSON(http.StatusCreated, &models.IncomeCreateResponse{ID: int(id)})
	ctx.Writer.Flush()
}

// HandleUpdateIncome handles an income update request
func (i *Incomes) HandleUpdateIncome(ctx *gin.Context) {

	var income models.Income
	err := json.NewDecoder(ctx.Request.Body).Decode(&income)
	if err != nil {
		log.Printf("could not decode update income body: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not decode income",
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

	income.ID = incomeID

	err = i.service.Update(ctx, income)
	if err != nil {
		log.Printf("could not update income: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not update income",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

// HandleGetIncomeByID handles a get income by id request
func (i *Incomes) HandleGetIncomeByID(ctx *gin.Context) {

	paramID := ctx.Param("id")

	incomeID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting income id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	incomeViewRecord, err := i.service.GetByID(ctx, int64(incomeID))
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

// HandleDeleteIncome handles an income delete request
func (i *Incomes) HandleDeleteIncome(ctx *gin.Context) {

	paramID := ctx.Param("id")

	incomeID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting income id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	err = i.service.Delete(ctx, int64(incomeID))
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

// HandleGetIncomesByCategory handles a get incomes by category request
func (i *Incomes) HandleGetIncomesByCategory(ctx *gin.Context) {

	paramCategory := ctx.Param("category")

	incomeViewRecords, err := i.service.GetByCategory(ctx, paramCategory)
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

// HandleGetIncomesByCard handles a get incomes by card request
func (i *Incomes) HandleGetIncomesByCard(ctx *gin.Context) {

	paramCard := ctx.Param("card")

	incomeViewRecords, err := i.service.GetByCard(ctx, paramCard)
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

// HandleGetIncomesByDates handles a get incomes by dates request
func (i *Incomes) HandleGetIncomesByDates(ctx *gin.Context) {

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

	incomeViewRecords, err := i.service.GetByDates(ctx, minDate, maxDate)
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
