package hospitals

// Doctor represents a doctor
type Doctor struct {
	ID          int32        `json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Departments []Department `json:"-"`
	PatientsIds []int32      `json:"patients"`
}
