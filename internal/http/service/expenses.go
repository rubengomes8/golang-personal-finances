package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// Expenses handles the expenses http requests
type Expenses struct {
	ExpensesRepository            repository.ExpenseRepo
	ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
	CardRepository                repository.CardRepo
}

// NewExpenses creates a new Expenses service
func NewExpenses(
	expRepo repository.ExpenseRepo,
	expSubCatRepo repository.ExpenseSubCategoryRepo,
	cardRepo repository.CardRepo,
) (Expenses, error) {
	return Expenses{
		ExpensesRepository:            expRepo,
		ExpensesSubCategoryRepository: expSubCatRepo,
		CardRepository:                cardRepo,
	}, nil
}

// CreateExpense creates an expense on the database
func (e *Expenses) CreateExpense(ctx *gin.Context) {

	var expense models.ExpenseCreateRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not decode expense: %v", err),
		})
		return
	}

	expSubCategory, card, err := e.getExpenseSubcategoryAndCardIDByNames(ctx, expense.SubCategory, expense.Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("unknown subcategory or card: %v", err),
		})
		return
	}

	date, err := dateStringToTime(expense.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not parse date (should use YYYY-MM-DD format): %v", err),
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

	id, err := e.ExpensesRepository.InsertExpense(ctx, expenseRecord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not insert expense: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &models.ExpenseCreateResponse{ID: int(id)})
	ctx.Writer.Flush()
}

// UpdateExpense updates an expense on the database
func (e *Expenses) UpdateExpense(ctx *gin.Context) {

	var expense models.ExpenseCreateRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not decode expense: %v", err),
		})
		return
	}

	expSubCategory, card, err := e.getExpenseSubcategoryAndCardIDByNames(ctx, expense.SubCategory, expense.Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("unknown subcategory or card: %v", err),
		})
		return
	}

	date, err := dateStringToTime(expense.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not parse date (should use YYYY-MM-DD format): %v", err),
		})
		return
	}

	paramID := ctx.Param("id")

	expenseID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("id parameter must be an integer: %v", err),
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

	_, err = e.ExpensesRepository.UpdateExpense(ctx, expenseRecord)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not update expense: %v", err),
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

// GetExpenseByID gets an expense from the database that match the id provided
func (e *Expenses) GetExpenseByID(ctx *gin.Context) {

	paramID := ctx.Param("id")

	expenseID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("id parameter must be an integer: %v", err),
		})
		return
	}

	expenseViewRecord, err := e.ExpensesRepository.GetExpenseByID(ctx, int64(expenseID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not get expense by id: %v", err),
		})
		return
	}

	responseExpense := expenseViewToExpenseGetResponse(expenseViewRecord)

	ctx.JSON(http.StatusOK, responseExpense)
	ctx.Writer.Flush()
}

// GetExpensesByCategory gets a list of expenses from the database that match the category provided
func (e *Expenses) GetExpensesByCategory(ctx *gin.Context) {

	paramCategory := ctx.Param("category")

	expenseViewRecords, err := e.ExpensesRepository.GetExpensesByCategory(ctx, paramCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not get expenses by category: %v", err),
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// GetExpensesBySubCategory gets a list of expenses from the database that match the subcategory provided
func (e *Expenses) GetExpensesBySubCategory(ctx *gin.Context) {

	paramSubCategory := ctx.Param("sub_category")

	expenseViewRecords, err := e.ExpensesRepository.GetExpensesBySubCategory(ctx, paramSubCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not get expenses by subcategory: %v", err),
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// GetExpensesByCard gets a list of expenses from the database that match the card provided
func (e *Expenses) GetExpensesByCard(ctx *gin.Context) {

	paramCard := ctx.Param("card")

	expenseViewRecords, err := e.ExpensesRepository.GetExpensesByCard(ctx, paramCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not get expenses by card: %v", err),
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// GetExpensesByDates gets a list of expenses from the database that match the dates' range provided
func (e *Expenses) GetExpensesByDates(ctx *gin.Context) {

	paramMinDate := ctx.Param("min_date")
	paramMaxDate := ctx.Param("max_date")

	minDate, err := dateStringToTime(paramMinDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not parse min_date (should use YYYY-MM-DD format): %v", err),
		})
		return
	}

	maxDate, err := dateStringToTime(paramMaxDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not parse max_date (should use YYYY-MM-DD format): %v", err),
		})
		return
	}

	expenseViewRecords, err := e.ExpensesRepository.GetExpensesByDates(ctx, minDate, maxDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not get expenses by dates: %v", err),
		})
		return
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	ctx.JSON(http.StatusOK, responseExpenses)
	ctx.Writer.Flush()
}

// DeleteExpense deletes an expense from the database that match the id provided
func (e *Expenses) DeleteExpense(ctx *gin.Context) {

	paramID := ctx.Param("id")

	expenseID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("id parameter must be an integer: %v", err),
		})
		return
	}

	err = e.ExpensesRepository.DeleteExpense(ctx, int64(expenseID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not delete expense: %v", err),
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
	subCategoryModel, err := e.ExpensesSubCategoryRepository.GetExpenseSubCategoryByName(ctx, subCategory)
	if err != nil {
		return dbModels.ExpenseSubCategoryTable{}, dbModels.CardTable{}, fmt.Errorf("could not get expense sub category by name: %v", err)
	}

	cardModel, err := e.CardRepository.GetCardByName(ctx, card)
	if err != nil {
		return dbModels.ExpenseSubCategoryTable{}, dbModels.CardTable{}, fmt.Errorf("could not get expense card by name: %v", err)
	}

	return subCategoryModel, cardModel, nil
}

func dateStringToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func timeToStringDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func expenseViewToExpenseGetResponse(expenseView dbModels.ExpenseView) models.ExpenseCreateRequest {
	return models.ExpenseCreateRequest{
		ID:          int(expenseView.ID),
		Value:       expenseView.Value,
		Date:        timeToStringDate(expenseView.Date),
		SubCategory: expenseView.SubCategory,
		Card:        expenseView.Card,
		Description: expenseView.Description,
	}
}

func expensesViewToExpensesGetResponse(expenseViewRecords []dbModels.ExpenseView) []models.ExpenseCreateRequest {
	var responseExpenses []models.ExpenseCreateRequest
	for _, exp := range expenseViewRecords {
		responseExpenses = append(responseExpenses, expenseViewToExpenseGetResponse(exp))
	}
	return responseExpenses
}
