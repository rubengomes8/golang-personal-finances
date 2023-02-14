package expense

import "errors"

var (
	ErrNoRowsAffectedOnDelete            = errors.New("there were no rows affected in exec expense delete statement")
	ErrNoRowsAffectedOnUpdate            = errors.New("there were no rows affected in exec expense update statement")
	ErrNoRowsAffectedOnCategoryDelete    = errors.New("there were no rows affected in exec expense category delete statement")
	ErrNoRowsAffectedOnSubcategoryDelete = errors.New("there were no rows affected in exec expense subcategory delete statement")
)
