package types

import "github.com/google/uuid"

type ReturnInitiated struct {
	ReturnID uuid.UUID `json:"returnID"`
}
