package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protoExpenses "github.com/rubengomes8/golang-personal-finances/proto/expenses"
)

const ADDR = "0.0.0.0:50051"

func createExpense(serviceClient protoExpenses.ExpensesServiceClient, expense *protoExpenses.ExpenseCreateRequest) {

	log.Println("createExpense was invoked")

	res, err := serviceClient.CreateExpense(context.Background(), expense)
	if err != nil {
		log.Fatalf("client could not request for create expense: %v\n", err)
	}

	log.Printf("Requested create expense with ID: %d\n", res.Id)
}

func main() {
	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	expense := protoExpenses.ExpenseCreateRequest{
		Value:       3,
		Date:        10,
		Category:    "House",
		SubCategory: "Rent",
		Card:        "CGD",
		Description: "Estefânia 92",
	}

	client := protoExpenses.NewExpensesServiceClient(conn)

	createExpense(client, &expense)
}
