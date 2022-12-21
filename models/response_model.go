package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{Message: message}
}

type BasketSummary struct {
	BasketId      uuid.UUID
	CountProducts int
	TotalPrice    float64
}

func (e *ErrorResponse) Error() string {
	errorString, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("failed to serialize the err: %v", err)
	}

	return string(errorString)
}
