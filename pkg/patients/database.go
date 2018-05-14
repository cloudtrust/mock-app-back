package patients

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

const (
	createPatientsTblStmt = `CREATE TABLE IF NOT EXISTS patients (
		id INT,
		first_name STRING,
		last_name STRING,
		birth_date DATE,
		avs_number STRING,
		PRIMARY KEY (id))`
	selectAllPatientsTblStmt = `SELECT * FROM patients`
)

// CockroachModule deals with the communication with CockroachDB.
type CockroachModule struct {
	db cockroachDB
}
type cockroachDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// InitDatabase initializes the database
func InitDatabase(db cockroachDB) (*CockroachModule, error) {
	// Init DB: create patients table.
	var _, err = db.Exec(createPatientsTblStmt)
	if err != nil {
		return nil, err
	}
	return &CockroachModule{
		db: db,
	}, nil
}

// ReadFromDb returns all the patients from the database
func (c *CockroachModule) ReadFromDb() ([]Patient, error) {
	var patients = []Patient{}
	var rows, err = c.db.Query(selectAllPatientsTblStmt)
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the patients from the database.")
	}
	var (
		id                             int32
		firstName, lastName, avsNumber string
		birthDate                      time.Time
	)
	defer rows.Close()
	for rows.Next() {
		var err = rows.Scan(&id, &firstName, &lastName, &birthDate, &avsNumber)
		if err != nil {
			return nil, errors.Wrapf(err, "error while returning all the patients from the database.")
		}
		patients = append(patients, Patient{
			ID:         id,
			FirstName:  firstName,
			LastName:   lastName,
			BirthDate:  birthDate,
			AVSNumber:  avsNumber,
			DoctorsIds: []int32{},
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the patients from the database.")
	}

	return patients, nil
}
