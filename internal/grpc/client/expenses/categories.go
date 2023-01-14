package client

import (
	"context"
	"log"

	categories "github.com/rubengomes8/golang-personal-finances/proto/expense_categories"
)

func CreateCategory(serviceClient categories.ExpenseCategoryServiceClient) {

	log.Println("CreateCategory was invoked")

	category := categories.ExpenseCategoryCreateRequest{
		Name: "Test",
	}

	res, err := serviceClient.CreateExpenseCategory(context.Background(), &category)
	if err != nil {
		log.Fatalf("client could not request for create category: %v\n", err)
	}

	log.Printf("Expense category created with ID: %d\n", res.Id)
}
