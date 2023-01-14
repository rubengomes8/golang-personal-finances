package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	client "github.com/rubengomes8/golang-personal-finances/internal/grpc/client/expenses"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

const ADDR = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// cardsclient.CreateCard(cCards)
	// time.Sleep(500 * time.Millisecond)

	// cCards := cards.NewCardServiceClient(conn)

	cExpenses := expenses.NewExpensesServiceClient(conn)
	// nowUnixDate := time.Now().UTC().Unix()
	// expensesclient.CreateExpense(
	// 	cExpenses,
	// 	expenses.ExpenseCreateRequest{
	// 		Value:       10,
	// 		Date:        nowUnixDate,
	// 		Category:    "House",
	// 		SubCategory: "Rent",
	// 		Card:        "CGD",
	// 		Description: "TEST",
	// 	},
	// )
	// time.Sleep(500 * time.Millisecond)

	// expensesclient.CreateExpense(
	// 	cExpenses,
	// 	expenses.ExpenseCreateRequest{
	// 		Value:       20,
	// 		Date:        nowUnixDate,
	// 		Category:    "Laser",
	// 		SubCategory: "Restaurants",
	// 		Card:        "Food allowance",
	// 		Description: "TEST",
	// 	},
	// )
	// time.Sleep(500 * time.Millisecond)

	// expensesclient.CreateExpense(
	// 	cExpenses,
	// 	expenses.ExpenseCreateRequest{
	// 		Value:       30,
	// 		Date:        nowUnixDate,
	// 		Category:    "Laser",
	// 		SubCategory: "Rent",
	// 		Card:        "Food allowance",
	// 		Description: "INVALID TEST",
	// 	},
	// )
	// time.Sleep(500 * time.Millisecond)

	card := expenses.ExpensesGetRequestByCard{
		Card: "CGD",
	}
	client.GetExpensesByCard(cExpenses, &card)
	time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByCategory(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesBySubCategory(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByDate(c)
	// time.Sleep(500 * time.Millisecond)
}
