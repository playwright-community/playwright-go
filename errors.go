package playwright

import (
	"errors"
	"fmt"
	"strings"
)

// Error represents a Playwright error
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
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

var (
	ErrPlaywright             = errors.New("playwright")
	ErrTargetClosed           = errors.New("target closed")
	ErrBrowserClosed          = errors.New("Browser has been closed")
	ErrBrowserOrContextClosed = errors.New("Target page, context or browser has been closed")
)

// TimeoutError represents a Playwright TimeoutError
var TimeoutError = &Error{
	Name: "TimeoutError",
}

func parseError(err Error) error {
	return fmt.Errorf("%w: %w", ErrPlaywright, &Error{
		Name:    err.Name,
		Message: err.Message,
		Stack:   err.Stack,
	})
}

const (
	errMsgBrowserClosed          = "Browser has been closed"
	errMsgBrowserOrContextClosed = "Target page, context or browser has been closed"
)

func isTargetClosedError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrTargetClosed) || errors.Is(err, ErrBrowserClosed) || errors.Is(err, ErrBrowserOrContextClosed)
}

func targetClosedError(reason *string) error {
	if reason == nil {
		return ErrTargetClosed
	}
	return fmt.Errorf("%w: %s", ErrTargetClosed, *reason)
}

func isSafeCloseError(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasSuffix(err.Error(), errMsgBrowserClosed) || strings.HasSuffix(err.Error(), errMsgBrowserOrContextClosed)
}
