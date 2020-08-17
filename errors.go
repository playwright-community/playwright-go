package playwright

type Error struct {
	Message string
	Stack   string
}

func (e *Error) Error() string {
	return e.Message
}

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
