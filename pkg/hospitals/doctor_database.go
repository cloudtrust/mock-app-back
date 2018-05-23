package hospitals

import "github.com/pkg/errors"

const (
	// TODO : A doctor can be in multiple departments
	createDoctorsTblStmt = `CREATE TABLE IF NOT EXISTS doctors (
		id INT,
		first_name STRING,
		last_name STRING,
		department_id INT,
		PRIMARY KEY (id))`
	selectAllDoctorsTblStmt = `SELECT * FROM doctors`
)

// DoctorsDatabase deals with the communication with CockroachDB.
type DoctorsDatabase struct {
	db cockroachDB
}

// InitDoctorsDatabase initializes the database
func InitDoctorsDatabase(db cockroachDB) (*DoctorsDatabase, error) {
	// Init DB: create doctors table.
	var _, err = db.Exec(createDoctorsTblStmt)
	if err != nil {
		return nil, err
	}

	return &DoctorsDatabase{
		db: db,
	}, nil
}

// ReadFromDb returns all the hospitals from the database
func (c *DoctorsDatabase) ReadFromDb() ([]Doctor, error) {
	var doctors = []Doctor{}
	var rows, err = c.db.Query(selectAllDoctorsTblStmt)
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the doctors from the database.")
	}
	var (
		id, departmentID    int32
		firstName, lastName string
	)
	defer rows.Close()
	for rows.Next() {
		var err = rows.Scan(&id, &firstName, &lastName, &departmentID)
		if err != nil {
			return nil, errors.Wrapf(err, "error while returning all the doctors from the database.")
		}
		doctors = append(doctors, Doctor{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Departments: []Department{Department{
				ID: departmentID,
			}},
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "error while returning all the doctors from the database.")
	}

	return doctors, nil
}
