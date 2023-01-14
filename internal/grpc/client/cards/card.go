package client

import (
	"context"
	"log"

	"github.com/rubengomes8/golang-personal-finances/proto/cards"
)

func CreateCard(serviceClient cards.CardServiceClient) {

	log.Println("CreateCard was invoked")

	card := cards.CardCreateRequest{
		Name: "Santander",
	}

	res, err := serviceClient.CreateCard(context.Background(), &card)
	if err != nil {
		log.Fatalf("client could not request for create card: %v\n", err)
	}

	log.Printf("Card created with ID: %d\n", res.Id)
}
