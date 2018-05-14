package hospitals

import "context"

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

// Module contains the business logic for the hospitals/departments.
type Module interface {
	ListAllHospitals(ctx context.Context) ([]Hospital, error)
	ListAllDepartments(ctx context.Context) ([]Department, error)
}

type module struct {
	cockroachModule CockroachModule
}

func (c *module) ListAllHospitals(ctx context.Context) ([]Hospital, error) {
	return c.cockroachModule.ReadHospitalsFromDb()
}

func (c *module) ListAllDepartments(ctx context.Context) ([]Department, error) {
	return c.cockroachModule.ReadDepartmentsFromDb()
}

// NewModule returns a hospital/departments module
func NewModule(cockroachModule CockroachModule) Module {
	return &module{
		cockroachModule: cockroachModule,
	}
}
