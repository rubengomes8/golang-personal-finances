package models

// CardTable is the rds card model
type CardTable struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
