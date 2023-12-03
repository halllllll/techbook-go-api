package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) UnWrap() error {
	return myErr.Err
}
