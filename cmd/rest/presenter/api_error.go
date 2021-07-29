package presenter

import (
	"micheltank/cryptocurrency-data-service/internal/domain"
)

type ApiError struct {
	Message string `json:"message"`
	Key     string `json:"key"`
	Detail  string `json:"detail,omitempty"`
}

func NewApiError(errorDomain domain.Error) ApiError {
	return ApiError{
		Message: errorDomain.Error(),
		Key:     errorDomain.Key(),
		Detail:  errorDomain.Detail(),
	}
}