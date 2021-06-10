package response

import (
	"net/http"
	"net/url"
)

// Code .
type Code string

// error codes
const (
	ErrService    Code = "ERR_SERVICE"
	ErrNotFound   Code = "ERR_NOT_FOUND"
	ErrBadRequest Code = "ERR_BAD_REQUEST"
	ErrBadParam   Code = "ERR_BAD_PARAM"
	ErrAuth       Code = "ERR_NOT_AUTHORIZED"
	ErrAccess     Code = "ERR_PERMISSION_DENIED"
)

// GetHTTPCode .
func (code Code) GetHTTPCode() int {
	switch code {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrBadRequest, ErrBadParam:
		return http.StatusBadRequest
	case ErrAuth:
		return http.StatusUnauthorized
	case ErrAccess:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// APIError .
type APIError struct {
	Code    Code   `json:"status"`
	Message string `json:"code"`
	err     error
}

// ValidationError .
type ValidationError struct {
	APIError
	Errors url.Values
}

// Error .
func (err APIError) Error() string {
	if err.err != nil {
		return err.err.Error()
	}
	return err.Message
}

// NewBadJwtError .
func NewBadJwtError(err error) APIError {
	return APIError{
		Code:    ErrAuth,
		Message: "Unauthorized",
		err:     err,
	}
}

// NewBadRequestError .
func NewBadRequestError(err error) APIError {
	return APIError{
		Code:    ErrBadRequest,
		Message: "Bad Request",
		err:     err,
	}
}

// NewAuthorizationError .
func NewAuthorizationError(err error) APIError {
	return APIError{
		Code:    ErrAuth,
		Message: "Unauthorized",
		err:     err,
	}
}

// NewNotFoundError .
func NewNotFoundError() APIError {
	return APIError{
		Code:    ErrNotFound,
		Message: "Not Found",
	}
}

// NewValidationError .
func NewValidationError(err url.Values) ValidationError {
	return ValidationError{
		APIError: APIError{
			Code:    ErrBadParam,
			Message: "Bad Request",
		},
		Errors: err,
	}
}
