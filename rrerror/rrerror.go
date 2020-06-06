package rrerror

// RRError is the custom error used for RR-specific errors
type RRError struct {
	message string
	code    int
}

// Error returns an error message
func (e *RRError) Error() string {
	return e.message
}

// Code returns a status code
func (e *RRError) Code() int {
	return e.code
}

// New returns a new RRError
func New(message string, code int) *RRError {
	e := RRError{}
	e.message = message
	e.code = code

	return &e
}
