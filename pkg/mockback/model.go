package mockback

import "time"

// Hospital represents a hospital
type Hospital struct {
	ID          int32
	Name        string
	City        string
	Departments []Department
}

// Department represents the department of a hospital
type Department struct {
	ID       int32
	Name     string
	Hospital Hospital
	Doctors  []Doctor
}

// Doctor represents a doctor
type Doctor struct {
	ID          int32
	FirstName   string
	LastName    string
	Departments []Department
	Patients    []Patient
}

// Patient represents a patient
type Patient struct {
	ID        int32
	FirstName string
	LastName  string
	BirthDate time.Time
	AVSNumber string
	Doctors   []Doctor
}
