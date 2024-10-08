package response

import (
	"fmt"
)

type AppError struct {
	StatusCode  int    `json:"statusCode"`
	Description string `json:"description,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf(e.Description)
}

func NewError(statusCode int, description string) *AppError {
	return &AppError{
		StatusCode:  statusCode,
		Description: description,
	}
}
