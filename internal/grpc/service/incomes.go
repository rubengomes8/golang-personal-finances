package service

import (
	"context"
	"fmt"
	"log"

	"github.com/rubengomes8/golang-personal-finances/internal/pb/incomes"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Incomes implements IncomesServiceServer methods
type Incomes struct {
	incomes.ServiceServer
	Repository         repository.IncomeRepo
	CategoryRepository repository.IncomeCategoryRepo
	CardRepository     repository.CardRepo
}

// NewIncomes creates a new Incomes service
func NewIncomes(
	repo repository.IncomeRepo,
	catRepo repository.IncomeCategoryRepo,
	cardRepo repository.CardRepo,
) (Incomes, error) {
	return Incomes{
		Repository:         repo,
		CategoryRepository: catRepo,
		CardRepository:     cardRepo,
	}, nil
}

// CreateIncome creates an income on the database
func (i *Incomes) CreateIncome(
	ctx context.Context,
	req *incomes.CreateRequest,
) (*incomes.CreateResponse, error) {

	card, err := i.CardRepository.GetCardByName(ctx, req.Card)
	if err != nil {
		log.Printf("grpc - could not get card by name: %v", err)
		return &incomes.CreateResponse{}, fmt.Errorf("could not get income card by name")

	}

	category, err := i.CategoryRepository.GetIncomeCategoryByName(ctx, req.Category)
	if err != nil {
		log.Printf("grpc - could not get income category by name: %v", err)
		return &incomes.CreateResponse{}, fmt.Errorf("could not get income category by name")
	}

	incomeRecord := models.IncomeTable{
		Value:       req.Value,
		Date:        req.Date.AsTime(),
		CategoryID:  category.ID,
		CardID:      card.ID,
		Description: req.Description,
	}

	id, err := i.Repository.InsertIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("grpc - could not insert income: %v", err)
		return &incomes.CreateResponse{}, fmt.Errorf("could not insert income")
	}

	return &incomes.CreateResponse{
		Id: id,
	}, nil
}

// UpdateIncome updates an income on the database
func (i *Incomes) UpdateIncome(
	ctx context.Context,
	req *incomes.UpdateRequest,
) (*incomes.UpdateResponse, error) {

	card, err := i.CardRepository.GetCardByName(ctx, req.Card)
	if err != nil {
		log.Printf("grpc - could not get card by name: %v", err)
		return &incomes.UpdateResponse{}, fmt.Errorf("could not get income card by name: %w", err)

	}

	category, err := i.CategoryRepository.GetIncomeCategoryByName(ctx, req.Category)
	if err != nil {
		log.Printf("grpc - could not get income category by name: %v", err)
		return &incomes.UpdateResponse{}, fmt.Errorf("could not get income category by name: %w", err)
	}

	incomeRecord := models.IncomeTable{
		ID:          req.Id,
		Value:       req.Value,
		Date:        req.Date.AsTime(),
		CardID:      card.ID,
		CategoryID:  category.ID,
		Description: req.Description,
	}

	id, err := i.Repository.UpdateIncome(ctx, incomeRecord)
	if err != nil {
		log.Printf("grpc - could not get update income: %v", err)
		return &incomes.UpdateResponse{}, fmt.Errorf("could not update income: %w", err)
	}

	return &incomes.UpdateResponse{
		Id: id,
	}, nil
}

// GetIncomesByDate gets the incomes from the database that are in the provided dates interval
func (i *Incomes) GetIncomesByDate(
	ctx context.Context,
	req *incomes.GetRequestByDate,
) (*incomes.GetSeveralResponse, error) {
	log.Printf("GetIncomesByDate was invoked with %v\n", req)

	incomeViewRecords, err := i.Repository.GetIncomesByDates(
		ctx,
		req.MinDate.AsTime(),
		req.MaxDate.AsTime(),
	)
	if err != nil {
		log.Printf("grpc - could not get incomes by dates %v", err)
		return &incomes.GetSeveralResponse{}, fmt.Errorf("could not get incomes by dates")
	}

	responseIncomes := incomeViewsToIncomesGetResponse(incomeViewRecords)

	return &incomes.GetSeveralResponse{
		Incomes: responseIncomes,
	}, nil
}

// GetIncomesByCategory gets the incomes from the database that match the category provided
func (i *Incomes) GetIncomesByCategory(
	ctx context.Context,
	req *incomes.GetRequestByCategory,
) (*incomes.GetSeveralResponse, error) {
	log.Printf("GetIncomesByCategory was invoked with %v\n", req)

	incomeViewRecords, err := i.Repository.GetIncomesByCategory(
		ctx,
		req.Category,
	)
	if err != nil {
		log.Printf("grpc - could not get incomes by category %v", err)
		return &incomes.GetSeveralResponse{}, fmt.Errorf("could not get incomes by category")
	}

	responseIncomes := incomeViewsToIncomesGetResponse(incomeViewRecords)

	return &incomes.GetSeveralResponse{
		Incomes: responseIncomes,
	}, nil
}

// GetIncomesByCard gets the incomes from the database that match the card provided
func (i *Incomes) GetIncomesByCard(
	ctx context.Context,
	req *incomes.GetRequestByCard,
) (*incomes.GetSeveralResponse, error) {
	log.Printf("GetIncomesByCard was invoked with %v\n", req)

	incomeViewRecords, err := i.Repository.GetIncomesByCard(
		ctx,
		req.Card,
	)
	if err != nil {
		log.Printf("grpc - could not get incomes by card %v", err)
		return &incomes.GetSeveralResponse{}, fmt.Errorf("could not get incomes by card")
	}

	responseIncomes := incomeViewsToIncomesGetResponse(incomeViewRecords)

	return &incomes.GetSeveralResponse{
		Incomes: responseIncomes,
	}, nil
}

func incomeViewsToIncomesGetResponse(
	incomeViewRecords []models.IncomeView,
) []*incomes.GetResponse {

	var responseIncomes []*incomes.GetResponse

	for _, inc := range incomeViewRecords {

		responseIncome := incomes.GetResponse{
			Id:          inc.ID,
			Value:       inc.Value,
			Date:        timestamppb.New(inc.Date),
			Category:    inc.Category,
			Card:        inc.Card,
			Description: inc.Description,
		}

		responseIncomes = append(responseIncomes, &responseIncome)
	}

	return responseIncomes
}
