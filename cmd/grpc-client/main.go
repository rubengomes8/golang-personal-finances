package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	expensesclient "github.com/rubengomes8/golang-personal-finances/internal/grpc/client/expenses"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

const ADDR = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	cExpenses := expenses.NewExpensesServiceClient(conn)
	// cCards := cards.NewCardServiceClient(conn)

	// cardsclient.CreateCard(cCards)
	// time.Sleep(500 * time.Millisecond)

	expensesclient.CreateExpense(cExpenses)
	time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByCard(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByCategory(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesBySubCategory(c)
	// time.Sleep(500 * time.Millisecond)

	// client.GetExpensesByDate(c)
	// time.Sleep(500 * time.Millisecond)
}
