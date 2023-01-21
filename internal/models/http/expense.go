package http

// Expense is the http expense model
type Expense struct {
	ID          int     `json:"id,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Date        string  `json:"date,omitempty"` // Should be on this format YYYY-MM-DD
	SubCategory string  `json:"sub_category,omitempty"`
	Card        string  `json:"card,omitempty"`
	Description string  `json:"description,omitempty"`
}

// ExpenseCreateResponse is the http create response model for expense
type ExpenseCreateResponse struct {
	ID int `json:"id,omitempty"`
}
