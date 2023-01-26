package models

// UserTable is the rds user model
type UserTable struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Passhash string `json:"passhash,omitempty"`
}
