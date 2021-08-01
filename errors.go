package playwright

// Error represents a Playwright error
type Error struct {
	Name    string
	Message string
	Stack   string
}

func (e *Error) Error() string {
	return e.Message
}

// TimeoutError represents a Playwright TimeoutError
type TimeoutError Error

func (e *TimeoutError) Error() string {
	return e.Message
}

func parseError(err errorPayload) error {
	if err.Name == "TimeoutError" {
		return &TimeoutError{
			Name:    "TimeoutError",
			Message: err.Message,
			Stack:   err.Stack,
		}
	}
	return &Error{
		Name:    err.Name,
		Message: err.Message,
		Stack:   err.Stack,
	}
}
