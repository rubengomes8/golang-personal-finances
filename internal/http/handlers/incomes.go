package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/service"
)

// Incomes handles the incomes http requests
type Incomes struct {
	service service.Incomes
}

// NewIncomes creates a new Incomes service
func NewIncomes(
	service service.Incomes,
) Incomes {
	return Incomes{
		service: service,
	}
}

// HandleCreateIncome handles a create income request
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

	incomeID, err := i.service.Create(ctx, income)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not create income",
		})
		return
	}

	ctx.JSON(http.StatusCreated, &models.IncomeCreateResponse{ID: incomeID})
	ctx.Writer.Flush()
}

// HandleUpdateIncome handles an update income request
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

	err = i.service.Update(ctx, income)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not update income",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
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

	err = i.service.Delete(ctx, incomeID)
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

// HandleGetByID handles a get income by id request
func (i *Incomes) HandleGetByID(ctx *gin.Context) {

	paramID := ctx.Param("id")

	incomeID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting income id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	income, err := i.service.GetByID(ctx, incomeID)
	if err != nil {
		log.Printf("could not get income by id - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income with this id does not exist",
		})
		return
	}

	ctx.JSON(http.StatusOK, income)
	ctx.Writer.Flush()
}

// HandleGetIncomesByCategory handles a get incomes by category request
func (i *Incomes) HandleGetIncomesByCategory(ctx *gin.Context) {

	paramCategory := ctx.Param("category")

	incomes, err := i.service.GetAllByCategory(ctx, paramCategory)
	if err != nil {
		log.Printf("could not get incomes by category - category is %v - %v", paramCategory, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income category does not exist",
		})
		return
	}

	ctx.JSON(http.StatusOK, incomes)
	ctx.Writer.Flush()
}

// HandleGetIncomesByCard handles a get incomes by card request
func (i *Incomes) HandleGetIncomesByCard(ctx *gin.Context) {

	paramCard := ctx.Param("card")

	incomes, err := i.service.GetAllByCard(ctx, paramCard)
	if err != nil {
		log.Printf("could not get incomes by card - card is %v - %v", paramCard, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "income card does not exist",
		})
		return
	}

	ctx.JSON(http.StatusOK, incomes)
	ctx.Writer.Flush()
}

// HandleGetIncomesByDates handles a get incomes by dates request
func (i *Incomes) HandleGetIncomesByDates(ctx *gin.Context) {

	paramMinDate := ctx.Param("min_date")
	paramMaxDate := ctx.Param("max_date")

	incomes, err := i.service.GetAllByDates(ctx, paramMinDate, paramMaxDate)
	if err != nil {
		log.Printf("could not get incomes by dates - min_date is %v | max_date is %v - err: %v", paramMinDate, paramMaxDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not get incomes by dates",
		})
		return
	}

	ctx.JSON(http.StatusOK, incomes)
	ctx.Writer.Flush()
}
