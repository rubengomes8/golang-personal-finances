package incomes

import "errors"

var (
	ErrCardNotFound       = errors.New("card does not exist")
	ErrCategoryNotFound   = errors.New("income category does not exist")
	ErrNotFound           = errors.New("income does not exist")
	ErrCouldNotGetByDates = errors.New("could not get incomes by dates")
	ErrCouldNotParseTime  = errors.New("could not parse time")
	ErrCouldNotInsert     = errors.New("could not insert income")
	ErrCouldNotUpdate     = errors.New("could not update income")
	ErrCouldNotDelete     = errors.New("could not delete income")
)
