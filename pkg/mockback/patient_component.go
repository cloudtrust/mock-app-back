package mockback

import "context"

// PatientComponent is the patient business component interface.
type PatientComponent interface {
	ListAllPatients(context.Context) []Patient
}

// PatientComponent is the patient business component.
type patientComponent struct {
	patientModule PatientModule
}

func (c *patientComponent) ListAllPatients(ctx context.Context) []Patient {
	return []Patient{}
}

// NewPatientComponent returns a patient business component
func NewPatientComponent(patientModule PatientModule) PatientComponent {
	return &patientComponent{
		patientModule: patientModule,
	}
}
