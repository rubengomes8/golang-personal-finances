package card

import "errors"

var (
	ErrNoRowsAffectedOnDelete = errors.New("there were no rows affected in exec expense card delete statement")
)
