package godriver

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
}

type DriverPatch struct {
	Name      *string    `json:"name,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
}
