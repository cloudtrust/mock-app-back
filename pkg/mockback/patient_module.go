package mockback

// PatientModule contains the business logic for the patients.
type PatientModule interface {
}

type patientModule struct {
}

// NewPatientModule returns a patient module
func NewPatientModule() PatientModule {
	return &patientModule{}
}
