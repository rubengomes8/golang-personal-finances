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

	expSubCategory, err := s.ExpensesSubCategoryRepository.GetExpenseSubCategoryByName(ctx, req.SubCategory)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not get expense sub category by name: %v", err)
	}

	card, err := s.CardRepository.GetCardByName(ctx, req.Card)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not get expense card by name: %v", err)
	}

	expenseTable := models.ExpenseTable{
		Value:         req.Value,
		Date:          unixToTime(req.Date),
		SubCategoryId: expSubCategory.Id,
		CardId:        card.Id,
		Description:   req.Description,
	}

	id, err := s.ExpensesRepository.InsertExpense(ctx, expenseTable)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not insert expense: %v", err)
	}

	return &expenses.ExpenseCreateResponse{
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

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func (s *ExpensesService) GetExpensesByCategory(ctx context.Context, req *expenses.ExpensesGetRequestByCategory) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByCategory was invoked with %v\n", req)

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func (s *ExpensesService) GetExpensesBySubCategory(ctx context.Context, req *expenses.ExpensesGetRequestBySubCategory) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpensesBySubCategory was invoked with %v\n", req)

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func (s *ExpensesService) GetExpensesByCard(ctx context.Context, req *expenses.ExpensesGetRequestByCard) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpensesByCard was invoked with %v\n", req)

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func getTestExpenses() []*expenses.ExpenseGetResponse {
	return []*expenses.ExpenseGetResponse{
		{Id: 1, Value: 20.5, Date: 1, Category: "House", SubCategory: "Rent", Card: "CGD", Description: "Test"},
		{Id: 2, Value: 10.5, Date: 2, Category: "House", SubCategory: "Rent", Card: "CGD", Description: "Test"},
	}
}

func unixToTime(unix int64) time.Time {
	return time.Unix(unix, 0).UTC()
}

func timeToUnix(date time.Time) int64 {
	return date.UTC().Unix()
}