package server

import (
	"context"
	"log"

	expenses "github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

func (s *FinancesServer) CreateExpense(ctx context.Context, req *expenses.ExpenseCreateRequest) (*expenses.ExpenseCreateResponse, error) {
	log.Printf("CreateExpense was invoked with %v\n", req)
	return &expenses.ExpenseCreateResponse{
		Id: 1,
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
