package patients

import (
	"context"
	"database/sql"
	"time"
)

// Patient represents a patient
type Patient struct {
	ID         int32     `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	BirthDate  time.Time `json:"birthDate"`
	AVSNumber  string    `json:"avsNumber"`
	DoctorsIds []int32   `json:"-"`
}

// Module contains the business logic for the patients.
type Module interface {
	ListAll(ctx context.Context) []Patient
}

type module struct {
	cockroachConn cockroach
}

type cockroach interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func (c *module) ListAll(ctx context.Context) []Patient {
	// TODO : Grab actual data from the database
	return []Patient{Patient{ID: 0, FirstName: "Mariusz", LastName: "Wiesniwski", AVSNumber: "756.1234.3333.55",
		BirthDate: time.Date(1984, time.May, 4, 0, 0, 0, 0, time.UTC), DoctorsIds: []int32{}}, Patient{ID: 1, FirstName: "Naimeng", LastName: "Liu", AVSNumber: "765.4321.0303.44",
		BirthDate: time.Date(1993, time.November, 11, 0, 0, 0, 0, time.UTC), DoctorsIds: []int32{}}, Patient{ID: 2, FirstName: "Julien", LastName: "Roch", AVSNumber: "333.4444.5555.66",
		BirthDate: time.Date(1984, time.December, 3, 0, 0, 0, 0, time.UTC), DoctorsIds: []int32{}}, Patient{ID: 3, FirstName: "Christophe", LastName: "Frattino", AVSNumber: "420.1337.1337.42",
		BirthDate: time.Date(1985, time.August, 8, 0, 0, 0, 0, time.UTC), DoctorsIds: []int32{}}}
}

// NewModule returns a patient module
func NewModule(cockroachConn cockroach) Module {
	return &module{
		cockroachConn: cockroachConn,
	}
}
