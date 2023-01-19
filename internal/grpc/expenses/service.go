package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	models "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
	"github.com/rubengomes8/golang-personal-finances/internal/pb/expenses"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
)

// ExpensesService implements ExpensesServiceServer methods
type ExpensesService struct {
	expenses.ExpensesServiceServer
	ExpensesRepository            repository.ExpenseRepo
	ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
	CardRepository                repository.CardRepo
}

// NewExpensesService creates a new ExpensesService
func NewExpensesService(
	expRepo repository.ExpenseRepo,
	expSubCatRepo repository.ExpenseSubCategoryRepo,
	cardRepo repository.CardRepo,
) (ExpensesService, error) {
	return ExpensesService{
		ExpensesRepository:            expRepo,
		ExpensesSubCategoryRepository: expSubCatRepo,
		CardRepository:                cardRepo,
	}, nil
}

// CreateExpense creates an expense on the database
func (s *ExpensesService) CreateExpense(
	ctx context.Context,
	req *expenses.ExpenseCreateRequest,
) (*expenses.ExpenseCreateResponse, error) {

	expSubCategory, card, err := s.getExpenseSubcategoryAndCardIDByNames(ctx, req.SubCategory, req.Card)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not get expense subcategory and/or card by name: %w", err)
	}

	expenseRecord := models.ExpenseTable{
		Value:         req.Value,
		Date:          unixToTime(req.Date),
		SubCategoryID: expSubCategory.ID,
		CardID:        card.ID,
		Description:   req.Description,
	}

	id, err := s.ExpensesRepository.InsertExpense(ctx, expenseRecord)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not insert expense: %w", err)
	}

	return &expenses.ExpenseCreateResponse{
		Id: id,
	}, nil
}

// UpdateExpense updates an expense on the database
func (s *ExpensesService) UpdateExpense(
	ctx context.Context,
	req *expenses.ExpenseUpdateRequest,
) (*expenses.ExpenseUpdateResponse, error) {

	expSubCategory, card, err := s.getExpenseSubcategoryAndCardIDByNames(ctx, req.SubCategory, req.Card)
	if err != nil {
		return &expenses.ExpenseUpdateResponse{}, fmt.Errorf("could not get expense subcategory and/or card by name: %w", err)
	}

	expenseRecord := models.ExpenseTable{
		ID:            req.Id,
		Value:         req.Value,
		Date:          unixToTime(req.Date),
		SubCategoryID: expSubCategory.ID,
		CardID:        card.ID,
		Description:   req.Description,
	}

	id, err := s.ExpensesRepository.UpdateExpense(ctx, expenseRecord)
	if err != nil {
		return &expenses.ExpenseUpdateResponse{}, fmt.Errorf("could not update expense: %w", err)
	}

	return &expenses.ExpenseUpdateResponse{
		Id: id,
	}, nil
}

// CreateExpenses creates a bulk of expenses on the database
func (s *ExpensesService) CreateExpenses(
	ctx context.Context,
	req *expenses.ExpensesCreateRequest,
) (*expenses.ExpensesCreateResponse, error) {
	log.Printf("TODO - CreateExpenses was invoked with %v\n", req)

	return &expenses.ExpensesCreateResponse{
		Ids: []*expenses.ExpenseCreateResponse{},
	}, nil
}

// GetExpensesByDate gets the expenses from the database that are in the provided dates interval
func (s *ExpensesService) GetExpensesByDate(
	ctx context.Context,
	req *expenses.ExpensesGetRequestByDate,
) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByDate was invoked with %v\n", req)

	expenseViewRecords, err := s.ExpensesRepository.GetExpensesByDates(
		ctx,
		unixToTime(req.MinDate),
		unixToTime(req.MaxDate),
	)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("could not get expenses by date: %w", err)
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	return &expenses.ExpensesGetResponse{
		Expenses: responseExpenses,
	}, nil
}

// GetExpensesByCategory gets the expenses from the database that match the category provided
func (s *ExpensesService) GetExpensesByCategory(
	ctx context.Context,
	req *expenses.ExpensesGetRequestByCategory,
) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByCategory was invoked with %v\n", req)

	expenseViewRecords, err := s.ExpensesRepository.GetExpensesByCategory(ctx, req.Category)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("could not get expenses by category: %w", err)
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	return &expenses.ExpensesGetResponse{
		Expenses: responseExpenses,
	}, nil
}

// GetExpensesBySubCategory gets the expenses from the database that match the subcategory provided
func (s *ExpensesService) GetExpensesBySubCategory(
	ctx context.Context,
	req *expenses.ExpensesGetRequestBySubCategory,
) (*expenses.ExpensesGetResponse, error) {

	log.Printf("GetExpensesBySubCategory was invoked with %v\n", req)

	expenseViewRecords, err := s.ExpensesRepository.GetExpensesBySubCategory(ctx, req.SubCategory)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("could not get expenses by subcategory: %v", err)
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	return &expenses.ExpensesGetResponse{
		Expenses: responseExpenses,
	}, nil
}

// GetExpensesByCard gets the expenses from the database that match the card provided
func (s *ExpensesService) GetExpensesByCard(
	ctx context.Context,
	req *expenses.ExpensesGetRequestByCard,
) (*expenses.ExpensesGetResponse, error) {

	log.Printf("GetExpensesByCard was invoked with %v\n", req)

	expenseViewRecords, err := s.ExpensesRepository.GetExpensesByCard(ctx, req.Card)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("could not get expenses by card: %v", err)
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	return &expenses.ExpensesGetResponse{
		Expenses: responseExpenses,
	}, nil
}

func unixToTime(unix int64) time.Time {
	return time.Unix(unix, 0).UTC()
}

func timeToUnix(date time.Time) int64 {
	return date.UTC().Unix()
}

func (s *ExpensesService) getExpenseSubcategoryAndCardIDByNames(
	ctx context.Context,
	subCategory, card string,
) (models.ExpenseSubCategoryTable, models.CardTable, error) {

	subCategoryModel, err := s.ExpensesSubCategoryRepository.GetExpenseSubCategoryByName(ctx, subCategory)
	if err != nil {
		return models.ExpenseSubCategoryTable{},
			models.CardTable{},
			fmt.Errorf("could not get expense sub category by name: %v", err)
	}

	cardModel, err := s.CardRepository.GetCardByName(ctx, card)
	if err != nil {
		return models.ExpenseSubCategoryTable{},
			models.CardTable{},
			fmt.Errorf("could not get expense card by name: %v", err)
	}

	return subCategoryModel, cardModel, nil
}

func expensesViewToExpensesGetResponse(
	expenseViewRecords []models.ExpenseView,
) []*expenses.ExpenseGetResponse {

	var responseExpenses []*expenses.ExpenseGetResponse

	var responseExpense expenses.ExpenseGetResponse

	for _, exp := range expenseViewRecords {

		unixDate := timeToUnix(exp.Date)

		responseExpense = expenses.ExpenseGetResponse{
			Id:          exp.ID,
			Value:       exp.Value,
			Date:        unixDate,
			Category:    exp.Category,
			SubCategory: exp.SubCategory,
			Card:        exp.Card,
			Description: exp.Description,
		}

		responseExpenses = append(responseExpenses, &responseExpense)
	}

	return responseExpenses
}
