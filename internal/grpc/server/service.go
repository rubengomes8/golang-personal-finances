package server

import (
	"context"
	"fmt"
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/models"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

func (s *FinancesServer) CreateExpense(ctx context.Context, req *expenses.ExpenseCreateRequest) (*expenses.ExpenseCreateResponse, error) {

	expense := models.Expense{
		Value:       req.Value,
		Date:        req.Date,
		Category:    req.Category,
		SubCategory: req.SubCategory,
		Card:        req.Card,
		Description: req.Description,
	}

	id, err := s.ExpensesRepository.InsertExpense(ctx, expense)
	if err != nil {
		return &expenses.ExpenseCreateResponse{}, fmt.Errorf("could not insert expense: %v", err)
	}

	return &expenses.ExpenseCreateResponse{
		Id: id,
	}, nil
}

func (s *FinancesServer) CreateExpenses(ctx context.Context, req *expenses.ExpensesCreateRequest) (*expenses.ExpensesCreateResponse, error) {
	log.Printf("CreateExpenses was invoked with %v\n", req)
	return &expenses.ExpensesCreateResponse{
		Ids: []*expenses.ExpenseCreateResponse{
			{Id: 1},
			{Id: 2},
		},
	}, nil
}

func (s *FinancesServer) GetExpensesByDate(ctx context.Context, req *expenses.ExpensesGetRequestByDate) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByDate was invoked with %v\n", req)

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func (s *FinancesServer) GetExpensesByCategory(ctx context.Context, req *expenses.ExpensesGetRequestByCategory) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpenseByCategory was invoked with %v\n", req)

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func (s *FinancesServer) GetExpensesBySubCategory(ctx context.Context, req *expenses.ExpensesGetRequestBySubCategory) (*expenses.ExpensesGetResponse, error) {
	log.Printf("GetExpensesBySubCategory was invoked with %v\n", req)

	return &expenses.ExpensesGetResponse{
		Expenses: getTestExpenses(),
	}, nil
}

func (s *FinancesServer) GetExpensesByCard(ctx context.Context, req *expenses.ExpensesGetRequestByCard) (*expenses.ExpensesGetResponse, error) {
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
