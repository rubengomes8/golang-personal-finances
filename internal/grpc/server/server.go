package server

import (
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

type FinancesServer struct {
	expenses.ExpensesServiceServer
}
