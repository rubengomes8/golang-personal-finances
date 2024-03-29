package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
	"github.com/rubengomes8/golang-personal-finances/internal/utils"
)

// Expenses handles the expenses http requests
type Expenses struct {
	Repository            repository.ExpenseRepo
	SubCategoryRepository repository.ExpenseSubCategoryRepo
	CardRepository        repository.CardRepo
}

// NewExpenses creates a new Expenses service
func NewExpenses(
	expRepo repository.ExpenseRepo,
	expSubCatRepo repository.ExpenseSubCategoryRepo,
	cardRepo repository.CardRepo,
) Expenses {
	return Expenses{
		Repository:            expRepo,
		SubCategoryRepository: expSubCatRepo,
		CardRepository:        cardRepo,
	}
}

// CreateExpense is used to create a new expense.
// ShowEntity godoc
// @tags Expenses
// @Summary Creates a new expense.
// @Description Endpoint to create an expense.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body models.ExpenseCreateRequest true "Create expense request"
// @Success 201 {object} models.ExpenseCreateResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/expense [post]
func (e *Expenses) CreateExpense(ctx *gin.Context) {

	var expense models.ExpenseCreateRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&expense)
	if err != nil {
		log.Printf("could not decode create expense body: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not decode expense",
		})
		return
	}

	expSubCategory, card, err := e.getExpenseSubcategoryAndCardIDByNames(ctx, expense.SubCategory, expense.Card)
	if err != nil {
		log.Printf("could not get expense subcategory and card ids by names: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "subcategory or card does not exist",
		})
		return
	}

	date, err := utils.DateStringToTime(expense.Date)
	if err != nil {
		log.Printf("error converting date string to time - %v: %v", expense.Date, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse date - must use YYYY-MM-DD date format",
		})
		return
	}

	expenseRecord := dbModels.ExpenseTable{
		Value:         expense.Value,
		Date:          date,
		SubCategoryID: expSubCategory.ID,
		CardID:        card.ID,
		Description:   expense.Description,
	}

	id, err := e.Repository.InsertExpense(ctx, expenseRecord)
	if err != nil {
		log.Printf("could not insert expense: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: "could not create expense",
		})
		return
	}

	ctx.JSON(http.StatusCreated, &models.ExpenseCreateResponse{ID: int(id)})
	ctx.Writer.Flush()
}

// UpdateExpense updates an expense on the database.
// ShowEntity godoc
// @tags Expenses
// @Summary Updates an existing expense.
// @Description Endpoint to update an expense.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The expense id"
// @Param body body models.ExpenseCreateRequest true "Update expense request"
// @Success 204 "No content"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/expense/{id} [put]
func (e *Expenses) UpdateExpense(ctx *gin.Context) {

	var expense models.ExpenseCreateRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&expense)
	if err != nil {
		log.Printf("could not decode update expense body: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not decode expense",
		})
		return
	}

	expSubCategory, card, err := e.getExpenseSubcategoryAndCardIDByNames(ctx, expense.SubCategory, expense.Card)
	if err != nil {
		log.Printf("could not get expense subcategory and card ids by names: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "subcategory or card does not exist",
		})
	}

	date, err := utils.DateStringToTime(expense.Date)
	if err != nil {
		log.Printf("error converting date string to time - %v: %v", expense.Date, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse date - must use YYYY-MM-DD date format",
		})
		return
	}

	paramID := ctx.Param("id")

	expenseID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting expense id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	expenseRecord := dbModels.ExpenseTable{
		ID:            int64(expenseID),
		Value:         expense.Value,
		Date:          date,
		SubCategoryID: expSubCategory.ID,
		CardID:        card.ID,
		Description:   expense.Description,
	}

	_, err = e.Repository.UpdateExpense(ctx, expenseRecord)
	if err != nil {
		log.Printf("could not update expense with param id = %v: %v", paramID, err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: "expense with this id does not exist",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

// GetExpenseByID gets an expense from the database that match the id provided.
// ShowEntity godoc
// @tags Expenses
// @Summary Gets an expense by its id.
// @Description Endpoint to get an expense by id.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The expense id"
// @Success 201 {object} models.ExpenseCreateRequest
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/expense/{id} [get]
func (e *Expenses) GetExpenseByID(ctx *gin.Context) {

	paramID := ctx.Param("id")

	expenseID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting expense id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	expenseViewRecord, err := e.Repository.GetExpenseByID(ctx, int64(expenseID))
	if err != nil {
		log.Printf("could not get expense by id - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "expense with this id does not exist",
		})
		return
	}

	responseExpense := expenseViewToExpenseGetResponse(expenseViewRecord)

	ctx.JSON(http.StatusOK, responseExpense)
	ctx.Writer.Flush()
}

// GetExpensesByCategory gets a list of expenses from the database that match the category provided.
// ShowEntity godoc
// @tags Expenses
// @Summary Gets a list of expenses by a category.
// @Description Endpoint to get a list of expenses by category.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category query string true "The expense category"
// @Success 201 {object} []models.ExpenseCreateRequest
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/expenses/category/{category} [get]
func (e *Expenses) GetExpensesByCategory(ctx *gin.Context) {

	paramCategory := ctx.Param("category")

	expenseViewRecords, err := e.Repository.GetExpensesByCategory(ctx, paramCategory)
	if err != nil {
		log.Printf("could not get expenses by category - category is %v - %v", paramCategory, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "expense category does not exist",
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// GetExpensesBySubCategory gets a list of expenses from the database that match the subcategory provided.
// ShowEntity godoc
// @tags Expenses
// @Summary Gets a list of expenses by subcategory.
// @Description Endpoint to get a list of expenses by subcategory.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category query string true "The expense subcategory"
// @Success 201 {object} []models.ExpenseCreateRequest
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/expenses/subcategory/{sub_category} [get]
func (e *Expenses) GetExpensesBySubCategory(ctx *gin.Context) {

	paramSubCategory := ctx.Param("sub_category")

	expenseViewRecords, err := e.Repository.GetExpensesBySubCategory(ctx, paramSubCategory)
	if err != nil {
		log.Printf("could not get expenses by subcategory - category is %v - %v", paramSubCategory, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "expense subcategory does not exist",
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// GetExpensesByCard gets a list of expenses from the database that match the card provided.
// ShowEntity godoc
// @tags Expenses
// @Summary Gets a list of expenses by card.
// @Description Endpoint to get a list of expense by card.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category query string true "The card"
// @Success 201 {object} []models.ExpenseCreateRequest
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/expenses/card/{card} [get]
func (e *Expenses) GetExpensesByCard(ctx *gin.Context) {

	paramCard := ctx.Param("card")

	expenseViewRecords, err := e.Repository.GetExpensesByCard(ctx, paramCard)
	if err != nil {
		log.Printf("could not get expenses by card - card is %v - %v", paramCard, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "expense card does not exist",
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// GetExpensesByDates gets a list of expenses from the database that match the dates' range provided
// ShowEntity godoc
// @tags Expenses
// @Summary Gets a list of expenses created on a range of dates.
// @Description Endpoint to get a list of expenses created on the provided range of dates.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param min_date query string true "The minimum date to consider"
// @Param max_date query string true "The maximum date to consider"
// @Success 201 {object} []models.ExpenseCreateRequest
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/expenses/dates/{min_date}/{max_date} [get]
func (e *Expenses) GetExpensesByDates(ctx *gin.Context) {

	paramMinDate := ctx.Param("min_date")
	paramMaxDate := ctx.Param("max_date")

	minDate, err := utils.DateStringToTime(paramMinDate)
	if err != nil {
		log.Printf("could not convert min date string to time - min date is %v - %v", paramMinDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse min date - must use YYYY-MM-DD date format",
		})
		return
	}

	maxDate, err := utils.DateStringToTime(paramMaxDate)
	if err != nil {
		log.Printf("could not convert max date string to time - max date is %v - %v", paramMaxDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not parse max date - must use YYYY-MM-DD date format",
		})
		return
	}

	expenseViewRecords, err := e.Repository.GetExpensesByDates(ctx, minDate, maxDate)
	if err != nil {
		log.Printf("could not get expenses by dates - min_date is %v | max_date is %v - err: %v", paramMinDate, paramMaxDate, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not get expenses by dates",
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// DeleteExpense deletes an expense from the database that match the id provided.
// ShowEntity godoc
// @tags Expenses
// @Summary Deletes an expense by its id.
// @Description Endpoint to delete an expense by id.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "The expense id"
// @Success 204 "No Content"
// @Failure 400 {object} models.ErrorResponse
// @Router /v1/expense/{id} [delete]
func (e *Expenses) DeleteExpense(ctx *gin.Context) {

	paramID := ctx.Param("id")

	expenseID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("error converting expense id to int - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "id parameter must be an integer",
		})
		return
	}

	err = e.Repository.DeleteExpense(ctx, int64(expenseID))
	if err != nil {
		log.Printf("could not delete expense with this id - param id is %v - %v", paramID, err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "expense with this id does not exist",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

func (e *Expenses) getExpenseSubcategoryAndCardIDByNames(
	ctx context.Context,
	subCategory, card string,
) (dbModels.ExpenseSubCategoryTable, dbModels.CardTable, error) {
	subCategoryModel, err := e.SubCategoryRepository.GetExpenseSubCategoryByName(ctx, subCategory)
	if err != nil {
		return dbModels.ExpenseSubCategoryTable{}, dbModels.CardTable{}, fmt.Errorf("could not get expense sub category by name: %v", err)
	}

	cardModel, err := e.CardRepository.GetCardByName(ctx, card)
	if err != nil {
		return dbModels.ExpenseSubCategoryTable{}, dbModels.CardTable{}, fmt.Errorf("could not get expense card by name: %v", err)
	}

	return subCategoryModel, cardModel, nil
}

func expenseViewToExpenseGetResponse(expenseView dbModels.ExpenseView) models.ExpenseCreateRequest {
	return models.ExpenseCreateRequest{
		ID:          int(expenseView.ID),
		Value:       expenseView.Value,
		Date:        utils.TimeToStringDate(expenseView.Date),
		SubCategory: expenseView.SubCategory,
		Card:        expenseView.Card,
		Description: expenseView.Description,
	}
}

func expensesViewToExpensesGetResponse(expenseViewRecords []dbModels.ExpenseView) []models.ExpenseCreateRequest {
	responseExpenses := []models.ExpenseCreateRequest{}
	for _, exp := range expenseViewRecords {
		responseExpenses = append(responseExpenses, expenseViewToExpenseGetResponse(exp))
	}
	return responseExpenses
}
