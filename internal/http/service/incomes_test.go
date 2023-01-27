package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

// repo mocks
var (
	incomeRepo   = mock.NewIncome()
	categoryRepo = mock.NewIncomeCategory()
	cardRepo     = mock.NewCard()
)

// http
var (
	salaryIncomeCreateRequest = models.Income{
		Value:       mock.IncomeSalary.Value,
		Date:        mock.IncomeSalary.Date.Format("2006-01-02"),
		Category:    mock.IncomeSalaryCategory.Name,
		Card:        mock.IncomeSalaryCard.Name,
		Description: mock.IncomeSalary.Description,
	}

	salaryIncomeCreateResponse = models.IncomeCreateResponse{
		ID: int(mock.IncomeSalary.ID),
	}
)

func TestIncomes_CreateIncome(t *testing.T) {

	incomesService, err := NewIncomes(incomeRepo, categoryRepo, cardRepo)
	if err != nil {
		t.Fatalf("error creating incomes service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

	type fields struct {
		Repository         repository.IncomeRepo
		CategoryRepository repository.IncomeCategoryRepo
		CardRepository     repository.CardRepo
	}

	type want struct {
		statusCode int
		response   models.IncomeCreateResponse
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		income models.Income
		want   want
	}{
		{
			name: "Success",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: salaryIncomeCreateRequest,
			want: want{
				statusCode: http.StatusCreated,
				response:   salaryIncomeCreateResponse,
			},
		},
		{
			name: "ErrorCardDoesNotExist",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card: "Unknown",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "card does not exist",
			},
		},
		{
			name: "ErrorCategoryDoesNotExist",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:     salaryIncomeCreateRequest.Card,
				Category: "Unknown",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "income category does not exist",
			},
		},
		{
			name: "ErrorCouldNotParseDate",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:     salaryIncomeCreateRequest.Card,
				Category: salaryIncomeCreateRequest.Category,
				Date:     "WrongFormat",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse date - must use YYYY-MM-DD date format",
			},
		},
		{
			name: "ErrorCouldNotCreate",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:        mock.IncomeSalaryCard.Name,
				Category:    mock.IncomeSalaryCategory.Name,
				Date:        salaryIncomeCreateRequest.Date,
				Description: "IsNotMock",
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				errorMsg:   "could not create income",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			data, err := json.Marshal(tt.income)
			if err != nil {
				t.Fatalf("error marshaling income: %v\n", err)
			}

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Request = &http.Request{
				Method: http.MethodPost,
				Body:   io.NopCloser(bytes.NewBuffer(data)),
			}

			// WHEN
			incomesService.CreateIncome(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusCreated:
				var r models.IncomeCreateResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.response, r)
			case http.StatusBadRequest, http.StatusInternalServerError:
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

func TestIncomes_UpdateIncome(t *testing.T) {

	incomesService, err := NewIncomes(incomeRepo, categoryRepo, cardRepo)
	if err != nil {
		t.Fatalf("error creating incomes service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

	type fields struct {
		Repository         repository.IncomeRepo
		CategoryRepository repository.IncomeCategoryRepo
		CardRepository     repository.CardRepo
	}

	type want struct {
		statusCode int
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		income models.Income
		want   want
		params map[string]string
	}{
		{
			name: "Success",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: salaryIncomeCreateRequest,
			params: map[string]string{"id": "1"},
			want: want{
				statusCode: http.StatusNoContent,
			},
		},
		{
			name: "ErrorCardDoesNotExist",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card: "Unknown",
			},
			params: map[string]string{"id": "1"},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "card does not exist",
			},
		},
		{
			name: "ErrorCategoryDoesNotExist",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:     salaryIncomeCreateRequest.Card,
				Category: "Unknown",
			},
			params: map[string]string{"id": "1"},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "income category does not exist",
			},
		},
		{
			name: "ErrorCouldNotParseDate",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:     salaryIncomeCreateRequest.Card,
				Category: salaryIncomeCreateRequest.Category,
				Date:     "WrongFormat",
			},
			params: map[string]string{"id": "1"},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "could not parse date - must use YYYY-MM-DD date format",
			},
		},
		{
			name: "ErrorParamIDNotInt",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:        mock.IncomeSalaryCard.Name,
				Category:    mock.IncomeSalaryCategory.Name,
				Date:        salaryIncomeCreateRequest.Date,
				Description: "IsNotMock",
			},
			params: map[string]string{"id": "abc"},
			want: want{
				statusCode: http.StatusBadRequest,
				errorMsg:   "id parameter must be an integer",
			},
		},
		{
			name: "ErrorIncomeDoesNotExist",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			income: models.Income{
				Card:        mock.IncomeSalaryCard.Name,
				Category:    mock.IncomeSalaryCategory.Name,
				Date:        salaryIncomeCreateRequest.Date,
				Description: "IsNotMock",
			},
			params: map[string]string{"id": "2"},
			want: want{
				statusCode: http.StatusInternalServerError,
				errorMsg:   "incomes with this id does not exist",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// GIVEN
			data, err := json.Marshal(tt.income)
			if err != nil {
				t.Fatalf("error marshaling income: %v\n", err)
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
			incomesService.UpdateIncome(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusNoContent:
			case http.StatusBadRequest, http.StatusInternalServerError:
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

func TestIncomes_GetIncomeByID(t *testing.T) {

	incomesService, err := NewIncomes(incomeRepo, categoryRepo, cardRepo)
	if err != nil {
		t.Fatalf("error creating incomes service: %v\n", err)
	}

	gin.SetMode(gin.TestMode)

	type fields struct {
		Repository         repository.IncomeRepo
		CategoryRepository repository.IncomeCategoryRepo
		CardRepository     repository.CardRepo
	}

	type want struct {
		statusCode int
		income     models.Income
		errorMsg   string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
		params map[string]string
	}{
		// TODO:
		/*{
			name: "Success",
			fields: fields{
				Repository:         &incomeRepo,
				CardRepository:     &cardRepo,
				CategoryRepository: &categoryRepo,
			},
			want: want{
				statusCode: http.StatusOK,
				income:     mock.IncomeSalary,
			},
			params: map[string]string{"id": "1"},
		},*/
	}
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
			incomesService.GetIncomeByID(ginCtx)

			// THEN
			assert.EqualValues(t, tt.want.statusCode, w.Code)

			switch w.Code {
			case http.StatusOK:
				var r models.Income
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, tt.want.income, r)
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
