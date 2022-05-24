package types

import (
	"github.com/google/uuid"
	"time"
)

type PickupAccepted struct {
	ReturnID uuid.UUID `json:"returnID"`
	DriverID uuid.UUID `json:"driverID"`
}

type ReturnInitiated struct {
	ReturnID   uuid.UUID `json:"returnID"`
	PickupDate time.Time `json:"pickupDate"`
}
