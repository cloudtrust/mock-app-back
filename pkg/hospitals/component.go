package hospitals

import "context"

// Component is the hospital business component interface.
type Component interface {
	ListAllHospitals(context.Context) ([]Hospital, error)
	ListAllDepartments(context.Context) ([]Department, error)
	ListAllDoctors(context.Context) ([]Doctor, error)
}

// component is the patient business component.
type component struct {
	hospDepModule HospDepModule
	doctorsModule DoctorsModule
}

func (c *component) ListAllHospitals(ctx context.Context) ([]Hospital, error) {
	return c.hospDepModule.ListAllHospitals(ctx)
}

func (c *component) ListAllDepartments(ctx context.Context) ([]Department, error) {
	return c.hospDepModule.ListAllDepartments(ctx)
}

func (c *component) ListAllDoctors(ctx context.Context) ([]Doctor, error) {
	return c.doctorsModule.ListAll(ctx)
}

// NewComponent returns a patient business component
func NewComponent(hospDepModule HospDepModule, doctorsModule DoctorsModule) Component {
	return &component{
		hospDepModule: hospDepModule,
		doctorsModule: doctorsModule,
	}
}
