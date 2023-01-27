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

// http body requests
var (
	salaryIncomeCreateRequest = models.IncomeCreateRequest{
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
		t.Fatalf("error creating expenses service: %v\n", err)
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
		income models.IncomeCreateRequest
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
			income: models.IncomeCreateRequest{
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
			income: models.IncomeCreateRequest{
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
			income: models.IncomeCreateRequest{
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
			income: models.IncomeCreateRequest{
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
			case http.StatusBadRequest:
				var r models.ErrorResponse
				err = json.NewDecoder(w.Body).Decode(&r)
				if err != nil {
					t.Fatalf("error decoding response: %v\n", err)
				}
				assert.Equal(t, r.ErrorMsg, tt.want.errorMsg)
			}
		})
	}
}
