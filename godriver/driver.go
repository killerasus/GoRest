package godriver

import "time"

type Driver struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
}
