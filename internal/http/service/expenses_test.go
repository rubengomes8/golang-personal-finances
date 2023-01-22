package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	httpModels "github.com/rubengomes8/golang-personal-finances/internal/models/http"
	rdsModels "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/cache"

	"github.com/stretchr/testify/assert"
)

var (
	firstFebruary2020ZeroHoursUTCTime = time.Date(2020, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	firstFebruary2020String           = "2020-02-01"

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

	houseRentExpenseHTTPModel = httpModels.Expense{
		ID:          1,
		Value:       10.0,
		Date:        firstFebruary2020String,
		SubCategory: "Rent",
		Card:        "CGD",
		Description: "Test",
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

	restaurantExpenseHTTPModel = httpModels.Expense{
		ID:          2,
		Value:       20.0,
		Date:        firstFebruary2020String,
		SubCategory: "Restaurants",
		Card:        "Food allowance",
		Description: "Test",
	}
)

func Test_timeToStringDate(t *testing.T) {
	type args struct {
		t time.Time
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Successful",
			args: args{
				t: firstFebruary2020ZeroHoursUTCTime,
			},
			want: firstFebruary2020String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeToStringDate(tt.args.t); got != tt.want {
				t.Errorf("timeToStringDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dateStringToTime(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				date: firstFebruary2020String,
			},
			want:    firstFebruary2020ZeroHoursUTCTime,
			wantErr: false,
		},
		{
			name: "ErrorWrongDateLayout",
			args: args{
				date: "2020-Feb-01",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dateStringToTime(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateStringToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateStringToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expenseViewToExpenseGetResponse(t *testing.T) {
	type args struct {
		expenseView rdsModels.ExpenseView
	}
	tests := []struct {
		name string
		args args
		want httpModels.Expense
	}{
		{
			name: "Success",
			args: args{
				expenseView: houseRentExpenseView,
			},
			want: houseRentExpenseHTTPModel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expenseViewToExpenseGetResponse(tt.args.expenseView); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expenseViewToExpenseGetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expensesViewToExpensesGetResponse(t *testing.T) {
	type args struct {
		expenseViewRecords []rdsModels.ExpenseView
	}
	tests := []struct {
		name string
		args args
		want []httpModels.Expense
	}{
		{
			name: "Success",
			args: args{
				expenseViewRecords: []rdsModels.ExpenseView{
					houseRentExpenseView, restaurantExpenseView,
				},
			},
			want: []httpModels.Expense{
				houseRentExpenseHTTPModel, restaurantExpenseHTTPModel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expensesViewToExpensesGetResponse(tt.args.expenseViewRecords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expensesViewToExpensesGetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	categoryCache = cache.NewExpenseCategory(categories)

	subCategories = []rdsModels.ExpenseSubCategoryTable{
		{ID: 1, Name: "Rent", CategoryID: 1},
		{ID: 2, Name: "Restaurants", CategoryID: 2},
	}
	subCategoriesCache = cache.NewExpenseSubCategory(subCategories)

	expenses = []rdsModels.ExpenseTable{
		{
			ID:            1,
			Value:         200.0,
			Date:          firstFebruary2020ZeroHoursUTCTime,
			SubCategoryID: 1,
			CardID:        1,
			Description:   "House rent",
		},
		{
			ID:            2,
			Value:         50.0,
			Date:          firstFebruary2020ZeroHoursUTCTime,
			SubCategoryID: 2,
			CardID:        2,
			Description:   "Dinner on Ramiro",
		},
	}
	expensesCache = cache.NewExpense(expenses)
)

func TestExpensesService_CreateExpense(t *testing.T) {

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenseID  int
		errorMsg   string
	}

	tests := []struct {
		name    string
		expense httpModels.Expense
		fields  fields
		want    want
	}{
		{
			name: "Success",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				Value:       200.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusCreated,
				expenseID:  1,
			},
		},
		{
			name: "ErrorUnknownCard",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				Value:       200.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "Unknown",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "unknown subcategory or card:",
			},
		},
		{
			name: "ErrorUnexpectedDateFormat",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				Value:       200.0,
				Date:        "01-Feb-2020",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse date (should use YYYY-MM-DD format):",
			},
		},
	}

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			data, err := json.Marshal(tt.expense)
			if err != nil {
				t.Fatalf("error marshaling expense: %v\n", err)
			}

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodPost,
				Body:   io.NopCloser(bytes.NewBuffer(data)),
				URL: &url.URL{
					Scheme: "http",
					Host:   "localhost:8080",
					Path:   "/v1/expense",
				},
			}

			// WHEN
			gin.SetMode(gin.TestMode)
			expensesController.CreateExpense(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusCreated:
				var r httpModels.ExpenseCreateResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenseID, r.ID)
			case http.StatusBadRequest:
				var r httpModels.ErrorResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Contains(t, r.ErrorMsg, tt.want.errorMsg)
			}
		})
	}
}

func TestExpensesService_UpdateExpense(t *testing.T) {

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenseID  int
		errorMsg   string
	}

	tests := []struct {
		name    string
		expense httpModels.Expense
		fields  fields
		want    want
		params  map[string]string
	}{
		{
			name: "Success",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				ID:          1,
				Value:       250.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusNoContent,
				expenseID:  1,
			},
			params: map[string]string{"id": "1"},
		},
		{
			name: "ErrorMissingParameterID",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				ID:          1,
				Value:       250.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "id parameter must be an integer:",
			},
		},
		{
			name: "ErrorUnknownCard",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				ID:          1,
				Value:       200.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "Unknown",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "unknown subcategory or card:",
			},
		},
		{
			name: "ErrorUnexpectedDateFormat",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: httpModels.Expense{
				ID:          1,
				Value:       200.0,
				Date:        "01-Feb-2020",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse date (should use YYYY-MM-DD format):",
			},
		},
	}

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			data, err := json.Marshal(tt.expense)
			if err != nil {
				t.Fatalf("error marshaling expense: %v\n", err)
			}

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodPut,
				Body:   io.NopCloser(bytes.NewBuffer(data)),
				URL: &url.URL{
					Scheme: "http",
					Host:   "localhost:8080",
					Path:   fmt.Sprintf("/v1/expense/%d", tt.expense.ID),
				},
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			gin.SetMode(gin.TestMode)
			expensesController.UpdateExpense(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusNoContent:
			case http.StatusBadRequest:
				var r httpModels.ErrorResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Contains(t, r.ErrorMsg, tt.want.errorMsg)
			}
		})
	}
}
