package grpc

import (
	"context"
	"reflect"
	"testing"
	"time"

	rdsModels "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
	grpc "github.com/rubengomes8/golang-personal-finances/internal/pb/expenses"
	"github.com/stretchr/testify/assert"

	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/cache"
)

var (
	firstFebruary2020ZeroHoursUTCTime = time.Date(2020, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	firstFebruary2020Unix             = int64(1580515200)

	houseRentExpenseView = rdsModels.ExpenseView{
		ID:            1,
		Value:         10.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
		Category:      "House",
		SubCategory:   "Rent",
		Card:          "CGD",
		CategoryID:    1,
		SubCategoryID: 1,
		CardID:        1,
		Description:   "Test",
	}

	houseRentGRPCExpenseCreateRequest = grpc.ExpenseCreateRequest{
		Value:       10.0,
		Date:        firstFebruary2020Unix,
		Category:    "House",
		SubCategory: "Rent",
		Card:        "CGD",
		Description: "Test",
	}

	houseRentGRPCExpenseCreateResponse = grpc.ExpenseCreateResponse{
		Id: 1,
	}

	houseRentExpenseTable = rdsModels.ExpenseTable{
		ID:            1,
		Value:         10.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
		SubCategoryID: 1,
		CardID:        1,
		Description:   "Test",
	}

	restaurantExpenseView = rdsModels.ExpenseView{
		ID:            2,
		Value:         20.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
		Category:      "Leisure",
		SubCategory:   "Restaurants",
		Card:          "Food allowance",
		CategoryID:    2,
		SubCategoryID: 2,
		CardID:        2,
		Description:   "Test",
	}

	restaurantExpenseTable = rdsModels.ExpenseTable{
		ID:            2,
		Value:         20.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
		SubCategoryID: 2,
		CardID:        2,
		Description:   "Test",
	}
)

// REPO
var (
	cards = []rdsModels.CardTable{
		{ID: 1, Name: "CGD"},
		{ID: 2, Name: "Food allowance"},
	}
	cardsCache = cache.NewCard(cards)

	categories = []rdsModels.ExpenseCategoryTable{
		{ID: 1, Name: "House"},
		{ID: 2, Name: "Leisure"},
	}
	categoriesCache = cache.NewExpenseCategory(categories)

	subCategories = []rdsModels.ExpenseSubCategoryTable{
		{ID: 1, Name: "Rent", CategoryID: 1},
		{ID: 2, Name: "Restaurants", CategoryID: 2},
	}
	subCategoriesCache = cache.NewExpenseSubCategory(subCategories)
)

func Test_timeToUnix(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Success",
			args: args{
				date: firstFebruary2020ZeroHoursUTCTime,
			},
			want: firstFebruary2020Unix,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeToUnix(tt.args.date); got != tt.want {
				t.Errorf("timeToUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unixToTime(t *testing.T) {
	type args struct {
		unix int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Success",
			args: args{
				unix: firstFebruary2020Unix,
			},
			want: firstFebruary2020ZeroHoursUTCTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unixToTime(tt.args.unix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpensesService_CreateExpense(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type args struct {
		ctx context.Context
		req *grpc.ExpenseCreateRequest
	}

	type want struct {
		response *grpc.ExpenseCreateResponse
		errorMsg string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
				CardRepository:                &cardsCache,
			},
			args: args{
				ctx: context.Background(),
				req: &houseRentGRPCExpenseCreateRequest,
			},
			want: want{
				response: &houseRentGRPCExpenseCreateResponse,
			},
			wantErr: false,
		},
		{
			name: "ErrorUnknownCard",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
				CardRepository:                &cardsCache,
			},
			args: args{
				ctx: context.Background(),
				req: &grpc.ExpenseCreateRequest{
					Value:       10.0,
					Date:        firstFebruary2020Unix,
					Category:    "House",
					SubCategory: "Rent",
					Card:        "Unknown",
					Description: "Test",
				},
			},
			want: want{
				errorMsg: "could not get expense subcategory and/or card by name:",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &ExpensesService{
				ExpensesRepository:            tt.fields.ExpensesRepository,
				ExpensesSubCategoryRepository: tt.fields.ExpensesSubCategoryRepository,
				CardRepository:                tt.fields.CardRepository,
			}

			got, err := s.CreateExpense(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpensesService.CreateExpense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			switch {
			case !tt.wantErr:
				if !reflect.DeepEqual(got, tt.want.response) {
					t.Errorf("ExpensesService.CreateExpense() = %v, want %v", got, tt.want.response)
				}
			case tt.wantErr:
				assert.Contains(t, err.Error(), tt.want.errorMsg)
			}

		})
	}
}
