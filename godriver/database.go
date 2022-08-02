package godriver

import (
	"errors"

	"github.com/google/uuid"
)

type Database struct {
	Drivers []Driver
}

func (db *Database) CreateDriver(d Driver) Driver {
	d.ID = uuid.New()
	db.Drivers = append(db.Drivers, d)
	return d
}

func (db *Database) GetDriver(id uuid.UUID) (Driver, error) {
	for i, d := range db.Drivers {
		if d.ID == id {
			return db.Drivers[i], nil
		}
	}
	return Driver{}, errors.New("Driver " + id.String() + " not found.")
}

func (db *Database) RemoveDriver(id uuid.UUID) (Driver, error) {
	for i, d := range db.Drivers {
		if d.ID == id {
			driver := db.Drivers[i]
			db.Drivers = append(db.Drivers[:i], db.Drivers[i+1:]...)
			return driver, nil
		}
	}

	return Driver{}, errors.New("Driver " + id.String() + " not found.")
}

func (db *Database) UpdateDriver(id uuid.UUID, driver Driver) (Driver, error) {
	d, err := db.RemoveDriver(id)
	if err == nil {
		driver.ID = id
		db.Drivers = append(db.Drivers, driver)
		return driver, nil
	}

	return d, err
}
