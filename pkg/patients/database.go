package patients

import "database/sql"

const (
	createPatientTblStmt = `CREATE TABLE patients (
		id INT,
		first_name STRING,
		last_name STRING,
		birth_date DATE,
		avs_number STRING
		PRIMARY KEY (id))`
)

type cockroachDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// InitDatabase initializes the database
func InitDatabase(db cockroachDB) {
	// Init DB: create patient table.
	db.Exec(createPatientTblStmt)
}
