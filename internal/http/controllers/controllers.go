package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	httpModels "github.com/rubengomes8/golang-personal-finances/internal/models/http"
	rdsModels "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
)

type ExpensesController struct {
	ExpensesRepository            repository.ExpenseRepo
	ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
	CardRepository                repository.CardRepo
	Database                      *sql.DB
}

func NewExpensesController(
	expRepo repository.ExpenseRepo,
	expSubCatRepo repository.ExpenseSubCategoryRepo,
	cardRepo repository.CardRepo,
) (ExpensesController, error) {
	return ExpensesController{
		ExpensesRepository:            expRepo,
		ExpensesSubCategoryRepository: expSubCatRepo,
		CardRepository:                cardRepo,
	}, nil
}

func (e *ExpensesController) CreateExpense(ctx *gin.Context) {

	var expense httpModels.Expense
	err := json.NewDecoder(ctx.Request.Body).Decode(&expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not decode expense: %v", err)})
		return
	}

	expSubCategory, card, err := e.getExpenseSubcategoryAndCardIdByNames(ctx, expense.SubCategory, expense.Card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unknown subcategory or card: %v", err)})
		return
	}

	date, err := dateStringToTime(expense.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not parse date (should use YYYY-MM-DD format): %v", err)})
		return
	}
	expenseRecord := rdsModels.ExpenseTable{
		Value:         expense.Value,
		Date:          date,
		SubCategoryId: expSubCategory.Id,
		CardId:        card.Id,
		Description:   expense.Description,
	}

	id, err := e.ExpensesRepository.InsertExpense(ctx, expenseRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not insert expense: %v", err)})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateExpense(ctx *gin.Context) {
	// TODO
}

func GetExpense(ctx *gin.Context) {
	// TODO
}

func DeleteExpense(ctx *gin.Context) {
	// TODO
}

func (e *ExpensesController) getExpenseSubcategoryAndCardIdByNames(
	ctx context.Context,
	subCategory, card string,
) (rdsModels.ExpenseSubCategoryTable, rdsModels.CardTable, error) {
	subCategoryModel, err := e.ExpensesSubCategoryRepository.GetExpenseSubCategoryByName(ctx, subCategory)
	if err != nil {
		return rdsModels.ExpenseSubCategoryTable{}, rdsModels.CardTable{}, fmt.Errorf("could not get expense sub category by name: %v", err)
	}

	cardModel, err := e.CardRepository.GetCardByName(ctx, card)
	if err != nil {
		return rdsModels.ExpenseSubCategoryTable{}, rdsModels.CardTable{}, fmt.Errorf("could not get expense card by name: %v", err)
	}

	return subCategoryModel, cardModel, nil
}

func dateStringToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
