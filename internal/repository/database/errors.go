package database

import "errors"

var (
	ErrNoRowsAffectedCardDelete = errors.New("there were no rows affected in exec expense card delete statement")

	ErrNoRowsAffectedExpenseDelete        = errors.New("there were no rows affected in exec expense delete statement")
	ErrNoRowsAffectedExpenseUpdate        = errors.New("there were no rows affected in exec expense update statement")
	ErrNoRowsAffectedExpCategoryDelete    = errors.New("there were no rows affected in exec expense category delete statement")
	ErrNoRowsAffectedExpSubcategoryDelete = errors.New("there were no rows affected in exec expense subcategory delete statement")

	ErrNoRowsAffectedIncomeDelete      = errors.New("there were no rows affected in exec income delete statement")
	ErrNoRowsAffectedIncomeUpdate      = errors.New("there were no rows affected in exec income update statement")
	ErrNoRowsAffectedIncCategoryDelete = errors.New("there were no rows affected in exec income category delete statement")
)