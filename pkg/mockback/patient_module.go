package mockback

import (
	"context"
	"database/sql"
	"time"
)

// PatientModule contains the business logic for the patients.
type PatientModule interface {
	ListAllPatients(ctx context.Context) []Patient
}

type cockroach interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type patientModule struct {
	cockroachConn cockroach
}

func (c *patientModule) ListAllPatients(ctx context.Context) []Patient {
	// TODO : Grab actual data from the database
	var mariuszWiesniewski = Patient{ID: next(), FirstName: "Mariusz", LastName: "Wiesniwski", AVSNumber: "756.1234.3333.55",
		BirthDate: time.Date(1984, time.May, 4, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{}}
	var naimengLiu = Patient{ID: next(), FirstName: "Naimeng", LastName: "Liu", AVSNumber: "765.4321.0303.44",
		BirthDate: time.Date(1993, time.November, 11, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{}}
	var julienRoch = Patient{ID: next(), FirstName: "Julien", LastName: "Roch", AVSNumber: "333.4444.5555.66",
		BirthDate: time.Date(1984, time.December, 3, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{}}
	var christopheFrattino = Patient{ID: next(), FirstName: "Christophe", LastName: "Frattino", AVSNumber: "420.1337.1337.42",
		BirthDate: time.Date(1985, time.August, 8, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{}}

	return []Patient{mariuszWiesniewski, naimengLiu, julienRoch, christopheFrattino}
}

// NewPatientModule returns a patient module
func NewPatientModule(cockroachConn cockroach) PatientModule {
	return &patientModule{
		cockroachConn: cockroachConn,
	}
}
