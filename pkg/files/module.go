package files

// File represents a patient file
type File struct {
	ID               int32  `json:"id,omitempty"`
	PatientAVSNumber string `json:"patientAvsNumber,omitempty"`
	DoctorID         int32  `json:"doctorId"`
	Data             string `json:"data"`
}
