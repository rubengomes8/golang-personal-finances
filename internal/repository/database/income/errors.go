package income

import "errors"

var (
	ErrNoRowsAffectedOnDelete         = errors.New("there were no rows affected in exec income delete statement")
	ErrNoRowsAffectedOnUpdate         = errors.New("there were no rows affected in exec income update statement")
	ErrNoRowsAffectedOnCategoryDelete = errors.New("there were no rows affected in exec income category delete statement")
)
