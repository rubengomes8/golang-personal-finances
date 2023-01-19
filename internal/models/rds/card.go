package rds

// CardTable is the rds card model
type CardTable struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
