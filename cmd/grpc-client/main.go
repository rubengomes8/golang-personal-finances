package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rubengomes8/golang-personal-finances/internal/grpc/client"
	"github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

const ADDR = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := expenses.NewExpensesServiceClient(conn)

	client.CreateExpense(c)
	time.Sleep(2 * time.Second)

	client.GetExpensesByCard(c)
	time.Sleep(2 * time.Second)

	client.GetExpensesByCategory(c)
	time.Sleep(2 * time.Second)

	client.GetExpensesBySubCategory(c)
	time.Sleep(2 * time.Second)

	client.GetExpensesByDate(c)
	time.Sleep(2 * time.Second)
}
