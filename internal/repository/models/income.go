package models

import "time"

// IncomeView is the db expense view model
type IncomeView struct {
	ID          int64     `json:"id,omitempty"`
	Value       float64   `json:"value,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Category    string    `json:"category,omitempty"`
	Card        string    `json:"card,omitempty"`
	CategoryID  int64     `json:"category_id,omitempty"`
	CardID      int64     `json:"card_id,omitempty"`
	Description string    `json:"description,omitempty"`
}

// IncomeTable is the db expense table model
type IncomeTable struct {
	ID          int64     `json:"id,omitempty"`
	Value       float64   `json:"value,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	CategoryID  int64     `json:"category_id,omitempty"`
	CardID      int64     `json:"card_id,omitempty"`
	Description string    `json:"description,omitempty"`
}

// IncomeCategoryTable is the db expense category table model
type IncomeCategoryTable struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
