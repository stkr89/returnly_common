package types

import (
	"github.com/google/uuid"
	"time"
)

type ReturnInitiated struct {
	ReturnID   uuid.UUID `json:"returnID"`
	PickupDate time.Time `json:"pickupDate"`
}
