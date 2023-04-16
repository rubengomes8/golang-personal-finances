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

// HandleCreateIncome handles a create income request.
// ShowEntity godoc
// @tags Incomes
// @Summary Creates a new income.
// @Description Endpoint to create an income.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.Income true "Create income request"
// @Success 201 {object} models.IncomeCreateResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/income [post]
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

// HandleUpdateIncome handles an update income request.
// ShowEntity godoc
// @tags Incomes
// @Summary Updates a new income.
// @Description Endpoint to update an income.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.Income true "Update income request"
// @Param id query string true "The income id"
// @Success 204 "No content"
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/income/{id} [put]
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

// HandleDeleteIncome handles an income delete request.
// ShowEntity godoc
// @tags Incomes
// @Summary Deletes a new income.
// @Description Endpoint to delete an income.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The income id"
// @Success 204 "No content"
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/income/{id} [delete]
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

// HandleGetByID handles a get income by id request.
// ShowEntity godoc
// @tags Incomes
// @Summary Gets an income by id.
// @Description Endpoint to get an income by its id.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The income id"
// @Success 200 {object} models.Income
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/income/{id} [get]
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

// HandleGetIncomesByCategory handles a get incomes by category request.
// ShowEntity godoc
// @tags Incomes
// @Summary Gets a list of incomes by category.
// @Description Endpoint to get a list of incomes by its category.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The income category"
// @Success 200 {object} []models.Income
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/incomes/category/{category} [get]
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

// HandleGetIncomesByCard handles a get incomes by card request.
// ShowEntity godoc
// @tags Incomes
// @Summary Gets a list of incomes by payment card.
// @Description Endpoint to get a list of incomes by payment card.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The payment card"
// @Success 200 {object} []models.Income
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/incomes/card/{card} [get]
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

// HandleGetIncomesByDates handles a get incomes by dates request.
// ShowEntity godoc
// @tags Incomes
// @Summary Gets a list of incomes by payment card.
// @Description Endpoint to get a list of incomes created on the provided range of dates.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param min_date query string true "The minimum date to consider"
// @Param max_date query string true "The maximum date to consider"
// @Success 200 {object} []models.Income
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/incomes/dates/{min_date}/{max_date} [get]
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
