package handlers

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
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/cache"
	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"
	"github.com/stretchr/testify/assert"
)

var (
	firstFebruary2020ZeroHoursUTCTime = time.Date(2020, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	firstFebruary2020String           = "2020-02-01"

	houseRentExpenseView = dbModels.ExpenseView{
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

	houseRentExpenseTable = dbModels.ExpenseTable{
		ID:            1,
		Value:         10.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
		SubCategoryID: 1,
		CardID:        1,
		Description:   "Test",
	}

	houseRentExpenseHTTPModel = models.ExpenseCreateRequest{
		ID:          1,
		Value:       10.0,
		Date:        firstFebruary2020String,
		SubCategory: "Rent",
		Card:        "CGD",
		Description: "Test",
	}

	restaurantExpenseView = dbModels.ExpenseView{
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

	restaurantExpenseTable = dbModels.ExpenseTable{
		ID:            2,
		Value:         20.0,
		Date:          firstFebruary2020ZeroHoursUTCTime,
		SubCategoryID: 2,
		CardID:        2,
		Description:   "Test",
	}

	restaurantExpenseHTTPModel = models.ExpenseCreateRequest{
		ID:          2,
		Value:       20.0,
		Date:        firstFebruary2020String,
		SubCategory: "Restaurants",
		Card:        "Food allowance",
		Description: "Test",
	}
)

func Test_expenseViewToExpenseGetResponse(t *testing.T) {
	type args struct {
		expenseView dbModels.ExpenseView
	}
	tests := []struct {
		name string
		args args
		want models.ExpenseCreateRequest
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
		expenseViewRecords []dbModels.ExpenseView
	}
	tests := []struct {
		name string
		args args
		want []models.ExpenseCreateRequest
	}{
		{
			name: "Success",
			args: args{
				expenseViewRecords: []dbModels.ExpenseView{
					houseRentExpenseView, restaurantExpenseView,
				},
			},
			want: []models.ExpenseCreateRequest{
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
	cards = []dbModels.CardTable{
		{ID: 1, Name: "CGD"},
		{ID: 2, Name: "Food allowance"},
	}
	cardsCache = cache.NewCard(cards)

	categories = []dbModels.ExpenseCategoryTable{
		{ID: 1, Name: "House"},
		{ID: 2, Name: "Leisure"},
	}
	categoriesCache = cache.NewExpenseCategory(categories)

	subCategories = []dbModels.ExpenseSubCategoryTable{
		{ID: 1, Name: "Rent", CategoryID: 1},
		{ID: 2, Name: "Restaurants", CategoryID: 2},
	}
	subCategoriesCache = cache.NewExpenseSubCategory(subCategories)
)

func TestExpenses_CreateExpense(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

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
		expense models.ExpenseCreateRequest
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
			expense: models.ExpenseCreateRequest{
				Value:       200.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "Unknown",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "subcategory or card does not exist",
			},
		},
		{
			name: "ErrorUnexpectedDateFormat",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: models.ExpenseCreateRequest{
				Value:       200.0,
				Date:        "01-Feb-2020",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse date - must use YYYY-MM-DD date format",
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
			expensesHandlers.CreateExpense(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusCreated:
				var r models.ExpenseCreateResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenseID, r.ID)
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_UpdateExpense(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

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
		expense models.ExpenseCreateRequest
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
			expense: models.ExpenseCreateRequest{
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
			expense: models.ExpenseCreateRequest{
				ID:          1,
				Value:       250.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "id parameter must be an integer",
			},
		},
		{
			name: "ErrorUnknownCard",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: models.ExpenseCreateRequest{
				ID:          1,
				Value:       200.0,
				Date:        "2020-02-01",
				SubCategory: "Rent",
				Card:        "Unknown",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "subcategory or card does not exist",
			},
		},
		{
			name: "ErrorUnexpectedDateFormat",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			expense: models.ExpenseCreateRequest{
				ID:          1,
				Value:       200.0,
				Date:        "01-Feb-2020",
				SubCategory: "Rent",
				Card:        "CGD",
				Description: "House Rent",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse date - must use YYYY-MM-DD date format",
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
			expensesHandlers.UpdateExpense(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusNoContent:
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_GetExpenseByID(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
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

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expense    models.ExpenseCreateRequest
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
			name: "ErrorUnknownExpense",
			fields: fields{
				ExpensesRepository:            &expensesCache,
				CardRepository:                &cardsCache,
				ExpensesSubCategoryRepository: &subCategoriesCache,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "expense with this id does not exist",
			},
			params: map[string]string{"id": "99"},
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
				errorMsg:   "id parameter must be an integer",
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
			expensesHandlers.GetExpenseByID(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r models.ExpenseCreateRequest
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expense, r)
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_GetExpensesByCategory(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
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

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []models.ExpenseCreateRequest
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
				expenses: []models.ExpenseCreateRequest{
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
				errorMsg:   "expense category does not exist",
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
			expensesHandlers.GetExpensesByCategory(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []models.ExpenseCreateRequest
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_GetExpensesBySubCategory(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
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

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []models.ExpenseCreateRequest
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
				expenses: []models.ExpenseCreateRequest{
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
				errorMsg:   "expense subcategory does not exist",
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
			expensesHandlers.GetExpensesBySubCategory(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []models.ExpenseCreateRequest
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_GetExpensesByCard(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
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

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []models.ExpenseCreateRequest
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
				expenses: []models.ExpenseCreateRequest{
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
				errorMsg:   "expense card does not exist",
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
			expensesHandlers.GetExpensesByCard(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []models.ExpenseCreateRequest
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_GetExpensesByDates(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

	type fields struct {
		ExpensesRepository            repository.ExpenseRepo
		ExpensesSubCategoryRepository repository.ExpenseSubCategoryRepo
		CardRepository                repository.CardRepo
	}

	type want struct {
		statusCode int
		expenses   []models.ExpenseCreateRequest
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
				expenses: []models.ExpenseCreateRequest{
					houseRentExpenseHTTPModel,
					restaurantExpenseHTTPModel,
				},
			},
			params: map[string]string{"min_date": "2020-01-31", "max_date": "2060-02-02"},
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
				errorMsg:   "could not parse min date - must use YYYY-MM-DD date format",
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
				errorMsg:   "could not parse max date - must use YYYY-MM-DD date format",
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
			expensesHandlers.GetExpensesByDates(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r []models.ExpenseCreateRequest
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.expenses, r)
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}

func TestExpenses_DeleteExpense(t *testing.T) {

	expenses := []dbModels.ExpenseTable{
		houseRentExpenseTable,
		restaurantExpenseTable,
	}
	expensesCache := cache.NewExpense(expenses, cardsCache, categoriesCache, subCategoriesCache)

	expensesHandlers := NewExpenses(&expensesCache, &subCategoriesCache, &cardsCache)

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
				errorMsg:   "expense with this id does not exist",
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
				errorMsg:   "id parameter must be an integer",
			},
			params: map[string]string{"id": "abc"},
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
			expensesHandlers.DeleteExpense(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusNoContent:
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err := json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.errorMsg, r.ErrorMsg)
			}
		})
	}
}
