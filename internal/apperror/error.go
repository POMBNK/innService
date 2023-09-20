package apperror

import (
	"encoding/json"
	"fmt"
)

type AppError struct {
	Err     error
	Code    int
	Message string
}

func (e *AppError) Error() string {
	err := e.Err.Error()
	return err
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return bytes
}

func NewAppError(transportCode int, message string) *AppError {
	return &AppError{
		Err:     fmt.Errorf(message),
		Code:    transportCode,
		Message: message,
	}
}
