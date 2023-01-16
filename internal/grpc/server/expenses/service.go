package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

func (s *ExpensesService) CreateExpense(ctx context.Context, req *expenses.ExpenseCreateRequest) (*expenses.ExpenseCreateResponse, error) {

	expSubCategory, card, err := s.getExpenseSubcategoryAndCardIdByNames(ctx, req.SubCategory, req.Card)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not get expense subcategory and/or card by name: %v", err)
	}

	expenseRecord := models.ExpenseTable{
		Value:         req.Value,
		Date:          unixToTime(req.Date),
		SubCategoryId: expSubCategory.Id,
		CardId:        card.Id,
		Description:   req.Description,
	}

	id, err := s.ExpensesRepository.InsertExpense(ctx, expenseRecord)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not insert expense: %v", err)
	}

	return &expenses.ExpenseCreateResponse{
		Id: id,
	}, nil
}

func (s *ExpensesService) UpdateExpense(ctx context.Context, req *expenses.ExpenseUpdateRequest) (*expenses.ExpenseUpdateResponse, error) {

	expSubCategory, card, err := s.getExpenseSubcategoryAndCardIdByNames(ctx, req.SubCategory, req.Card)
	if err != nil {
		return &expenses.ExpenseUpdateResponse{}, fmt.Errorf("could not get expense subcategory and/or card by name: %v", err)
	}

	expenseRecord := models.ExpenseTable{
		Id:            req.Id,
		Value:         req.Value,
		Date:          unixToTime(req.Date),
		SubCategoryId: expSubCategory.Id,
		CardId:        card.Id,
		Description:   req.Description,
	}

	id, err := s.ExpensesRepository.UpdateExpense(ctx, expenseRecord)
	if err != nil {
		return &expenses.ExpenseUpdateResponse{}, fmt.Errorf("could not update expense: %v", err)
	}

	return &expenses.ExpenseUpdateResponse{
		Id: id,
	}, nil
}

func (s *ExpensesService) CreateExpenses(ctx context.Context, req *expenses.ExpensesCreateRequest) (*expenses.ExpensesCreateResponse, error) {
	log.Printf("TODO - CreateExpenses was invoked with %v\n", req)
	return &expenses.ExpensesCreateResponse{
		Ids: []*expenses.ExpenseCreateResponse{
			{Id: 1},
			{Id: 2},
		},
	}, nil
}

func (s *ExpensesService) GetExpensesByDate(ctx context.Context, req *expenses.ExpensesGetRequestByDate) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByDate was invoked with %v\n", req)

	expenseViewRecords, err := s.ExpensesRepository.GetExpensesByDates(ctx, unixToTime(req.MinDate), unixToTime(req.MaxDate))
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("could not get expenses by date: %v", err)
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	return &expenses.ExpensesGetResponse{
		Expenses: responseExpenses,
	}, nil
}

func (s *ExpensesService) GetExpensesByCategory(ctx context.Context, req *expenses.ExpensesGetRequestByCategory) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByCategory was invoked with %v\n", req)

	expenseViewRecords, err := s.ExpensesRepository.GetExpensesByCategory(ctx, req.Category)
	if err != nil {
		return &expenses.ExpensesGetResponse{}, fmt.Errorf("could not get expenses by category: %v", err)
	}

	responseExpenses := expensesViewToExpensesGetResponse(expenseViewRecords)

	return &expenses.ExpensesGetResponse{
		Expenses: responseExpenses,
	}, nil
}

func (s *ExpensesService) GetExpensesBySubCategory(ctx context.Context, req *expenses.ExpensesGetRequestBySubCategory) (*expenses.ExpensesGetResponse, error) {
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

func (s *ExpensesService) GetExpensesByCard(ctx context.Context, req *expenses.ExpensesGetRequestByCard) (*expenses.ExpensesGetResponse, error) {
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

func (s *ExpensesService) getExpenseSubcategoryAndCardIdByNames(
	ctx context.Context,
	subCategory, card string,
) (models.ExpenseSubCategoryTable, models.CardTable, error) {
	subCategoryModel, err := s.ExpensesSubCategoryRepository.GetExpenseSubCategoryByName(ctx, subCategory)
	if err != nil {
		return models.ExpenseSubCategoryTable{}, models.CardTable{}, fmt.Errorf("could not get expense sub category by name: %v", err)
	}

	cardModel, err := s.CardRepository.GetCardByName(ctx, card)
	if err != nil {
		return models.ExpenseSubCategoryTable{}, models.CardTable{}, fmt.Errorf("could not get expense card by name: %v", err)
	}

	return subCategoryModel, cardModel, nil
}

func expensesViewToExpensesGetResponse(expenseViewRecords []models.ExpenseView) []*expenses.ExpenseGetResponse {

	var responseExpenses []*expenses.ExpenseGetResponse
	var responseExpense expenses.ExpenseGetResponse
	for _, exp := range expenseViewRecords {

		responseExpense = expenses.ExpenseGetResponse{
			Id:          exp.Id,
			Value:       exp.Value,
			Date:        exp.Date,
			Category:    exp.Category,
			SubCategory: exp.SubCategory,
			Card:        exp.Card,
			Description: exp.Description,
		}

		responseExpenses = append(responseExpenses, &responseExpense)
	}

	return responseExpenses
}
