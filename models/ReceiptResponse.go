package models

import "github.com/google/uuid"

type ReceiptResponse struct {
	ID     uuid.UUID `json:"id"`
	Points int       `json:"points"`
}
