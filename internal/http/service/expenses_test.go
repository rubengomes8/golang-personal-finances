package service

import (
	"reflect"
	"testing"
	"time"

	httpModels "github.com/rubengomes8/golang-personal-finances/internal/models/http"
	rdsModels "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
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
		Category:      "Laser",
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
