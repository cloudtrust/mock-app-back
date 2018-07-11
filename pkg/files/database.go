package files

import (
	"database/sql"

	"github.com/pkg/errors"
)

const (
	createFilesTblStmt = `CREATE TABLE IF NOT EXISTS files (
		id INT,
		patient_avs_number STRING,
		doctor_id INT,
		data STRING,
		PRIMARY KEY (id))`
	selectAllFilesTblStmt  = `SELECT * FROM files`
	selectSomeFilesTblStmt = `SELECT * FROM files OFFSET $1 LIMIT $2`
	countFilesTblStmt      = `SELECT count(*) FROM files`
)

// Database deals with the communication with CockroachDB.
type Database struct {
	db cockroachDB
}
type cockroachDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// InitDatabase initializes the database
func InitDatabase(db cockroachDB) (*Database, error) {
	// Init DB: create files table.
	var _, err = db.Exec(createFilesTblStmt)
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

// ReadFromDb returns all the files from the database
func (c *Database) ReadFromDb(first int32, count int32) ([]File, error) {
	var files = []File{}
	var rows *sql.Rows
	var err error
	if first != -1 && count != -1 {
		rows, err = c.db.Query(selectSomeFilesTblStmt, first, count)
	} else {
		rows, err = c.db.Query(selectAllFilesTblStmt)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning files from the database.")
	}
	var (
		id, doctorID           int32
		patientAvsNumber, data string
	)
	defer rows.Close()
	for rows.Next() {
		var err = rows.Scan(&id, &patientAvsNumber, &doctorID, &data)
		if err != nil {
			return nil, errors.Wrapf(err, "error while returning files from the database.")
		}
		files = append(files, File{
			ID:               id,
			PatientAVSNumber: patientAvsNumber,
			DoctorID:         doctorID,
			Data:             data,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning files from the database.")
	}

	return files, nil
}

// Count count all the rows from the database
func (c *Database) Count() (int32, error) {
	var rows, err = c.db.Query(countFilesTblStmt)
	if err != nil {
		return 0, errors.Wrapf(err, "error while counting the files from the database.")
	}
	var count int32
	defer rows.Close()
	for rows.Next() {
		var err = rows.Scan(&count)
		if err != nil {
			return 0, errors.Wrapf(err, "error while counting the files from the database.")
		}
	}
	err = rows.Err()
	if err != nil {
		return 0, errors.Wrapf(err, "error while counting the files from the database.")
	}
	return count, nil
}
