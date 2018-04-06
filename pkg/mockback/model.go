package mockback

import "time"

// Hospital represents a hospital
type Hospital struct {
	ID          int32        `json:"id"`
	Name        string       `json:"name"`
	City        string       `json:"city"`
	Departments []Department `json:"departments"`
}

// Department represents the department of a hospital
type Department struct {
	ID       int32    `json:"id"`
	Name     string   `json:"name"`
	Hospital Hospital `json:"-"`
	Doctors  []Doctor `json:"doctors"`
}

// Doctor represents a doctor
type Doctor struct {
	ID          int32        `json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Departments []Department `json:"-"`
	Patients    []Patient    `json:"patients"`
}

// Patient represents a patient
type Patient struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
	AVSNumber string    `json:"avsNumber"`
	Doctors   []Doctor  `json:"-"`
}
