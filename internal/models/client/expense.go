package models

type Expense struct {
	Value       float64 `json:"value,omitempty"`
	Date        int64   `json:"date,omitempty"`
	SubCategory string  `json:"sub_category,omitempty"`
	Card        string  `json:"card,omitempty"`
	Description string  `json:"description,omitempty"`
}
