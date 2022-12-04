package errors

var (
	InvalidElementError int = 1
	JsonMarshalError    int = 2
)

type CustomError struct {
	Err  error
	Code int
}
