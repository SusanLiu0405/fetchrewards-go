package models

import "github.com/google/uuid"

type ReceiptIdResponse struct {
	ID uuid.UUID `json:"id"`
}
