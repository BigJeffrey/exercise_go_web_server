package models

import "github.com/google/uuid"

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
