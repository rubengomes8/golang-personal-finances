package models

import "time"

// ExpenseView is the db expense view model
type ExpenseView struct {
	ID            int64     `json:"id,omitempty"`
	Value         float64   `json:"value,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	Category      string    `json:"category,omitempty"`
	SubCategory   string    `json:"sub_category,omitempty"`
	Card          string    `json:"card,omitempty"`
	CategoryID    int64     `json:"category_id,omitempty"`
	SubCategoryID int64     `json:"sub_category_id,omitempty"`
	CardID        int64     `json:"card_id,omitempty"`
	Description   string    `json:"description,omitempty"`
}

// ExpenseTable is the db expense table model
type ExpenseTable struct {
	ID            int64     `json:"id,omitempty"`
	Value         float64   `json:"value,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	SubCategoryID int64     `json:"sub_category_id,omitempty"`
	CardID        int64     `json:"card_id,omitempty"`
	Description   string    `json:"description,omitempty"`
}

// ExpenseCategoryTable is the db expense category table model
type ExpenseCategoryTable struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ExpenseSubCategoryTable is the db expense subcategory table model
type ExpenseSubCategoryTable struct {
	ID         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CategoryID int64  `json:"category_id,omitempty"`
}
