package apierror

import "fmt"

type APIError struct {
	StatusCode int    // HTTP status to return
	Message    string // safe message for the client
	Err        error  // underlying error (optional)
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewUnauthorized(msg string, err error) *APIError {
	return &APIError{StatusCode: 401, Message: msg, Err: err}
}

func NewForbidden(msg string, err error) *APIError {
	return &APIError{StatusCode: 403, Message: msg, Err: err}
}

func NewConflict(msg string, err error) *APIError {
	return &APIError{StatusCode: 409, Message: msg, Err: err}
}

func NewUnprocessableEntity(msg string, err error) *APIError {
	return &APIError{StatusCode: 422, Message: msg, Err: err}
}

func NewTooManyRequests(msg string, err error) *APIError {
	return &APIError{StatusCode: 429, Message: msg, Err: err}
}

func NewInternal(msg string, err error) *APIError {
	return &APIError{StatusCode: 500, Message: msg, Err: err}
}

func NewBadRequest(msg string, err error) *APIError {
	return &APIError{StatusCode: 400, Message: msg, Err: err}
}

func NewNotFound(msg string, err error) *APIError {
	return &APIError{StatusCode: 404, Message: msg, Err: err}
}

func NewUnsupportedMedia(msg string, err error) *APIError {
	return &APIError{StatusCode: 415, Message: msg, Err: err}
}
