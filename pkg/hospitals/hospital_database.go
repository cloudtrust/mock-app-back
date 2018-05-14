package hospitals

import (
	"database/sql"

	"github.com/pkg/errors"
)

const (
	createHospitalsTblStmt = `CREATE TABLE IF NOT EXISTS hospitals (
		id INT,
		name STRING,
		city STRING,
		PRIMARY KEY (id))`
	selectAllHospitalsTblStmt = `SELECT * FROM hospitals`
	createDepartmentsTblStmt  = `CREATE TABLE IF NOT EXISTS departments (
		id INT,
		name STRING,
		hospital_id INT,
		PRIMARY KEY (id))`
	selectAllDepartmentsTblStmt = `SELECT * FROM departments`
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
	// Init DB: create hospitals table.
	var _, err = db.Exec(createHospitalsTblStmt)
	if err != nil {
		return nil, err
	}

	// Init DB: create departments table.
	_, err = db.Exec(createDepartmentsTblStmt)
	if err != nil {
		return nil, err
	}

	return &CockroachModule{
		db: db,
	}, nil
}

// ReadHospitalsFromDb returns all the hospitals from the database
func (c *CockroachModule) ReadHospitalsFromDb() ([]Hospital, error) {
	var hospitals = []Hospital{}
	var rows, err = c.db.Query(selectAllHospitalsTblStmt)
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the hospitals from the database.")
	}
	var (
		id         int32
		name, city string
	)
	defer rows.Close()
	for rows.Next() {
		var err = rows.Scan(&id, &name, &city)
		if err != nil {
			return nil, errors.Wrapf(err, "error while returning all the hospitals from the database.")
		}
		hospitals = append(hospitals, Hospital{
			ID:   id,
			Name: name,
			City: city,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the hospitals from the database.")
	}

	return hospitals, nil
}

// ReadDepartmentsFromDb returns all the departments from the database
func (c *CockroachModule) ReadDepartmentsFromDb() ([]Department, error) {
	var departments = []Department{}
	var rows, err = c.db.Query(selectAllDepartmentsTblStmt)
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the departments from the database.")
	}
	var (
		id, hospitalID int32
		name           string
	)
	defer rows.Close()
	for rows.Next() {
		var err = rows.Scan(&id, &name, &hospitalID)
		if err != nil {
			return nil, errors.Wrapf(err, "error while returning all the departments from the database.")
		}
		departments = append(departments, Department{
			ID:   id,
			Name: name,
			Hospital: Hospital{
				ID: hospitalID,
			},
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the departments from the database.")
	}

	return departments, nil
}
