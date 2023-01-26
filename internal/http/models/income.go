package models

// IncomeCreateRequest is the http expense model
type IncomeCreateRequest struct {
	ID          int     `json:"id,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Date        string  `json:"date,omitempty"` // Should be on this format YYYY-MM-DD
	Category    string  `json:"category,omitempty"`
	Card        string  `json:"card,omitempty"`
	Description string  `json:"description,omitempty"`
}

// IncomeCreateResponse is the http create response model for expense
type IncomeCreateResponse struct {
	ID int `json:"id,omitempty"`
}
