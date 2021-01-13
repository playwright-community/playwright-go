package playwright

type Error struct {
	Message string
	Stack   string
}

// Error represents a Playwright error
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
			Message: err.Message,
			Stack:   err.Stack,
		}
	}
	return &Error{
		Message: err.Message,
		Stack:   err.Stack,
	}
}
