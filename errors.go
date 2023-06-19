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

func (e *Error) Is(target error) bool {
	err, ok := target.(*Error)
	if !ok {
		return false
	}
	if err.Name != e.Name {
		return false
	}
	if e.Name != "Error" {
		return true // same name and not normal error
	}
	return e.Message == err.Message
}

// TimeoutError represents a Playwright TimeoutError
var TimeoutError = &Error{
	Name: "TimeoutError",
}

func parseError(err errorPayload) error {
	return &Error{
		Name:    err.Name,
		Message: err.Message,
		Stack:   err.Stack,
	}
}
