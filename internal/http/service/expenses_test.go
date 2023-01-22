package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

	houseRentExpenseTable = rdsModels.ExpenseTable{
		ID:            1,
		Value:         10.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
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

	restaurantExpenseTable = rdsModels.ExpenseTable{
		ID:            2,
		Value:         20.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
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
	categoriesCache = cache.NewExpenseCategory(categories)

	subCategories = []rdsModels.ExpenseSubCategoryTable{
		{ID: 1, Name: "Rent", CategoryID: 1},
		{ID: 2, Name: "Restaurants", CategoryID: 2},
	}
	subCategoriesCache = cache.NewExpenseSubCategory(subCategories)
)

func TestExpensesService_CreateExpense(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

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
			expense: houseRentExpenseHTTPModel,
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
			}

			// WHEN
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

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

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
				Date:        firstFebruary2020String,
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "Test",
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
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
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

func TestExpensesService_GetExpenseByID(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
		{
			ID:            3,
			Value:         20.0,
			Date:          firstFebruary2020ZeroHoursUTCTime,
			SubCategoryID: 2,
			CardID:        3,
			Description:   "Unknown card",
		},
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expense    httpModels.Expense
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		{
			name: "Success",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusOK,
				expense:    houseRentExpenseHTTPModel,
			},
			params: map[string]string{"id": "1"},
		},
		{
			name: "ErrorUnknownCard",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not get expense by id:",
			},
			params: map[string]string{"id": "3"},
		},
		{
			name: "ErrorParameterIDNotInteger",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "id parameter must be an integer:",
			},
			params: map[string]string{"id": "abc"},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodGet,
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			expensesController.GetExpenseByID(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r httpModels.Expense
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expense, r)
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

func TestExpensesService_GetExpensesByCategory(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
		{
			ID:            3,
			Value:         250.0,
			Date:          firstFebruary2020ZeroHoursUTCTime,
			SubCategoryID: 1,
			CardID:        1,
			Description:   "Other House expense",
		},
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []httpModels.Expense
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		{
			name: "SuccessCategoryHouse",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusOK,
				expenses: []httpModels.Expense{
					houseRentExpenseHTTPModel,
					{
						ID:          3,
						Value:       250.0,
						Date:        firstFebruary2020String,
						SubCategory: "Rent",
						Card:        "CGD",
						Description: "Other House expense",
					},
				},
			},
			params: map[string]string{"category": "House"},
		},
		{
			name: "ErrorUnknownCategory",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not get expenses by category:",
			},
			params: map[string]string{"category": "Unknown"},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodGet,
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			expensesController.GetExpensesByCategory(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []httpModels.Expense
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
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

func TestExpensesService_GetExpensesBySubCategory(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
		{
			ID:            3,
			Value:         250.0,
			Date:          firstFebruary2020ZeroHoursUTCTime,
			SubCategoryID: 1,
			CardID:        1,
			Description:   "Other Rent expense",
		},
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []httpModels.Expense
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		{
			name: "SuccessSubCategoryRent",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusOK,
				expenses: []httpModels.Expense{
					houseRentExpenseHTTPModel,
					{
						ID:          3,
						Value:       250.0,
						Date:        firstFebruary2020String,
						SubCategory: "Rent",
						Card:        "CGD",
						Description: "Other Rent expense",
					},
				},
			},
			params: map[string]string{"sub_category": "Rent"},
		},
		{
			name: "ErrorUnknownSubCategory",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not get expenses by subcategory:",
			},
			params: map[string]string{"sub_category": "Unknown"},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodGet,
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			expensesController.GetExpensesBySubCategory(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []httpModels.Expense
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
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

func TestExpensesService_GetExpensesByCard(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
		{
			ID:            3,
			Value:         250.0,
			Date:          firstFebruary2020ZeroHoursUTCTime,
			SubCategoryID: 1,
			CardID:        1,
			Description:   "Other CGD card expense",
		},
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []httpModels.Expense
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		{
			name: "SuccessCardCGD",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusOK,
				expenses: []httpModels.Expense{
					houseRentExpenseHTTPModel,
					{
						ID:          3,
						Value:       250.0,
						Date:        firstFebruary2020String,
						SubCategory: "Rent",
						Card:        "CGD",
						Description: "Other CGD card expense",
					},
				},
			},
			params: map[string]string{"card": "CGD"},
		},
		{
			name: "ErrorUnknownCard",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not get expenses by card:",
			},
			params: map[string]string{"card": "Unknown"},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodGet,
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			expensesController.GetExpensesByCard(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []httpModels.Expense
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
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

func TestExpensesService_GetExpensesByDates(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []httpModels.Expense
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		{
			name: "Success",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusOK,
				expenses: []httpModels.Expense{
					houseRentExpenseHTTPModel,
					restaurantExpenseHTTPModel,
				},
			},
			params: map[string]string{"min_date": "2020-01-31", "max_date": "2020-02-02"},
		},
		{
			name: "ErrorWrongMinDateFormat",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse min_date (should use YYYY-MM-DD format):",
			},
			params: map[string]string{"min_date": "2020-Jan-31", "max_date": "2020-Feb-02"},
		},
		{
			name: "ErrorWrongMaxDateFormat",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse max_date (should use YYYY-MM-DD format):",
			},
			params: map[string]string{"min_date": "2020-01-31", "max_date": "2020-Feb-02"},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodGet,
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			expensesController.GetExpensesByDates(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []httpModels.Expense
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
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

func TestExpensesService_DeleteExpense(t *testing.T) {

	expenses := []rdsModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesController, err := NewExpensesService(&expensesCache, &subCategoriesCache, &cardsCache)
	if err != nil {
		t.Fatalf("error creating expenses service: %v\n", err)
	}

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		{
			name: "Success",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusNoContent,
			},
			params: map[string]string{"id": "1"},
		},
		{
			name: "ErrorUnexistingExpense",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not delete expense:",
			},
			params: map[string]string{"id": "5"},
		},
		{
			name: "ErrorParameterIDNotInteger",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not delete expense:",
			},
			params: map[string]string{"id": "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodDelete,
			}

			for k, v := range tt.params {
				ginCtx.Params = append(ginCtx.Params, gin.Param{Key: k, Value: v})
			}

			// WHEN
			expensesController.DeleteExpense(ginCtx)

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
