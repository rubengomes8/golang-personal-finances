package rds

import "time"

// ExpenseView is the rds expense view model
type ExpenseView struct {
	Id            int64     `json:"id,omitempty"`
	Value         float64   `json:"value,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	Category      string    `json:"category,omitempty"`
	SubCategory   string    `json:"sub_category,omitempty"`
	Card          string    `json:"card,omitempty"`
	CategoryId    int64     `json:"category_id,omitempty"`
	SubCategoryId int64     `json:"sub_category_id,omitempty"`
	CardId        int64     `json:"card_id,omitempty"`
	Description   string    `json:"description,omitempty"`
}

// ExpenseTable is the rds expense table model
type ExpenseTable struct {
	Id            int64     `json:"id,omitempty"`
	Value         float64   `json:"value,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	SubCategoryId int64     `json:"sub_category_id,omitempty"`
	CardId        int64     `json:"card_id,omitempty"`
	Description   string    `json:"description,omitempty"`
}

// ExpenseCategoryTable is the rds expense category table model
type ExpenseCategoryTable struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ExpenseSubCategoryTable is the rds expense subcategory table model
type ExpenseSubCategoryTable struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CategoryId int64  `json:"category_id,omitempty"`
}
