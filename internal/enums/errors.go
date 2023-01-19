package enums

import "errors"

var (
	NoRowsAffectedCardDeleteErr           = errors.New("there were no rows affected in exec expense card delete statement")
	NoRowsAffectedExpenseDeleteErr        = errors.New("there were no rows affected in exec expense delete statement")
	NoRowsAffectedExpenseUpdateErr        = errors.New("there were no rows affected in exec expense update statement")
	NoRowsAffectedExpCategoryDeleteErr    = errors.New("there were no rows affected in exec expense category delete statement")
	NoRowsAffectedExpSubcategoryDeleteErr = errors.New("there were no rows affected in exec expense subcategory delete statement")
)
