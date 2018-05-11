package patients

import (
	"context"
	"time"
)

// Patient represents a patient
type Patient struct {
	ID         int32     `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	BirthDate  time.Time `json:"birthDate"`
	AVSNumber  string    `json:"avsNumber"`
	DoctorsIds []int32   `json:"-"`
}

// Module contains the business logic for the patients.
type Module interface {
	ListAll(ctx context.Context) ([]Patient, error)
}

type module struct {
	cockroachModule CockroachModule
}

func (c *module) ListAll(ctx context.Context) ([]Patient, error) {
	return c.cockroachModule.ReadFromDb()
}

// NewModule returns a patient module
func NewModule(cockroachModule CockroachModule) Module {
	return &module{
		cockroachModule: cockroachModule,
	}
}
