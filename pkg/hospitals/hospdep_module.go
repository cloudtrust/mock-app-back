package hospitals

import "context"

// Hospital represents a hospital
type Hospital struct {
	ID          int32        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	City        string       `json:"city,omitempty"`
	Departments []Department `json:"departments,omitempty"`
}

// Department represents the department of a hospital
type Department struct {
	ID       int32    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Hospital Hospital `json:"hospital,omitempty"`
	Doctors  []Doctor `json:"doctors,omitempty"`
}

// HospDepModule contains the business logic for the hospitals/departments.
type HospDepModule interface {
	ListAllHospitals(ctx context.Context) ([]Hospital, error)
	ListAllDepartments(ctx context.Context) ([]Department, error)
}

type hospDepModule struct {
	hospDepDatabase HospDepDatabase
}

func (c *hospDepModule) ListAllHospitals(ctx context.Context) ([]Hospital, error) {
	return c.hospDepDatabase.ReadHospitalsFromDb()
}

func (c *hospDepModule) ListAllDepartments(ctx context.Context) ([]Department, error) {
	return c.hospDepDatabase.ReadDepartmentsFromDb()
}

// NewHospDepModule returns a hospital/departments module
func NewHospDepModule(hospDepDatabase HospDepDatabase) HospDepModule {
	return &hospDepModule{
		hospDepDatabase: hospDepDatabase,
	}
}
