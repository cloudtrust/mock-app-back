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
	ListSome(ctx context.Context, first int32, rows int32) ([]File, error)
	Count(ctx context.Context) (int32, error)
}

type module struct {
	database Database
}

func (c *module) ListAll(ctx context.Context) ([]File, error) {
	return c.database.ReadFromDb(-1, -1)
}

func (c *module) ListSome(ctx context.Context, first int32, rows int32) ([]File, error) {
	return c.database.ReadFromDb(first, rows)
}

func (c *module) Count(ctx context.Context) (int32, error) {
	return c.database.Count()
}

// NewModule returns a files module
func NewModule(database Database) Module {
	return &module{
		database: database,
	}
}
