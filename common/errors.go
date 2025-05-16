package common

type UserError struct {
	Code          string
	Detail        any
	InternalError error
}

func NewUserErrorWithDetail(code string, reason error, detail any) *UserError {
	return &UserError{
		Code:          code,
		Detail:        detail,
		InternalError: reason,
	}
}

func NewUserError(code string, reason error) *UserError {
	return &UserError{
		Code:          code,
		InternalError: reason,
	}
}

func (e *UserError) Error() string {
	return e.Code
}

const ErrInternalServerError = "INTERNAL_SERVER_ERROR"
