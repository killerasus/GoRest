package godriver

import "github.com/google/uuid"

type Passenger struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type PassengerPatch struct {
	Name *string `json:"name,omitempty"`
}
