package http

type Expense struct {
	Id          int     `json:"id,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Date        string  `json:"date,omitempty"` // Should be on this format YYYY-MM-DD
	SubCategory string  `json:"sub_category,omitempty"`
	Card        string  `json:"card,omitempty"`
	Description string  `json:"description,omitempty"`
}
