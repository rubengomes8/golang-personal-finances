package models

type Expense struct {
	Id          int64   `json:"id,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Date        int64   `json:"date,omitempty"`
	Category    string  `json:"category,omitempty"`
	SubCategory string  `json:"sub_category,omitempty"`
	Card        string  `json:"card,omitempty"`
	Description string  `json:"description,omitempty"`
}

type ExpenseCategory struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExpenseSubCategory struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CategoryId int64  `json:"category_id,omitempty"`
}

type ExpenseWithIDs struct {
	Id            int64   `json:"id,omitempty"`
	Value         float64 `json:"value,omitempty"`
	Date          int64   `json:"date,omitempty"`
	CategoryId    int64   `json:"category_id,omitempty"`
	SubCategoryId int64   `json:"sub_category_id,omitempty"`
	CardId        int64   `json:"card_id,omitempty"`
	Description   string  `json:"description,omitempty"`
}