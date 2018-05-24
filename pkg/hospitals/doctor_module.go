package hospitals

import "context"

// Doctor represents a doctor
type Doctor struct {
	ID          int32        `json:"id,omitempty"`
	FirstName   string       `json:"firstName,omitempty"`
	LastName    string       `json:"lastName,omitempty"`
	Departments []Department `json:"departments,omitempty"`
	PatientsIds []int32      `json:"patients,omitempty"`
}

// DoctorsModule contains the business logic for the doctors.
type DoctorsModule interface {
	ListAll(ctx context.Context) ([]Doctor, error)
}

type doctorsModule struct {
	doctorsDatabase DoctorsDatabase
}

func (c *doctorsModule) ListAll(ctx context.Context) ([]Doctor, error) {
	return c.doctorsDatabase.ReadFromDb()
}

// NewDoctorModule returns a hospital/departments module
func NewDoctorModule(doctorsDatabase DoctorsDatabase) DoctorsModule {
	return &doctorsModule{
		doctorsDatabase: doctorsDatabase,
	}
}
