package http_handlers

type CustomError struct {
	ErrorText string
}

func (e *CustomError) Error() string {
	return e.ErrorText
}

var (
	ErrorJson    = CustomError{"Invalid JSON request"}
	ErrorService = CustomError{"Service no execute request"}
	ErrorOther   = CustomError{"Other error"}
)
