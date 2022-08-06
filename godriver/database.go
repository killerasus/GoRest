package godriver

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Database struct {
	Drivers        []Driver
	Passengers     []Passenger
	TravelRequests []TravelRequest
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

func (db *Database) PatchDriver(id uuid.UUID, patch DriverPatch) error {
	d, err := db.GetDriver(id)
	if err != nil {
		return err
	}

	if patch.Name != nil {
		d.Name = *patch.Name
	}
	if patch.Birthdate != nil {
		d.Birthdate = *patch.Birthdate
	}

	if _, err = db.UpdateDriver(id, d); err != nil {
		return err
	}

	return nil
}

func (db *Database) CreatePassenger(p Passenger) Passenger {
	p.ID = uuid.New()
	db.Passengers = append(db.Passengers, p)
	return p
}

func (db *Database) GetPassenger(id uuid.UUID) (Passenger, error) {
	for i, p := range db.Passengers {
		if p.ID == id {
			return db.Passengers[i], nil
		}
	}
	return Passenger{}, errors.New("Passenger " + id.String() + " not found.")
}

func (db *Database) RemovePassenger(id uuid.UUID) (Passenger, error) {
	for i, p := range db.Passengers {
		if p.ID == id {
			passenger := db.Passengers[i]
			db.Passengers = append(db.Passengers[:i], db.Passengers[i+1:]...)
			return passenger, nil
		}
	}

	return Passenger{}, errors.New("Passenger " + id.String() + " not found.")
}

func (db *Database) UpdatePassenger(id uuid.UUID, passenger Passenger) (Passenger, error) {
	p, err := db.RemovePassenger(id)
	if err == nil {
		passenger.ID = id
		db.Passengers = append(db.Passengers, passenger)
		return passenger, nil
	}

	return p, err
}

func (db *Database) PatchPassenger(id uuid.UUID, patch PassengerPatch) error {
	p, err := db.GetPassenger(id)
	if err != nil {
		return err
	}

	if patch.Name != nil {
		p.Name = *patch.Name
	}

	if _, err = db.UpdatePassenger(id, p); err != nil {
		return err
	}

	return nil
}

func (db *Database) SaveTravelRequest(t TravelRequestInput) (TravelRequest, error) {
	var tr TravelRequest

	_, err := db.GetPassenger(*t.Passenger)
	if err != nil {
		return tr, err
	}

	if *t.Origin == "" {
		return tr, errors.New("origin unknown")
	}

	if *t.Destination == "" {
		return tr, errors.New("destination unknown")
	}

	tr.ID = uuid.New()
	tr.Passenger = *t.Passenger
	tr.Origin = *t.Origin
	tr.Destination = *t.Destination
	tr.Status = Created
	tr.CreationDate = time.Now()
	db.TravelRequests = append(db.TravelRequests, tr)

	return tr, nil
}
