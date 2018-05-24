package hospitals

import "context"

// Doctor represents a doctor
type Doctor struct {
	ID          int32        `json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Departments []Department `json:"departments"`
	PatientsIds []int32      `json:"patients"`
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
