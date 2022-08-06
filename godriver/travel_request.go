package godriver

import (
	"time"

	"github.com/google/uuid"
)

type TravelRequestStatus uint8

const (
	Created TravelRequestStatus = iota
	Accepted
	Refused
)

type TravelRequestInput struct {
	Passenger   *uuid.UUID `json:"passenger,omitempty"`
	Origin      *string    `json:"origin,omitempty"`
	Destination *string    `json:"destination,omitempty"`
}

type TravelRequest struct {
	ID           uuid.UUID           `json:"id"`
	Passenger    uuid.UUID           `json:"passenger"`
	Origin       string              `json:"origin"`
	Destination  string              `json:"destination"`
	Status       TravelRequestStatus `json:"status"`
	CreationDate time.Time           `json:"date"`
}
