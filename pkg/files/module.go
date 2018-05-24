package files

import "context"

// File represents a patient file
type File struct {
	ID               int32  `json:"id,omitempty"`
	PatientAVSNumber string `json:"patientAvsNumber,omitempty"`
	DoctorID         int32  `json:"doctorId"`
	Data             string `json:"data"`
}

// Module contains the business logic for the files.
type Module interface {
	ListAll(ctx context.Context) ([]File, error)
}

type module struct {
	database Database
}

func (c *module) ListAll(ctx context.Context) ([]File, error) {
	return c.database.ReadFromDb()
}

// NewModule returns a files module
func NewModule(database Database) Module {
	return &module{
		database: database,
	}
}
